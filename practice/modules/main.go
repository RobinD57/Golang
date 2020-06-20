package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	quoteV3 "rsc.io/quote/v3"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logger, _ := zap.NewProduction()
	logger.Info("Successfully performed HTTP request")
	logger.Info(quoteV3.HelloV3())
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}