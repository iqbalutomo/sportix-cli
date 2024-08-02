package cli

import (
	"os"
	"sportix-cli/constants"
	"sportix-cli/internal/entity"
	"sportix-cli/internal/session"

	"github.com/rivo/tview"
)

func AuthModal(app *tview.Application, handler Handler) {
	modal := tview.NewModal().
		SetText("WELCOME to SPORTIX").
		AddButtons([]string{"Register", "Login", "Exit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Register" {
				RegisterPage(app, handler)
			} else if buttonLabel == "Login" {
				LoginPage(app, handler)
			} else if buttonLabel == "Exit" {
				app.Stop()
				os.Exit(0)
			}
		})

	if err := app.SetRoot(modal, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func RegisterPage(app *tview.Application, handler Handler) {
	var user entity.User

	//roleOptions := []string{"user", "owner"}
	selectedRole := constants.User

	form := tview.NewForm().
		AddInputField("Username:", "", 40, nil, func(text string) {
			user.Username = text
		}).
		AddInputField("Email:", "", 40, nil, func(text string) {
			user.Email = text
		}).
		AddPasswordField("Password:", "", 40, '*', func(text string) {
			user.Password = text
		}).
		AddDropDown("Role:", constants.RoleOptions, 0, func(option string, index int) {
			selectedRole = option
			user.Role = selectedRole
		}).
		AddButton("Register", func() {
			if user.Username == "" || user.Email == "" || user.Password == "" || user.Role == "" {
				errorModal := tview.NewModal().
					SetText("Username, email, password, and role cannot be empty").
					AddButtons([]string{"OK"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						RegisterPage(app, handler)
					})
				app.SetRoot(errorModal, true).EnableMouse(true).Run()
				return
			}

			err := handler.User.Register(user.Username, user.Email, user.Password, user.Role)
			if err != nil {
				errorModal := tview.NewModal().
					SetText(err.Error()).
					AddButtons([]string{"OK"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						RegisterPage(app, handler)
					})
				app.SetRoot(errorModal, true).EnableMouse(true).Run()
				return
			}

			successModal := tview.NewModal().
				SetText("Registered successfully").
				AddButtons([]string{"OK"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					LoginPage(app, handler)
				})
			app.SetRoot(successModal, true).EnableMouse(true).Run()
		}).
		AddButton("Back", func() {
			AuthModal(app, handler)
		}).
		AddButton("Exit", func() {
			app.Stop()
			os.Exit(0)
		})

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func LoginPage(app *tview.Application, handler Handler) {
	var user entity.User

	form := tview.NewForm().
		AddInputField("Email:", "", 40, nil, func(text string) {
			user.Email = text
		}).
		AddPasswordField("Password:", "", 40, '*', func(text string) {
			user.Password = text
		}).
		AddButton("Login", func() {
			if user.Email == "" || user.Password == "" {
				errorModal := tview.NewModal().
					SetText("Email and password cannot be empty").
					AddButtons([]string{"OK"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						LoginPage(app, handler)
					})
				app.SetRoot(errorModal, true).EnableMouse(true).Run()
				return
			}

			user, err := handler.User.Login(user.Email, user.Password)
			if err != nil {
				errorModal := tview.NewModal().
					SetText(err.Error()).
					AddButtons([]string{"OK"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						LoginPage(app, handler)
					})
				app.SetRoot(errorModal, true).EnableMouse(true).Run()
				return
			}

			session.UserSession = &entity.CurrentUser{
				UserID:   user.UserID,
				Username: user.Username,
				Email:    user.Email,
				Role:     user.Role,
			}

			successModal := tview.NewModal().
				SetText("Login successfully").
				AddButtons([]string{"OK"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if user.Role == constants.User {
						UserDashboardPage(app, handler)
					} else {
						OwnerDashboardPage(app, handler)
					}
				})
			app.SetRoot(successModal, true).EnableMouse(true).Run()
		}).
		AddButton("Back", func() {
			AuthModal(app, handler)
		}).
		AddButton("Exit", func() {
			app.Stop()
			os.Exit(0)
		})

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
