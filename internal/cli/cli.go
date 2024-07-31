package cli

import (
	handler "sportix-cli/internal/handler/user"
	"sportix-cli/internal/session"

	"github.com/rivo/tview"
)

type Handler struct {
	User handler.UserHandler
}

func MainCLI(app *tview.Application, handler Handler) {
	if session.UserSession == nil {
		AuthModal(app, handler)
	} else {
		if session.UserSession.Role == "user" {
			UserDashboardPage(app, handler)
		} else {
			// TODO: admin dasboard ya, semangat!
		}
	}
}
