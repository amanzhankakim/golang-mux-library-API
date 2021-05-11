package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Router *mux.Router
	DB *mongo.Database
}

func (a *App) Initialize() {
	

	a.Router = mux.NewRouter()
	a.initializeRoutes()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://amanzhan:1nQFGlOL7YqZB5Y1@cluster0.c9cpf.mongodb.net/test"))
	if err != nil {
		log.Fatal(err)
	}

	a.DB = client.Database("test")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/books", a.getBooks).Methods("GET")
	a.Router.HandleFunc("/book", a.createBook).Methods("POST")
	a.Router.HandleFunc("/book/{name}", a.getBook).Methods("GET")
	a.Router.HandleFunc("/book/{name}", a.updateBook).Methods("PUT")
	a.Router.HandleFunc("/book/{name}", a.deleteBook).Methods("DELETE")
}

func (a *App) getBooks(w http.ResponseWriter, r *http.Request) {
	
	col := a.DB.Collection("Books")

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	
	result, err := col.Find(ctx, bson.M{})
	if err != nil {
		fmt.Print(err)
	}

	var b []bson.M
	if err = result.All(ctx, &b); err != nil {
    	log.Fatal(err)
	}
	json.NewEncoder(w).Encode(b)

}
func (a *App) getBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	col := a.DB.Collection("Books")

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	
	result, err := col.Find(ctx, bson.M{"name": params["name"]})
	if err != nil {
		fmt.Print(err)
	}

	var b []bson.M
	if err = result.All(ctx, &b); err != nil {
    	log.Fatal(err)
	}
	json.NewEncoder(w).Encode(b)

}
func (a *App) createBook(w http.ResponseWriter, r *http.Request) {
	
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)

	col := a.DB.Collection("Books")

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	result, err := col.InsertOne(ctx, book)
	if err != nil {
		fmt.Print(err)
	}

	json.NewEncoder(w).Encode(result)
	
}
func (a *App) updateBook(w http.ResponseWriter, r *http.Request) {


}
func (a *App) deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	col := a.DB.Collection("Books")

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	
	result, err := col.DeleteOne(ctx, bson.M{"name": params["name"]})
	if err != nil {
		fmt.Print(err)
	}

	json.NewEncoder(w).Encode(result)
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}