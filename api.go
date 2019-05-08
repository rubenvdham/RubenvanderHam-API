package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const epochAge int64 = 894697740

/* Routes */
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Welcome to the API, For documentation visit '[prefix]/documentation'")
	logRequest(r)
}

func age(w http.ResponseWriter, r *http.Request) {
	years, days := calculateAge()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `[{"years": %.2f, "days": %d}]`, years, days)
	logRequest(r)
}

func documentation(w http.ResponseWriter, r *http.Request) {
	body := `[{"location": "/age", "methods": "GET", "returnType": "JSON", "description": "Returns my age with keys 'years' and 'days' and their corresponding values"},
	 {"location": "/myip", "methods": "GET", "returnType": "HTML", "description": "Returns the IP address of the client"},
	{"location": "/documentation", "methods": "GET", "returnType": "JSON", "description": "Returns the documentation of this API in JSON"}]`

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, body)
	logRequest(r)
}

func myIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", strings.Split(r.RemoteAddr, ":")[0])
	logRequest(r)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Nothing here, sorry\nTry '\\documentation' for more info")
	} else {
		fmt.Fprintln(w, "Error ", status)
	}
	log.Printf("- %s - %s \t%d\n", strings.Split(r.RemoteAddr, ":")[0], r.URL.Path, status)
}

/* Helper functions */
func calculateAge() (float64, int64) {
	seconds := time.Now().Unix() - epochAge
	hours := seconds / 3600
	days := hours / 24
	years := float64(days) / float64(365.25)

	return years, days
}

func logRequest(r *http.Request) {
	log.Printf("- %s - %s \t%d\n", strings.Split(r.RemoteAddr, ":")[0], r.URL.Path, http.StatusOK)
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/age", age)
	http.HandleFunc("/myip", myIP)
	http.HandleFunc("/documentation", documentation)
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}

func main() {
	handleRequests()
}
