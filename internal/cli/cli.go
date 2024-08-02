package cli

import (
	"sportix-cli/internal/entity"
	"sportix-cli/internal/handler"
	"sportix-cli/internal/session"

	"github.com/rivo/tview"
)

type Handler struct {
	User     handler.UserHandler
	Location handler.LocationHandler
	Category handler.CategoryHandler
	Field    handler.FieldHandler
}

func MainCLI(app *tview.Application, handler Handler) {
	user, err := handler.User.Login("juragan@lapangan.com", "1234")
	if err != nil {
		return
	}
	saldo, err := handler.User.GetBalanceByEmail(user.Email)
	if err != nil {
		return
	}
	currentUser := &entity.CurrentUser{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Balance:  saldo,
	}
	session.UserSession = currentUser

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
