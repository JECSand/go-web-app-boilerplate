package controllers

import (
	"fmt"
	"github.com/JECSand/go-web-app-boilerplate/models"
	"github.com/JECSand/go-web-app-boilerplate/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// AuthController structures the set of app page views
type AuthController struct {
	manager     *ControllerManager
	authService *services.AuthService
}

// RegisterPage renders Index Page
func (p *AuthController) RegisterPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, _ := p.manager.authCheck(r)
	model := models.IndexModel{Name: "home", Title: "Home", Auth: auth}
	if auth.Authenticated {
		p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	rModel := models.IndexModel{Name: "register", Title: "Register", Auth: auth}
	p.manager.Viewer.RenderTemplate(w, "templates/register.html", &rModel)
}

// LoginPage renders the Login Page
func (p *AuthController) LoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, _ := p.manager.authCheck(r)
	fmt.Println("\n\nCHECK AUTH: ", auth)
	model := models.IndexModel{Name: "home", Title: "Home", Auth: auth}
	lModel := models.LoginModel{Name: "login", Title: "Login", Auth: auth}
	//lModel.BuildRoute()
	if auth.Authenticated {
		p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	fmt.Println("\n\nCHECK AUTH 2: ", auth)
	p.manager.Viewer.RenderTemplate(w, "templates/login.html", &lModel)
}

// LoginHandler controls the login process
func (p *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, cookie := p.manager.authCheck(r)
	model := models.LoginModel{Name: "login", Title: "Login", Auth: auth}
	if r.Method != http.MethodPost {
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &model)
		return
	}
	user := &models.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	fmt.Println("\n\nCHECK AUTH REQUEST USER: ", user)
	auth, err := p.authService.Authenticate(user)
	fmt.Println("\n\nCHECK AUTH REPLY: ", auth, err)
	if err != nil || auth.Status != http.StatusOK {
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &model)
		return
	}
	cookie, err = p.manager.SessionManager.NewSession(auth)
	if err != nil {
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &model)
		return
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 303)
}

// RegistrationHandler ...
func (p *AuthController) RegistrationHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, _ := p.manager.authCheck(r)
	model := models.IndexModel{Name: "register", Title: "Register", Auth: auth}
	if r.Method != http.MethodPost {
		p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	user := &models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Username:  r.FormValue("username"),
		Password:  r.FormValue("password"),
	}
	auth, err := p.authService.Register(user)
	if err != nil || auth.Status != http.StatusCreated {
		p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	cookie, err := p.manager.SessionManager.NewSession(auth)
	if err != nil {
		p.manager.Viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 303)
}

// LogoutHandler controls the logout process
func (p *AuthController) LogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here, Build in Check to ensure logout worked
	auth, cookie := p.manager.authCheck(r)
	if auth.Authenticated {
		auth, _ = p.authService.Invalidate(auth)
		_ = p.manager.SessionManager.DeleteSession(cookie)
	}
	http.Redirect(w, r, "/login", 303)
}
