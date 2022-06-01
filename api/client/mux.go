package client

import (
	"context"
	"log"
	"net/http"
	"os"

	"go-api-auction/db"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var ctx context.Context
var collection *mongo.Collection
var router *mux.Router

func MuxHandler() {
	router, ctx, client, collection = db.MongoClient("auction-collection")

	router.HandleFunc("/api/draw/", getDraws).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func getDraws(w http.ResponseWriter, r *http.Request) {

}
