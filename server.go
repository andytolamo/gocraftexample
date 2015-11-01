package main

import (
	"github.com/gocraft/web"
	"fmt"
	"net/http"
	"io/ioutil"
)

const ViewFolder = "views/";

type Context struct {
	Input string
}


func GetPost(rw web.ResponseWriter, req *web.Request) {
	exist := req.FormValue("test")

	if len(exist) > 0 {
		fmt.Fprint(rw, exist)
	} else {
		fmt.Fprint(rw, "Does not exist")
	}
}

func GetGet(rw web.ResponseWriter, req *web.Request) {
	exist := req.URL.Query().Get("test")

	if len(exist) > 0 {
		fmt.Fprint(rw, exist)
	} else {
		fmt.Fprint(rw, "Does not exist")
	}
}

func ShowView(rw web.ResponseWriter, req *web.Request) {
	dat, _ := ioutil.ReadFile(ViewFolder+"show.html")
	fmt.Fprint(rw, string(dat))

}



func main() {
	router := web.New(Context{})                 // Create your router
	router.Middleware(web.LoggerMiddleware)           // Use some included middleware
	router.Middleware(web.ShowErrorsMiddleware)       // ...
	router.Post("/posthere", GetPost) // route where to post 'test' key and value
	router.Get("/gethere", GetGet) // route where to give GET paramter 'test' and value
	router.Get("/showview", ShowView) // route that shows view file

	http.ListenAndServe("localhost:3000", router)   // Start the server!
}
