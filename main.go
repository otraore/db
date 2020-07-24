package main

import (
	"io/ioutil"
	"log"
	"net/url"

	"github.com/webview/webview"
)

func main() {
	content, err := ioutil.ReadFile("./views/app.html")
	if err != nil {
		log.Fatal("failed to load file: ", err)
	}

	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle("Delete By Time")
	w.SetSize(1100, 600, webview.HintNone)
	w.Navigate("data:text/html," + url.QueryEscape(string(content)))
	w.Run()
}
