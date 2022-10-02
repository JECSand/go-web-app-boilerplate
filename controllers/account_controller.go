package controllers

import (
	"fmt"
	"github.com/JECSand/go-web-app-boilerplate/models"
	"github.com/JECSand/go-web-app-boilerplate/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// AccountController structures the set of app page views
type AccountController struct {
	manager     *ControllerManager
	userService *services.UserService
}

// AccountPage renders the Account Page
func (p *AccountController) AccountPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	fmt.Println("VarsTEST:", params.ByName("child"))
	model := models.AccountModel{Title: "Account", SubRoute: params.ByName("child"), Name: "account", Auth: auth}
	if model.SubRoute == "settings" {
		user, err := p.userService.Get(auth, &models.User{Id: auth.UserId})
		if err != nil {
			panic(err)
		}
		model.User = user
		model.Settings = models.InitializeUserSettings(model.User)
	}
	model.BuildRoute()
	if !auth.Authenticated {
		lModel := models.LoginModel{Title: "Login", Name: "login", Auth: auth, Heading: models.NewHeading("Login", "w3-wide text")}
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &lModel)
		return
	}
	model.Heading = models.NewHeading("My Account", "w3-wide text")
	if model.SubRoute == "settings" {
		model.Title = "Account Settings"
		model.Name = "Account Settings"
		model.Heading = models.NewHeading("Account Settings", "w3-wide text")
	}
	p.manager.Viewer.RenderTemplate(w, "templates/account.html", &model)
}

// AccountSettingsHandler controls the account settings update process
func (p *AccountController) AccountSettingsHandler(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.manager.authCheck(r)
	lModel := models.LoginModel{Name: "Login", Title: "Login", Auth: auth}
	if r.Method != http.MethodPatch {
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &lModel)
		return
	}
	user := &models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Username:  r.FormValue("username"),
	}
	user, err := p.userService.Update(auth, user)
	model := models.AccountModel{Name: "account", Title: "Account Settings", SubRoute: "settings", Auth: auth}
	model.BuildRoute()
	if err != nil {
		p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	model.Settings = models.InitializeUserSettings(model.User)
	p.manager.Viewer.RenderTemplate(w, "templates/account.html", &model)
}
