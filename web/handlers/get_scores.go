package handlers

import (
	"github.com/infernalfire72/flame/cache/users"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/osuapi"
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

	queryMode, err := qs.GetUint("m")
	if err != nil {

	}

	mode := constants.Mode(queryMode)
	mode.Clamp()

	queryMods, err := qs.GetUint("mods")
	if err != nil {

	}

	mods := constants.Mod(queryMods)

	filter, err := qs.GetUint("v")
	if err != nil {

	}

	relax := mods.Has(constants.ModRelax)

	p.SetRelaxing(relax)

	b := beatmaps.Get(md5)
	if b == nil {
		// Fetch from osu api if no result, 
		// get from /web/maps/filename if empty, not submitted, if content need update
		if content := osuapi.GetBeatmapContent(file); len(content) == 0 {
			ctx.WriteString("-1|false")
		} else if len(content) < 64 {
			ctx.WriteString("-2|false")
		} else {
			ctx.WriteString("1|false")
		}

		ctx.SetConnectionClose()
		return
	}

	if lb := leaderboards.Get(leaderboards.Identifier{md5, relax, mode}); lb != nil {
		var scores []*layouts.Score

		switch filter {
		case 2:
			scores = lb.ModsFilter(mods)
		case 3:
			scores = lb.FriendsFilter(p.ID)
		case 4:
			scores = lb.CountryFilter(p.Country)
		default:
			lb.Mutex.RLock()
			scores = append(scores, lb.Scores...)
			lb.Mutex.RUnlock()
		}

		ctx.WriteString(b.Online(len(scores)))

		/*if personalBest, index := scores.GetPersonalBest(p.ID); personalBest != nil {
			ctx.WriteString(personalBest.Online(!lb.Relax || b.Status == constants.StatusLoved, p.FullName(), index+1))
		} else {
			ctx.WriteString("\n")
		}*/

		limit := 100
		if p.Privileges.Has(constants.UserPremium) {
			limit = 500
		} else if p.Privileges.Has(constants.UserDonor) {
			limit = 250
		}

		if len(scores) > limit {
			scores = scores[:limit]
		}

		for _, score := range scores {
			if u := users.Get(score.User); u != nil && (u.Privileges.Has(constants.UserPublic) || u.ID == p.ID) {
				/*ctx.WriteString(score.Online(!lb.Relax || b.Status == constants.StatusLoved, u.FullName(), i+1))*/
			}
		}
	}
}
