package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/utils"

	"github.com/infernalfire72/flame/bancho/players"
	"github.com/infernalfire72/flame/web/router"
)

func UploadScreenshot(ctx router.WebCtx) {
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Error(err)
		return
	}

	if len(form.File) != 1 {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	username, ok := form.Value["u"]
	if !ok || len(username) != 1 {
		ctx.SetStatusCode(http.StatusUnauthorized)
		return
	}

	password, ok := form.Value["p"]
	if !ok || len(password) != 1 {
		ctx.SetStatusCode(http.StatusUnauthorized)
		return
	}

	if p := players.FindUsername(username[0]); p != nil {
		if !p.VerifyPassword(password[0]) {
			ctx.SetStatusCode(http.StatusUnauthorized)
			return
		}

		var (
			fileName string
			filePath string
		)

	generate:
		fileName = utils.RandString(10)
		filePath = fmt.Sprintf(config.Web.ScreenshotPath, fileName)

		if _, err = os.Stat(filePath); err == nil {
			goto generate
		}

		if !os.IsNotExist(err) {
			log.Error(err)
			ctx.SetStatusCode(http.StatusInternalServerError)
			return
		}

		file, err := form.File["ss"][0].Open()
		if err != nil {
			log.Error(err)
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}
		defer file.Close()

		dst, err := os.Create(filePath)
		if err != nil {
			log.Error(err)
			ctx.SetStatusCode(http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			log.Error(err)
			ctx.SetStatusCode(http.StatusInternalServerError)
			return
		}

		if index := strings.LastIndex(filePath, "/") + 1; index != 0 {
			fileName = filePath[index:]
		}

		ctx.WriteString(fileName)
	} else {
		ctx.SetStatusCode(http.StatusUnauthorized)
	}
}
