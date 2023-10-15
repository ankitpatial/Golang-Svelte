package main

import (
	"fmt"
	"net/http"
	"roofix/config"
	"roofix/server/router"
)

func main() {
	r := router.NewRouter()
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Read().Website.Port), r); err != nil {
		panic(err)
	}
}
