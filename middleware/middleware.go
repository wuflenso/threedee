package middleware

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type Handler func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error

func Middleware(handle Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		SetupLogger()

		start := time.Now()
		query := r.URL.Query()
		/* body := []byte{}
		if r.Body != nil {
			defer r.Body.Close()
			body, _ = ioutil.ReadAll(r.Body)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		} */

		// this executes the methods in handler!
		err := handle(w, r, params)

		elapsed := time.Since(start).Seconds() * 1000
		elapsedStr := strconv.FormatFloat(elapsed, 'f', -1, 64)

		if err != nil {
			log.WithFields(log.Fields{
				"time":   elapsedStr,
				"method": r.Method,
				"path":   r.URL.Path,
				"query":  query.Encode(),
				//"body":   string(body),
			}).Warning(err.Error()) // this is where the 'msg' comes from
		} else {
			log.WithFields(log.Fields{
				"time":   elapsedStr,
				"method": r.Method,
				"path":   r.URL.Path,
				"query":  query.Encode(),
				// "body":   string(body),
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
