package web

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/log"

	"github.com/infernalfire72/flame/web/handlers"
	"github.com/infernalfire72/flame/web/router"
)

func wrap(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		log.Warn("Web", string(ctx.Path()))

		next(ctx)
	}
}

func Start(conf *config.WebConfig) {
	r := router.NewRouter()

	g := r.Group("/web")
	{
		g.Get("/osu-osz2-getscores.php", handlers.GetScores)
		g.Post("/screenshot_upload.php", handlers.UploadScreenshot)
	}
	//g.POST("/osu-submit-modular-selector.php", ssTest)

	port := fmt.Sprintf(":%d", conf.Port)
	log.Info("Started Web at", port)
	go fasthttp.ListenAndServe(port, wrap(r.Handler))
}

func allT(ctx router.WebCtx) {
	fmt.Println(string(ctx.Path()))
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
