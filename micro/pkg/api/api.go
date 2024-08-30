package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"stepic-go-basic/micro/pkg/db"
	"stepic-go-basic/micro/pkg/logger"
	"stepic-go-basic/micro/pkg/metrics"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
)

type URLMapping struct {
	OriginalURL string `json:"original_url"`
}

type API struct {
	router *mux.Router
	db     db.Interface
}

func New(db db.Interface) *API {
	api := API{
		router: mux.NewRouter(),
		db:     db,
	}

	api.Endpoints()

	return &api
}

func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) Endpoints() {

	api.router.Handle("/metrics", promhttp.Handler())
	api.router.HandleFunc("/add", api.addLinkHandler).Methods(http.MethodPost)
	api.router.HandleFunc("/{short}", api.redirectHandler).Methods(http.MethodGet)

}

func (api *API) addLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	start := time.Now()
	defer func() {
		metrics.RequestDuration.WithLabelValues("shorten").Observe(time.Since(start).Seconds())
		metrics.RequestCounter.WithLabelValues("shorten").Inc()
	}()

	var mapping URLMapping

	if err := json.NewDecoder(r.Body).Decode(&mapping); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shortLink := generateShortLink()

	response := db.URL{Short: "http://localhost:8080/" + shortLink, Orig: mapping.OriginalURL}

	api.db.AddURL(r.Context(), response)

	logger.Logger.Info().Str("method", r.Method).Str("short_link", shortLink).Str("long_link", mapping.OriginalURL).Msg("Shortened link created")

	//response := map[string]string{"short_link": "https://lynks.org/" + shortLink}
	//w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (api *API) redirectHandler(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	defer func() {
		metrics.RequestDuration.WithLabelValues("redirect").Observe(time.Since(start).Seconds())
		metrics.RequestCounter.WithLabelValues("redirect").Inc()
	}()

	shortLink := r.URL.Path[len("/"):]

	println(shortLink)

	res, _ := api.db.Redirect(r.Context(), shortLink)

	println(res)

	logger.Logger.Info().Str("method", r.Method).Str("short_link", shortLink).Str("long_link", res).Msg("Redirecting to long link")

	http.Redirect(w, r, res, http.StatusFound)

	//response := map[string]string{"original_url": res}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(response)

}

func generateShortLink() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortLink := make([]byte, 5)
	for i := range shortLink {
		shortLink[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortLink)
}
