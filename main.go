package main

import(
	"fmt"
	"net/http"

)

func main() {

	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		_, _ = fmt.Fprintf(w, "Welcome to my website！")
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/",http.StripPrefix("/static/",fs))

	_ = http.ListenAndServe(":80", nil)
}
