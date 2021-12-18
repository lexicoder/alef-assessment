package main

import (
	"assessment/models"
	"assessment/routes"
	"assessment/sessions"
	"assessment/utils"
	"fmt"
	"log"
	"net/http"
	"os"
)

//Based of https://github.com/EricLau1/Youtube/tree/master/go-webapp
//Related youtube playlist videos: https://www.youtube.com/playlist?list=PLFXr22TafQUs6JqgVOqst70iLtmGJ8nmc
//https://cloud.google.com/go/getting-started/using-cloud-datastore
//https://github.com/GoogleCloudPlatform/golang-samples/tree/steps/getting-started/bookshelf
//many thanks to : Eric Lau de Oliveira @ https://github.com/EricLau1
func main() {
	models.TestConnection()
	sessions.SessionOptions("localhost", "/", 3600, true)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	utils.LoadTemplates("/app/views/*.html")
	fmt.Println("loaded all templates successfully!")
	r := routes.NewRouter()
	http.Handle("/", r)
	fmt.Printf("Listening Port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
