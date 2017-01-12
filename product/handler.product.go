package product

import (
    "net/http"
    "fmt"
    "encoding/json"
    "strconv"
    "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func GetProductHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  w.Header().Set("content-type", "application/json")
  product := GetProduct()

  json.NewEncoder(w).Encode(product)
  return
}

func SearchProductHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  w.Header().Set("content-type", "application/json")
  idint, _ := strconv.ParseInt(ps.ByName("id"), 10, 64)
  produc, err := GetProductById(idint)
  if(err != nil){
      json.NewEncoder(w).Encode("false")
      return
  }
  json.NewEncoder(w).Encode(produc)
  return
}

func AddProductHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  w.Header().Set("content-type", "application/json")
  idint, _ := strconv.ParseInt(ps.ByName("id"), 10, 64)

  AddProduct(Product{Name:ps.ByName("name"),Id:idint, Description:"ini mobile terbagusefef"})

  json.NewEncoder(w).Encode("ok")
  return
}


  //fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
