package main

import (
	"app/src/interfaces"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"log"
	"os"
)

// EnvLoad 環境変数の取得
func EnvLoad() {
	if os.Getenv("ENV_MODE") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		// TODO dev環境以外の場合は、.envファイルではなく、HOSTサーバーの環境変数を使用。
	}
}

func main() {
	EnvLoad()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET},
	}))

	t := &Template{
		templates: template.Must(template.ParseGlob("src/public/views/*.html")),
	}
	e.Renderer = t

	interfaces.Routes(e)
	interfaces.Run(e, os.Getenv("PORT"))
}

//HTML templeteのレンダリング用
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
