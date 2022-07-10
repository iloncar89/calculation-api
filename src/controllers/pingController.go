package controllers

import (
	"github.com/iloncar89/calculation-api/src/utils/logger"
	"net/http"
)

const (
	pong = "pong"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct{}

//Ping handler function for /divide url. Resolves GET http function on /ping url.
func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		logger.InfoLogger.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		_, err := w.Write([]byte(pong))
		if err != nil {
			logger.ErrorLogger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger.InfoLogger.Print("Response sent ", pong, "for GET request  /ping\n")
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
