package afr_kafka

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name             string
	Method           string
	Pattern          string
	HandlerFunc      http.HandlerFunc
	AddToAccessEntry bool
}

type Routes []Route

func CreateRoutes(UC *UserControl) Routes {
	return Routes{
		// Route{
		// 	"HTTP_PurchaseBundle",
		// 	"GET",
		// 	"/HTTP_PurchaseBundle/",
		// 	//Use(UC.HTTP_PurchaseBundle, UC.JWSAuthentication),
		// 	Uc.HTTP_PurchaseBundle,
		// 	false,
		// },

		//Kafak Producer
		// Route{
		// 	"HTTP_KafkaProduceMSG",
		// 	"GET",
		// 	"/HTTP_KafkaProduceMSG/{topic}/{msg}/",
		// 	UC.HTTP_KafkaProduceMSG,
		// 	false,
		// },

		// Route{
		// 	"HTTP_INLiveFeed_SSR_Producer",
		// 	"GET",
		// 	"/HTTP_INLiveFeed_SSR_Producer/",
		// 	UC.HTTP_INLiveFeed_SSR_Producer,
		// 	false,
		// },
	}
}

func (Uc *UserControl) AddToRouter(router *mux.Router, UC *UserControl) {
	// When StrictSlash is set to true, if the route path is "/path/", accessing "/path" will redirect
	// to the former and vice versa
	routes := CreateRoutes(UC)
	for _, route := range routes {
		handler := route.HandlerFunc
		//handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	//Registering Prometheus http handler to metrics endpoint
	router.Path("/metrics").Handler(CustomPrometheusHandler())
	router.Path("/metrics_latency").Handler(CustomPrometheusLatencyHandler())
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////
// authentication functions
// /////////////////////////////////////////////////////////////////////////////////////////////////////
func Use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
