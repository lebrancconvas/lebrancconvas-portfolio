package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/node_modules/",http.StripPrefix("/node_modules/",http.FileServer(http.Dir("node_modules"))))
	http.HandleFunc("/",index)
	fmt.Println("Server Run Complete")
	http.ListenAndServe(":2000",nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"index.html")
}