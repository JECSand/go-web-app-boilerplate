package models

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

// Auth models the authentication structure of an app user
type Auth struct {
	Username      string
	UserId        string
	GroupId       string
	Role          string
	AuthToken     string
	APIKey        string
	Authenticated bool
	LastLogin     string
	LoginIP       string
	Status        int
}

// GetAuthString converts the structures attributes into a string (todo - look into using json marshall/unmarshall)
func (a *Auth) GetAuthString() string {
	authStr := "false"
	if a.Authenticated {
		authStr = "true"
	}
	var reString string
	reString = a.Username + "||" + a.UserId + "||" + a.GroupId + "||" + a.Role + "||" + authStr + "||" + a.AuthToken + "||" + a.LastLogin + "||" + a.LoginIP
	return reString
}

// LoadAuthString loads a stringed Auth struct
func (a *Auth) LoadAuthString(authString string) {
	authBool := false
	sString := strings.Split(authString, "||")
	authStr := sString[4]
	if authStr == "true" {
		authBool = true
	}
	a.Username = sString[0]
	a.UserId = sString[1]
	a.GroupId = sString[2]
	a.Role = sString[3]
	a.Authenticated = authBool
	a.AuthToken = sString[5]
	a.LastLogin = sString[6]
	a.LoginIP = sString[7]
}

// Delete all data in Auth struct
func (a *Auth) Delete() {
	a.Username = ""
	a.UserId = ""
	a.GroupId = ""
	a.Role = ""
	a.Authenticated = false
	a.AuthToken = ""
	a.LastLogin = ""
	a.LoginIP = ""
}

// Load Auth with data from a http.Response
func (a *Auth) Load(resp *http.Response) error {
	currentTime := time.Now().UTC()
	authToken := resp.Header["Auth-Token"]
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		return err
	}
	if err = resp.Body.Close(); err != nil {
		return err
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}
	a.Status = resp.StatusCode
	if a.Status == http.StatusOK {
		a.Username = user.Username
		a.UserId = user.Id
		a.GroupId = user.GroupId
		a.Role = user.Role
		a.AuthToken = authToken[0]
		a.LastLogin = currentTime.String()
		a.Authenticated = true
		a.LoginIP = "Unknown"
	}
	return nil
}
