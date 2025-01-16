package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	U "github.com/bschau/tools/internal/usage"
)

var doc = `Usage: sauk v1.2
sauk [OPTIONS]
[OPTIONS]
 -D              Enable debugging
 -d path         Document root
 -h              Help (this page)
 -p num          Listen on port number
`

func main() {
	help := flag.Bool("h", false, "Help")
	docRoot := flag.String("d", ".", "Document root")
	portNum := flag.Int("p", 8080, "Port number")
	debug := flag.Bool("D", false, "Debug")
	flag.Parse()

	if *help {
		U.Usage(doc, 0)
	}

	fs := dotFileHidingFileSystem{http.Dir(*docRoot)}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		past := time.Now().UTC().AddDate(-1, 0, 0).Format(http.TimeFormat)
		headers := w.Header()
		// Cannot set Last-Modified as it is handled by the FileServer
		headers.Set("Cache-Control", "no-store,max-age=0")
		headers.Set("Expires", past)

		if *debug {
			fmt.Println(html.EscapeString(r.URL.Path))
		}

		http.FileServer(fs).ServeHTTP(w, r)
	})
	port := ":" + strconv.Itoa(*portNum)
	if *debug {
		fmt.Println("Listening on port " + port)
	}
	log.Fatal(http.ListenAndServe(port, nil))
	os.Exit(0)
}
