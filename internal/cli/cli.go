package cli

import (
	"sportix-cli/internal/handler"
	"sportix-cli/internal/session"

	"github.com/rivo/tview"
)

type Handler struct {
	User        handler.UserHandler
	Location    handler.LocationHandler
	Category    handler.CategoryHandler
	Field       handler.FieldHandler
	Reservation handler.ReservationHandler
}

func MainCLI(app *tview.Application, handler Handler) {
	if session.UserSession == nil {
		AuthModal(app, handler)
	} else {
		if session.UserSession.Role == "user" {
			UserDashboardPage(app, handler)
		} else {
			// TODO: admin dasboard ya, semangat!
			OwnerDashboardPage(app, handler)
		}
	}
}
