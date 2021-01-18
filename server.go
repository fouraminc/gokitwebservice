package gokitwebservice

import (
	"context"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func NewHttpServer(ctx context.Context, endpoints Endpoints) http.Handler{

	r:= mux.NewRouter()
	r.Use(commonMiddleWare)

	r.Methods("GET").Path("/status").Handler(httptransport.NewServer(
		endpoints.StatusEndPoint,
		decodeStatusRequest,
		encodeResponse,
		))

	r.Methods("GET").Path("/get").Handler(httptransport.NewServer(
		endpoints.GetEndPoint,
		decodeGetRequest,
		encodeResponse,
		))

	r.Methods("POST").Path("/validate").Handler(httptransport.NewServer(
		endpoints.ValidateEndpoint,
		decodeValidateRequest,
		encodeResponse,
		))

	return r
}


func commonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
