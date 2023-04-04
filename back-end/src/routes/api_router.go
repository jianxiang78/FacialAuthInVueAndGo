package routes

import (
	"../conf"
	"../handlers"
	"fmt"
	"log"
	"net/http"
)

func ApiRouter() {
	http.HandleFunc("/dashboard", handlers.Dashboard)
	http.HandleFunc("/config", handlers.Config)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/loginWithFace", handlers.LoginWithFace)
	http.HandleFunc("/userEdit", handlers.UserEdit)
	port := conf.GetEnv().ServerPort
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}