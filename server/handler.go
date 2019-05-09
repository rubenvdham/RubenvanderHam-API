package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Handlers struct {
	logger   *log.Logger
	epochAge int64
}

func NewHandler(logger *log.Logger, epochAge int64) *Handlers {
	return &Handlers{
		logger:   logger,
		epochAge: epochAge}
}

/* Routes */
func (h *Handlers) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.Header().Set("Status", string(http.StatusOK))
	fmt.Fprint(w, "Welcome to the API, For documentation visit '[prefix]/documentation'")
}

func (h *Handlers) age(w http.ResponseWriter, r *http.Request) {
	years, days := h.calculateAge()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `[{"years": %.2f, "days": %d}]`, years, days)
}

func (h *Handlers) documentation(w http.ResponseWriter, r *http.Request) {
	body := `[{"location": "/age", "methods": "GET", "returnType": "JSON", "description": "Returns my age with keys 'years' and 'days' and their corresponding values"},
	 {"location": "/myip", "methods": "GET", "returnType": "HTML", "description": "Returns the IP address of the client"},
	{"location": "/documentation", "methods": "GET", "returnType": "JSON", "description": "Returns the documentation of this API in JSON"}]`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, body)
}

func (h *Handlers) myIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", strings.Split(r.RemoteAddr, ":")[0])
}

func (h *Handlers) errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Nothing here, sorry\nTry '\\documentation' for more info")
	} else {
		fmt.Fprintln(w, "Error ", status)
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.log(h.home))
	mux.HandleFunc("/age", h.log(h.age))
	mux.HandleFunc("/myip", h.log(h.myIP))
	mux.HandleFunc("/documentation", h.log(h.documentation))
}

func (h *Handlers) log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if h.logger != nil {
			startTime := time.Now()
			defer h.logger.Printf("- %s - %s - %.2f ms\n", strings.Split(r.RemoteAddr, ":")[0], r.URL.Path, float64(time.Now().Sub(startTime).Nanoseconds())/1000)
		}
		next(w, r)
	}
}

/* Helper functions */
func (h *Handlers) calculateAge() (float64, int64) {
	seconds := time.Now().Unix() - h.epochAge
	hours := seconds / 3600
	days := hours / 24
	years := float64(days) / float64(365.25)

	return years, days
}
