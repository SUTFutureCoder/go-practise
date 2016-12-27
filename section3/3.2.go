package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {

}

func (p *MyMux) ServeHttp(w http.ResponseWriter, r *http.Request){
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request){
	fmt.Println(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
