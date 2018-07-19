// demonstrate how to use buffalo package
package main

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

func main() {
	app := newApp()
	if app == nil {
		return
	}

	app.Serve()
}

func newApp() *buffalo.App {

	app := buffalo.New(buffalo.Options{}) //Note: port defaults to 3000

	app.GET("/hello", hellobuffalo)

	return app

}

func hellobuffalo(c buffalo.Context) error {
	return c.Render(http.StatusOK, render.String("hello from buffalo"))
}
