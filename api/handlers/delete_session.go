package handlers

import (
	"net/http"

	"github.com/keratin/authn-server/api/sessions"
	"github.com/keratin/authn-server/app"
	"github.com/keratin/authn-server/services"
)

func DeleteSession(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := services.SessionEnder(app.RefreshTokenStore, sessions.GetRefreshToken(r))
		if err != nil {
			app.Reporter.ReportRequestError(err, r)
		}

		sessions.Set(app.Config, w, "")

		w.WriteHeader(http.StatusOK)
	}
}
