package product

import (
  "database/sql"
  "log"
  _ "github.com/lib/pq"
)

var DB *sql.DB

func init(){
  constring := ""
  db, errOpen := sql.Open("postgres", constring)

  if errOpen != nil {
    log.Fatal(errOpen)
  }

  DB = db
}
