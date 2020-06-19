package web

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/log"
)

func Start(conf *config.WebConfig) {
	r := router.New()

	port := fmt.Sprintf(":%d", conf.Port)

	r.POST("/web/osu-submit-modular-selector.php", ssTest)

	log.Info("Started Web at", port)
	fasthttp.ListenAndServe(port, r.Handler)
}

func ssTest(ctx *fasthttp.RequestCtx) {
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Error(err)
		return
	}

	for k, v := range form.Value {
		log.Info(k, v)
	}
}