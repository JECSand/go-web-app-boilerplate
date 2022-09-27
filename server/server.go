package server

import (
	"flag"
	"fmt"
	"github.com/JECSand/go-web-app-boilerplate/controllers"
	"github.com/JECSand/go-web-app-boilerplate/models"
	"github.com/JECSand/go-web-app-boilerplate/services"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

// Server is a struct that stores the API Apps high level attributes such as the router, config, and services
type Server struct {
	manager *controllers.ControllerManager
	Router  *httprouter.Router
}

// NewServer is a function used to initialize a new Server struct
func NewServer(manager *controllers.ControllerManager, auth *services.AuthService, u *services.APIService[*models.User], g *services.APIService[*models.Group]) *Server {
	s := Server{manager: manager}
	basicController := s.manager.NewBasicController()
	authController := s.manager.NewAuthController(auth)
	accountController := s.manager.NewAccountController(u)
	adminController := s.manager.NewAdminController(u, g)
	s.Router = GetRouter(manager, basicController, authController, accountController, adminController)
	return &s
}

// Start starts the initialized server
func (s *Server) Start() {
	port := ":" + os.Getenv("PORT")
	listen := flag.String("listen ", port, "Interface and port to listen on")
	flag.Parse()
	fmt.Println("Listening on ", *listen)
	log.Fatal(http.ListenAndServe(*listen, s.Router))
}
