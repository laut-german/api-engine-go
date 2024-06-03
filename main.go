package main

import (
	"fmt"
	"log"
	"net/http"
	serverengine "web-engine-go/server-engine"
	userroute "web-engine-go/user"
)


func main() {
	topLevelRoutes := make(map[string] *serverengine.Route)
	topLevelRoutes["user"] = userroute.New()
	server := serverengine.New(topLevelRoutes)
	fmt.Println("Server listening at port 8000")
	if err:= http.ListenAndServe(":8000", server); err !=nil {
		log.Fatal("An error occurred while server starting")
	}
}