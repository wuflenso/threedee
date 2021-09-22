package middleware

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type Handler func(w http.ResponseWriter, r *http.Request, params httprouter.Params) (int, error)

func Middleware(handle Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		SetupLogger()

		start := time.Now()
		query := r.URL.Query()

		// this executes the methods in handler!
		code, err := handle(w, r, params)

		elapsed := time.Since(start).Seconds() * 1000
		elapsedStr := strconv.FormatFloat(elapsed, 'f', -1, 64)

		if err != nil {
			log.WithFields(log.Fields{
				"time":   elapsedStr,
				"method": r.Method,
				"path":   r.URL.Path,
				"query":  query.Encode(),
				"status": code,
			}).Warning(err.Error()) // this is where the 'msg' comes from
		} else {
			log.WithFields(log.Fields{
				"time":   elapsedStr,
				"method": r.Method,
				"path":   r.URL.Path,
				"query":  query.Encode(),
				"status": code,
			}).Info("success")
		}
	}
}

func SetupLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
}
