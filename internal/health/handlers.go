package health

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"todorestapp.ut/cmd/config"
)

func HealthHandler(appConfig *config.AppEnv) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		if r.URL.Path != "/" {
			appConfig.AppError.NotFound(w)
			return
		}

		w.Write([]byte("Success"))
	}
}
