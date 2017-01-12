package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	p "github.com/primaap/hello-world/product"
)

func main() {

	router := httprouter.New()
  router.GET("/", p.GetProductHandler)
	router.GET("/product", p.GetProductHandler)
	router.GET("/product/:id", p.SearchProductHandler)
  router.POST("/product/:id/:name", p.AddProductHandler)

  log.Fatal(http.ListenAndServe(":8080", router))
}
