package infrastructures

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Router interface {
	Start() error
	AddRoute(path string, handler func(http.ResponseWriter, *http.Request), methods ...string)
}

type RouterHandler struct {
	Router
	MuxRouter *mux.Router
}

func NewRouter() *RouterHandler {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	return &RouterHandler{
		MuxRouter: router,
	}
}

func (router *RouterHandler) Start() error {
	log.Println("Listening on port:", os.Getenv("API_PORT"))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("API_PORT")), router.MuxRouter)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	return err
}

func (router *RouterHandler) AddRoute(
	path string, handler func(http.ResponseWriter, *http.Request), methods ...string) {
	router.MuxRouter.HandleFunc(path, handler).Methods(methods...)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("API REQUEST:", request.RequestURI)
		next.ServeHTTP(writer, request)
	})
}
