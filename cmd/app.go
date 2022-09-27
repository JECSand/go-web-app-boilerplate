package cmd

import (
	"github.com/JECSand/go-web-app-boilerplate/controllers"
	"github.com/JECSand/go-web-app-boilerplate/server"
	"github.com/JECSand/go-web-app-boilerplate/services"
	"github.com/JECSand/go-web-app-boilerplate/views"
)

// App is the highest level struct of the rest_api application. Stores the server, client, and config settings.
type App struct {
	server *server.Server
}

// Initialize is a function used to initialize a new instantiation of the API Application
func (a *App) Initialize(env string) error {
	config, err := ConfigurationSettings(env)
	if err != nil {
		return err
	}
	config.InitializeEnvironment()
	var globalSessions *services.SessionService
	globalSessions = services.NewSessionService()
	v := views.InitializeViewer()
	v.InitializeTemplates()
	authService := services.NewAuthService()
	uService := services.NewUserService()
	gService := services.NewGroupService()
	manager := controllers.NewManager(v, globalSessions)
	a.server = server.NewServer(manager, authService, uService, gService)
	return nil
}

// Run is a function used to run a previously initialized Application
func (a *App) Run() {
	a.server.Start()
}
