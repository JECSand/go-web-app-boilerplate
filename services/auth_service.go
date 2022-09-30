package services

import (
	"fmt"
	"github.com/JECSand/fetch"
	"github.com/JECSand/go-web-app-boilerplate/models"
	"os"
)

// AuthService manages the authentication system of the app
type AuthService struct {
	host string
}

// AuthHeaders returns authenticated headers
func (a *AuthService) AuthHeaders(auth *models.Auth) [][]string {
	reqHeaders := fetch.JSONDefaultHeaders()
	if auth.AuthToken != "" {
		return fetch.AppendHeaders(reqHeaders, []string{"Auth-Token", auth.AuthToken})
	}
	return reqHeaders
}

// authenticate a user session
func (a *AuthService) authenticate(body []byte, authType string) (*fetch.Fetch, error) {
	url := a.host + "/auth"
	if authType == "registration" {
		url = a.host + "/auth/register"
	}
	newReq := NewRequest(url, "")
	return newReq.Post(body)
}

// invalidate a user session
func (a *AuthService) invalidate(authToken string) (*fetch.Fetch, error) {
	newReq := NewRequest(a.host+"/auth", authToken)
	return newReq.Delete()
}

// Register a new User
func (a *AuthService) Register(user *models.User) (*models.Auth, error) {
	var auth models.Auth
	jsonBytes := user.GetJSON()
	f, err := a.authenticate(jsonBytes, "registration")
	if err != nil {
		return nil, err
	}
	f.Resolve()
	err = auth.Load(f.Res)
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

// Authenticate a User
func (a *AuthService) Authenticate(user *models.User) (*models.Auth, error) {
	var auth models.Auth
	bodyStr := `{"email":"` + user.Email + `","password":"` + user.Password + `"}`
	fmt.Println("\nCHECK BODY STRING: ", bodyStr)
	f, err := a.authenticate([]byte(bodyStr), "login")
	fmt.Println("\nCHECK AUTH ERROR: ", err)
	if err != nil {
		return nil, err
	}
	f.Resolve()
	fmt.Println("\nCHECK AUTH res: ", f.Res)
	err = auth.Load(f.Res)
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

// Invalidate a User session
func (a *AuthService) Invalidate(auth *models.Auth) (*models.Auth, error) {
	f, err := a.invalidate(auth.AuthToken)
	if err != nil {
		return auth, err
	}
	f.Resolve()
	auth.Delete()
	return auth, nil
}

// NewAuthService initializes a new AuthService
func NewAuthService() *AuthService {
	return &AuthService{host: os.Getenv("API_HOST")}
}
