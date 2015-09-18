package main

import (
	"fmt"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"github.com/thoas/stats"
	"net/http"
	"time"
	"runtime"
	"github.com/jelinden/selfjs"
	"io/ioutil"
	"encoding/json"
)

type Lorem struct {
	Text string `json:"text"`
}

var lorem = Lorem{"Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."}

func apiFrontPage(c *echo.Context) error {
	return c.JSON(http.StatusOK, rssObject)
}

func apiAnotherPage(c *echo.Context) error {
	return c.JSON(http.StatusOK, lorem)
}

func loremJSON() string {
	json, _ := json.Marshal(lorem)
	return string(json)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fetchFeed()

	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.StripTrailingSlash())
	e.Use(mw.Gzip())
	e.Use(cors.Default().Handler)

	bundle, _ := ioutil.ReadFile("./build/bundle.js")

	// stats
	s := stats.New()
	e.Use(s.Handler)
	e.Get("/stats", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, s.Data())
	})
	// static files
	e.Static("/public/css", "public/css")
	e.Static("/universal.js", "./build/bundle.js")
	e.Favicon("public/favicon.ico")

	e.Get("/", selfjs.New(runtime.NumCPU(), string(bundle), rss))
	e.Get("/about", selfjs.New(runtime.NumCPU(), string(bundle), loremJSON()))

	e.Get("/api/data", apiFrontPage)
	e.Get("/api/anotherpage", apiAnotherPage)
	go tick()
	fmt.Println("serving at port 3000")
	e.Run(":3000")
}

func tick() {
	ticker := time.NewTicker(70 * time.Second)
	for {
		fmt.Println("fetching feed")
		fetchFeed()
		<-ticker.C
	}
}
