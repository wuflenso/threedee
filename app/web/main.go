package main

import (
	"fmt"
	"log"
	"net/http"
	"threedee"
)

/*
 * FIRST LAYER a.k.a Entry Point => main.go
 *
 * Steps to do here:
 * 1. Create an instance of the app that contains the http.Handler item that is needed by
 *    the Http.ListenAndServe method.
 * 2. Log the launch notif using log.Println
 * 3. Execute http.ListenAndServe to launch the server.
 *
 * Leave the main package simple. The objectives of main are to create the app Handler and
 * launch the server. Leave the routing inside threedee.go. Leave the logic/remote api call/
 * db data retrieval/modification in service directory.
 */

func main() {
	app := threedee.NewThreedee()
	log.Println("Threedee service is ready to listen at port 3000")
	err := http.ListenAndServe(":3000", app.Router)
	if err != nil {
		panic(fmt.Sprintf("%s: %s", "Failed to listen and serve", err))
	}
}
