package server

import (
	"github.com/JECSand/go-web-app-boilerplate/controllers"
	"github.com/julienschmidt/httprouter"
)

// GetRouter returns a new HTTP Router
func GetRouter(p *controllers.ControllerManager, b *controllers.BasicController, au *controllers.AuthController, ac *controllers.AccountController, ad *controllers.AdminController) *httprouter.Router {

	// mux handler
	router := httprouter.New()

	// Index route
	router.GET("/", b.IndexPage)
	// TODO Change "/" to "/register" and Make index (/) a slash page
	router.GET("/register", au.RegisterPage)
	// Register new Account/User at Index Page Form
	router.POST("/register", au.RegistrationHandler)

	// Login Route
	router.GET("/login", au.LoginPage)
	router.POST("/login", au.LoginHandler)

	// Logout Route
	router.GET("/logout", au.LogoutHandler)

	// About Route
	router.GET("/about", b.AboutPage)
	router.GET("/about/:child", b.AboutPage)

	// Admin Page Routes
	router.Handler("GET", "/admin", p.Protected(ad.AdminPage)) // 1) GET GENERIC ADMIN VIEW

	router.Handler("GET", "/admin/:child", p.Protected(ad.AdminPage)) // 2) GET ADMIN USER OR GROUP DATA TABLE
	// ID CAN BE EITHER A UUID OR CREATE - CREATE LOADS CREATE FORM
	router.Handler("GET", "/admin/:child/:id", p.Protected(ad.AdminPage)) // 3) GET ADMIN USER OR GROUP UPDATE FORM
	// Admin Group Handler Routes
	//router.Handler("POST", "/admin/:child", p.Protected(ad.CreateHandler))
	// TO LOAD AN UPDATE FORM SPECIFICALLY WHEN APP USER CLICKS UPDATE FOR A GROUP LISTED IN GROUP DATATABLE
	// HANDLERS TO SUBMIT UPDATE FORM OR DELETE A GROUP
	//router.Handler("PATCH", "/admin/:child/:id", p.Protected(ad.UpdateHandler))
	//router.Handler("DELETE", "/admin/:child/:id", p.Protected(ad.DeleteHandler))

	//router.Handler("PATCH", "/admin/usermenu", p.Protected(p.AdminPage)) // 2) GET ADMIN USER OR GROUP DATA TABLE

	// Admin User Handler Routes
	//router.Handler("POST", "/admin/users", p.Protected(p.CreateUserHandler))
	// TO LOAD AN UPDATE FORM SPECIFICALLY WHEN APP USER CLICKS UPDATE FOR A USER LISTED IN USER DATATABLE
	//router.Handler("GET", "/admin/users/:id", p.Protected(p.AdminPage))
	// HANDLERS TO SUBMIT UPDATE FORM OR DELETE A USER
	//router.Handler("PATCH", "/admin/users/:id", p.Protected(p.UpdateUserHandler))
	//router.Handler("DELETE", "/admin/users/:id", p.Protected(p.DeleteUserHandler))

	// Account Route
	router.Handler("GET", "/account", p.Protected(ac.AccountPage))
	router.Handler("GET", "/account/:child", p.Protected(ac.AccountPage))

	// Account Settings Route
	router.Handler("PATCH", "/account/settings", p.Protected(ac.AccountSettingsHandler))

	// Variable Route
	router.GET("/variable", b.VariablePage)
	router.GET("/variable/:child", b.VariablePage)

	// Example route that encounters an error
	router.GET("/broken/handler", b.BrokenPage)

	// Serve static assets via the "static" directory
	router.ServeFiles("/static/*filepath", p.Viewer.Statics)

	return router
}
