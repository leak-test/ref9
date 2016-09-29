package main

import "github.com/leak-test/mux"

var router = mux.NewRouter()

func main() {
    http.Handle("/", router)
}
