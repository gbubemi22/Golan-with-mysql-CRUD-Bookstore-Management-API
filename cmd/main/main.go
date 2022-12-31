package main

import (
	"fmt"
	"github.com/gbubemi22/go-bookstore/pkg/routers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	fmt.Println(http.ListenAndServe("localhost:9010",r))
}
