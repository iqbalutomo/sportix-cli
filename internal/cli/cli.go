package cli

import (
	"sportix-cli/internal/entity"
	"sportix-cli/internal/handler"
	"sportix-cli/internal/session"

	"github.com/rivo/tview"
)

type Handler struct {
	User  handler.UserHandler
	Field handler.FieldHandler
}

func MainCLI(app *tview.Application, handler Handler) {
	user, err := handler.User.Login("chloe@mail.com", "password123")
	if err != nil {
		return
	}
	currentUser := &entity.CurrentUser{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
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
