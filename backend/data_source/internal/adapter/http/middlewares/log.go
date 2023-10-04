package middlewares

import (
	"net/http"
)

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//fields := map[string]interface{}{
		//	"remote_addr": r.RemoteAddr,
		//	"request_id":  r.Context().Value("x-req-id"),
		//}

		//l := logger.GetLoggerWithFields(fields)
		//
		//l.Infof("started %s %s", r.Method, r.RequestURI)
		//
		//start := time.Now()
		//rw := &api.ResponseWriter{ResponseWriter: w, Code: http.StatusOK}
		//h.ServeHTTP(rw, r)
		//
		//var level logrus.Level
		//switch {
		//case rw.Code >= 500:
		//	level = logrus.ErrorLevel
		//case rw.Code >= 400:
		//	level = logrus.WarnLevel
		//default:
		//	level = logrus.InfoLevel
		//}
		//l.Logf(
		//	level,
		//	"completed with %d %s in %v",
		//	rw.Code,
		//	http.StatusText(rw.Code),
		//	time.Now().Sub(start),
		//)
	})
}
