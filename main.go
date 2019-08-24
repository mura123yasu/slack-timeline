/*
 TODO:write something...
*/

package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mura123yasu/slack-timeline/command"
	"github.com/mura123yasu/slack-timeline/domain"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK!")
	})

	e.POST("/timeline", func(c echo.Context) error {
		// get request param
		req := new(domain.SlackRequest)
		if err := c.Bind(req); err != nil {
			return err
		}

		var data domain.SlackResponse
		switch req.Text {
		case "":
			data = command.Suggest(*req)
		case "fin":
			data = command.Aggregate(*req)
		default:
			data = command.Cache(*req)
		}

		return c.JSON(http.StatusOK, data)
	})

	// サーバー起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}