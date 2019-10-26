package main

import (
	templates "./templates"
	"fmt"
	"os"
	"time"
)

func Init() {
	totalStart := time.Now()
	start := time.Now()
	for _, dirName := range []string{"assets", "pages", "posts"} {
		_ = os.Mkdir(dirName, 0777)
	}
	stop := time.Now()
	fmt.Println(stop.Sub(start))

	start = time.Now()
	BuildSass("./assets/style.css")
	stop = time.Now()
	fmt.Println(stop.Sub(start))

	start = time.Now()
	config := templates.LoadPostblog("_data/_postblog.yml")
	config.Save("postblog.yml")
	//fmt.Printf("%+v\n", config)
	stop = time.Now()
	fmt.Println(stop.Sub(start))

	start = time.Now()
	page := templates.LoadPage("_data/_page.yml")
	page.Save("pages/index.yml")
	//fmt.Printf("%+v\n", page)
	stop = time.Now()
	fmt.Println(stop.Sub(start))

	start = time.Now()
	fmt.Println(templates.RSS(config))
	stop = time.Now()
	fmt.Println(stop.Sub(start))

	totalStop := time.Now()
	fmt.Println("_____________")
	fmt.Println(totalStop.Sub(totalStart))
}

func main() {
	Init()
}
