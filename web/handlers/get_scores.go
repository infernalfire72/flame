package handlers

import (
	"net/http"

	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/log"

	"github.com/infernalfire72/flame/bancho/players"
	"github.com/infernalfire72/flame/cache/beatmaps"
	"github.com/infernalfire72/flame/cache/leaderboards"
	"github.com/infernalfire72/flame/web/router"
)

func GetScores(ctx router.WebCtx) {
	qs := ctx.QueryArgs()

	username := string(qs.Peek("us"))
	if len(username) == 0 {
		ctx.Error("nouser")
		return
	}

	password := string(qs.Peek("ha"))
	if len(password) != 32 {
		ctx.Error("pass")
		return
	}

	p := players.FindUsername(username)
	if p == nil {
		ctx.Error("nouser")
		return
	}

	if password != p.Password {
		log.Warn(username, "failed auth", "get_scores")
		ctx.Error("pass")
		return
	}

	md5 := string(qs.Peek("c"))
	if len(md5) != 32 {
		ctx.Error("beatmap")
		return
	}

	log.Info(p, "requested leaderboard for", md5)

	_, err := qs.GetUint("i")
	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.Error("beatmap")
		return
	}

	file := string(qs.Peek("f"))
	if len(file) < 5 {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.Error("beatmap")
		return
	}

	mode, err := qs.GetUint("m")
	if err != nil {

	}

	mods, err := qs.GetUint("mods")
	if err != nil {

	}

	filter, err := qs.GetUint("v")
	if err != nil {

	}

	relax := (mods & 128) != 0

	if relax {
		clamp(&mode, 0, 2)
	} else {
		clamp(&mode, 0, 3)
	}

	p.SetRelaxing(relax)

	b := beatmaps.Get(md5)
	if b == nil {
		// Fetch from osu api if no result, get from /web/maps/filename if empty, not submitted, if content need update
		b = beatmaps.FetchFromApi(md5, file)
	}

	if b == nil {
		ctx.SetConnectionClose()
		return
	}

	if b.Status <= constants.StatusNeedUpdate {
		ctx.WriteString(b.Online())
		return
	}

	limit := 100
	if p.Privileges.Has(constants.UserPremium) {
		limit = 500
	} else if p.Privileges.Has(constants.UserDonor) {
		limit = 250
	}

	if lb := leaderboards.Get(leaderboards.Identifier{md5, byte(mode), relax}); lb != nil {
		var scores leaderboards.Scores

		switch filter {
		case 2:
			scores = lb.Mods(mods)
		case 3:
			scores = lb.Friends(p.ID)
		case 4:
			scores = lb.Country(p.Country)
		default:
			lb.Mutex.RLock()
			scores = append(make(leaderboards.Scores, 0), lb.Scores...)
			lb.Mutex.RUnlock()
		}

		ctx.WriteString(b.OnlineRanked(len(scores)))

		if personalBest, index := scores.GetPersonalBest(p.ID); personalBest != nil {
			ctx.WriteString(personalBest.String(!lb.Relax || b.Status == constants.StatusLoved, index+1))
		} else {
			ctx.WriteString("\n")
		}

		if len(scores) > limit {
			scores = scores[:limit]
		}

		for i, score := range scores {
			ctx.WriteString(score.String(!lb.Relax || b.Status == constants.StatusLoved, i+1))
		}
	}
}

func clamp(mod *int, min, max int) {
	if *mod > max {
		*mod = max
	} else if *mod < min {
		*mod = min
	}
}
