package product

import (
    "errors"
)

var Products []Product

func GetProduct() []Product{
  return Products
}

func AddProduct(pro Product) error {
  Products = append(Products, pro)
  return nil
}

func GetProductById(id int64) (pr Product, er error) {
  for _, v := range Products{
    if(v.Id == id){
      return v, nil
    }
  }
  return pr, errors.New("notfound")
}
