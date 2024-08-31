package api

import (
	"encoding/json"
	"net/http"
	"stepic-go-basic/cashe/pkg/db"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
)

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
	api.router.HandleFunc("/add", api.saveHandler).Methods(http.MethodPost)
	api.router.HandleFunc("/{short}", api.getHandler).Methods(http.MethodGet)

}

func (api *API) saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var mapping db.URL

	if err := json.NewDecoder(r.Body).Decode(&mapping); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response := db.URL{Short: mapping.Short, Orig: mapping.Orig}

	api.db.SaveURL(r.Context(), response)

	//logger.Logger.Info().Str("method", r.Method).Str("short_link", shortLink).Str("long_link", mapping.OriginalURL).Msg("Shortened link created")

	//response := map[string]string{"short_link": "https://lynks.org/" + shortLink}
	//w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (api *API) getHandler(w http.ResponseWriter, r *http.Request) {

	shortLink := r.URL.Path[len("/"):]

	println(shortLink)

	res, _ := api.db.GetOriginal(r.Context(), shortLink)

	println(res)

	//logger.Logger.Info().Str("method", r.Method).Str("short_link", shortLink).Str("long_link", res).Msg("Redirecting to long link")

	//http.Redirect(w, r, res, http.StatusFound)

	response := map[string]string{"original_url": res}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
