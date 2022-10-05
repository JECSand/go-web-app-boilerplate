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

	// Account Route
	router.Handler("GET", "/account", p.Protected(ac.AccountPage))
	router.Handler("GET", "/account/:child", p.Protected(ac.AccountPage))

	// Account Settings Route
	router.Handler("POST", "/account/settings", p.Protected(ac.AccountSettingsHandler))

	// Admin Page Routes
	router.Handler("GET", "/admin", p.Protected(ad.AdminPage))
	router.Handler("GET", "/admin/users", p.Protected(ad.AdminPage))
	router.Handler("GET", "/admin/users/:id", p.Protected(ad.AdminUserPage))

	// Admin Group Handler Routes
	router.Handler("GET", "/admin/groups", p.Protected(ad.AdminGroupsPage))
	router.Handler("GET", "/admin/groups/:id", p.Protected(ad.AdminGroupPage))
	router.Handler("POST", "/admin/groups", p.Protected(ad.AdminCreateGroupHandler))
	router.Handler("POST", "/admin/groups/:id/update", p.Protected(ad.AdminUpdateGroupHandler))
	router.Handler("GET", "/admin/groups/:id/delete", p.Protected(ad.AdminDeleteGroupHandler))

	// Admin User Handler Routes
	router.Handler("POST", "/admin/groups/:id", p.Protected(ad.AdminCreateUserHandler))
	//router.Handler("POST", "/admin/users/:id/update", p.Protected(ad.UpdateUserHandler))
	router.Handler("GET", "/admin/users/:id/delete", p.Protected(ad.AdminDeleteUserHandler))

	// Task Route
	router.GET("/tasks", b.VariablePage)
	router.GET("/tasks/:child", b.VariablePage)

	// Example route that encounters an error
	router.GET("/broken/handler", b.BrokenPage)

	// Serve static assets via the "static" directory
	router.ServeFiles("/static/*filepath", p.Viewer.Statics)

	return router
}
