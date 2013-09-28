package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	PORT     int    = 9000
	WEB_ROOT string = "./webroot"
)

var PERMIT_SUFFIX = []string{"html", "jpg", "gif", "png"}

func server(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.URL.Path

	for _, v := range PERMIT_SUFFIX {
		if strings.HasSuffix(url, v) {
			resource, err := os.Open(WEB_ROOT + url)
			if err != nil {
				// log.Fatal("Fatal:", err)
				fmt.Println(err)
				fmt.Fprint(w, "404 NOT FOUND!\n")
				w.WriteHeader(404)
				return
			}
			buffer := make([]byte, 1024)
			for {
				n, err := resource.Read(buffer)
				if err != nil && err != io.EOF {
					fmt.Println(err)
					fmt.Fprint(w, "500")
					w.WriteHeader(500)
				}
				if err == io.EOF {
					return
				}
				fmt.Fprintf(w, "%s", buffer[:n])
			}
			return
		}
	}

	fmt.Fprintf(w, "Welcome ~ ~ \n")
	fmt.Fprintf(w, "You are visiting: %s\n", url)
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	if r.Form != nil {
		fmt.Fprint(w, "Here are parameters:\n")
		for key, value := range r.Form {
			fmt.Fprintf(w, "%s => %s", key, value)
		}
	}
}

func main() {
	http.HandleFunc("/", server)
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		log.Fatal("Fatal:", err)
	}
}
