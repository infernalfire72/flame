package handlers

import (
	"crypto/cipher"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/celso-wo/rijndael256"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"

	"github.com/infernalfire72/flame/bancho/bot"
	"github.com/infernalfire72/flame/bancho/channels"
	"github.com/infernalfire72/flame/bancho/players"
	"github.com/infernalfire72/flame/cache/beatmaps"
	"github.com/infernalfire72/flame/cache/leaderboards"
	"github.com/infernalfire72/flame/web/router"
)

func SubmitModular(ctx router.WebCtx) {
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Error(err)
		return
	}

	if len(form.File) != 1 {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	formValue := func(key string) (string, bool) {
		if values, ok := form.Value[key]; ok && len(values) == 1 {
			return values[0], ok
		}

		return "", false
	}

	password, ok := formValue("pass")
	if !ok || len(password) != 32 {
		ctx.SetStatusCode(http.StatusUnauthorized)
		return
	}

	score, ok := formValue("score")
	if !ok {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	iv, ok := formValue("iv")
	if !ok || len(iv) != 44 { // ((4 * 32 / 3) + 3) & ~3 == 44
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	version, ok := formValue("osuver")
	if !ok || len(version) != 8 {
		ctx.Error("oldver")
		return
	}

	decrypted, err := decryptScoreData(score, iv, version)
	if err != nil {
		log.Error(err)
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	parts := strings.SplitN(decrypted, ":", 18)
	if len(parts) != 18 {
		return
	}

	username := parts[1]
	p := players.FindUsername(username)
	if p == nil {
		log.Warn("Score Submission | ", username, "is not online.")
		return
	}

	if !p.VerifyPassword(password) {
		// ctx.SetStatusCode(http.StatusUnauthorized)
		ctx.Error("pass")
		return
	}

	log.Info(p, "wants to submit a score.")

	// Check if they're banned
	if !p.Privileges.Has(constants.UserPublic | constants.UserNormal) {
		return
	}

	restricted := !p.Privileges.Has(constants.UserPublic)

	var s layouts.Score
	err = layouts.ReadScoreLayout(parts, &s)
	if err != nil {
		log.Error(err)
		return
	}

	beatmap := beatmaps.Get(s.BeatmapHash)
	if beatmap == nil {
		ctx.Error("beatmap")
		log.Warn("Score Submission |", "Beatmap doesn't exist in Database.")
		return
	}

	/*
	Code Removed because peppy let unranked shit submit so lets go ig
	if beatmap.Status < constants.StatusRanked {
		ctx.Error("beatmap")
		log.Warn("Score Submission |", "Beatmap is unranked.")
		return
	}*/

	// Playtime here

	// Calculate pp

	maxPerformanceThreshold := func(mode constants.Mode, relax bool) float32 {
		return 1700;
	}

	if !restricted {
		if max := maxPerformanceThreshold(s.Mode, s.Relax); s.Performance > max {

		}
	}

	lb := leaderboards.Get(leaderboards.Identifier{s.BeatmapHash, s.Mode, s.Relax})
	if lb == nil {
		ctx.Error("beatmap")
		return
	}

	oldPersonalBest, _ := lb.FindUserScore(p.ID)


	if oldPersonalBest == nil || s.Performance >= oldPersonalBest.Performance {
		s.Status = constants.ScoreBestPerformance
		// replace score in cache here
	} else {
		s.Status = constants.ScorePassed
	}

	err = s.AddToDatabase(config.Database)
	if err != nil {
		log.Error(err)
	}

	if c := channels.Get("#announce"); c != nil {
		bot.WriteMessagef(c, "%v submitted a score on %s", p, s.BeatmapHash)
	}
}

func decryptScoreData(encrypted, iv, version string) (string, error) {
	c, err := rijndael256.NewCipher([]byte("osu!-scoreburgr---------" + version))
	if err != nil {
		return "", err
	}

	bytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	ivbytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}

	bm := cipher.NewCBCDecrypter(c, ivbytes)

	//dst := make([]byte, len(bytes))
	bm.CryptBlocks(bytes, bytes)
	end := strings.IndexRune(string(bytes), 18)
	if end != -1 && end < len(bytes) {
		bytes = bytes[:end]
	}

	return string(bytes), nil
}
