package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"./server"
)

const epochAge int64 = 894697740
const welcomeMessage = "Started rubenvanderhamAPI"
const serveAddr = "0.0.0.0:8081"

func main() {
	logging := flag.Bool("log", false, "Turn on Logging (to stdout)")
	flag.Parse()

	var logger *log.Logger = nil
	if *logging {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}

	h := server.NewHandler(logger, epochAge)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)
	srv := server.New(mux, serveAddr)

	if *logging {
		logger.Println(welcomeMessage)
		logger.Fatal(srv.ListenAndServe())
	} else {
		fmt.Println(welcomeMessage)
		srv.ListenAndServe()
	}
}
