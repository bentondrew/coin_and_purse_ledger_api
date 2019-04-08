package router

import (
  "net/http"
  "log"
  "time"
  "github.com/Drewan-Tech/coin_and_purse_ledger_service/app/api"
)

/*Router struct contains the service API instance, http.ServeMux
instance, and logger instance. It also contains the baseURL string
which sets the base of the URL endpoints.*/
type Router struct {
  api *api.API
  Mux *http.ServeMux
  logger *log.Logger
  baseURL string
} 

/*NewRouter generates a new instance of a router struct given the
provided logger instance, base URL string, and API instance.*/
func NewRouter(logger *log.Logger, baseURL string, api *api.API) *Router {
  return &Router{
    api: api,
    Mux: http.NewServeMux(),
    logger: logger,
    baseURL: baseURL,
  }
}

func (rtr *Router) endpointLogger(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    startTime := time.Now()
    next(w, r)
    rtr.logger.Printf("Request to path %s processed in %s seconds.\n", r.URL, time.Now().Sub(startTime))
  }
}

/*SetupRoutes sets up the route endpoints in the MUX with the given API.
The base URL is appended to all the endpoints. The logger is used to
log response times to endpoint requests.*/
func (rtr *Router) SetupRoutes() {
  rtr.Mux.HandleFunc("/", rtr.endpointLogger(rtr.api.HandleDefault))
  rtr.Mux.HandleFunc(rtr.baseURL + "/hello", rtr.endpointLogger(rtr.api.HandleHello))
  rtr.Mux.HandleFunc(rtr.baseURL + "/transactions", rtr.endpointLogger(rtr.api.HandleTransactions))
}
