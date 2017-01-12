package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	p "github.com/primaap/hello-world/product"
)

func main() {
	fmt.Print("starting..")

	shop1 := p.Shop{Id: 1, ShopName: "pisang"}

	review1 := p.Review{Id: 1, Review: "produk ini bagus"}
	review2 := p.Review{Id: 2, Review: "produk itu jelek banget"}
	reviews := []p.Review{review1, review2}
	product := p.Product{Id: 1, Name: "hp bagus", Description: "ini produk bagus lho", Review: reviews, Shop: shop1}

	res, _ := json.Marshal(product)
	log.Printf("%+v",product)

	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, string(res))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
