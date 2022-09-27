package models

import (
	"encoding/json"
	"time"
)

// Group is a root struct that is used to store the json encoded data for/from a mongodb group doc.
type Group struct {
	Id           string    `json:"id,omitempty"`
	RootAdmin    bool      `json:"root_admin,omitempty"`
	Name         string    `json:"name,omitempty"`
	LastModified time.Time `json:"last_modified,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

// GetJSON marshals the Group struct data into JSON bytes
func (g *Group) GetJSON() []byte {
	b, _ := json.Marshal(g)
	return b
}

// GetBodyString returns a string version of the Group model for session management
func (g *Group) GetBodyString() string {
	bodyStr := `{"name": "` + g.Name + `"}`
	return bodyStr
}

// GetID returns the Group ID
func (g *Group) GetID() string {
	return g.Id
}

// User is a root struct that is used to store the json encoded data for/from a mongodb user doc.
type User struct {
	Id           string    `json:"id,omitempty"`
	Username     string    `json:"username,omitempty"`
	Password     string    `json:"password,omitempty"`
	FirstName    string    `json:"firstname,omitempty"`
	LastName     string    `json:"lastname,omitempty"`
	Email        string    `json:"email,omitempty"`
	RootAdmin    bool      `json:"root_admin,omitempty"`
	Role         string    `json:"role,omitempty"`
	GroupId      string    `json:"group_id,omitempty"`
	LastModified time.Time `json:"last_modified,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

// GetJSON marshals the User struct data into JSON bytes
func (u *User) GetJSON() []byte {
	b, _ := json.Marshal(u)
	return b
}

// GetBodyString returns a string version of the User model for session management
func (u *User) GetBodyString(strType string) string {
	var bodyStr string
	if strType == "Settings" {
		bodyStr = `{"username": "` + u.Username + `","firstname": "` + u.FirstName + `","lastname":"` + u.LastName + `","email":"` + u.Email + `"}`
	} else if strType == "Admin" {
		bodyStr = `{"username": "` + u.Username + `","firstname": "` + u.FirstName + `","lastname":"` + u.LastName + `","email":"` + u.Email + `","password":"` + u.Password + `","group_id":"` + u.GroupId + `","role":"` + u.Role + `"}`
	}
	return bodyStr
}

// GetID returns the User ID
func (u *User) GetID() string {
	return u.Id
}

// Task is a root struct that is used to store the json encoded data for/from a mongodb todos doc.
type Task struct {
	Id           string    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Completed    bool      `json:"completed,omitempty"`
	Due          time.Time `json:"due,omitempty"`
	Description  string    `json:"description,omitempty"`
	UserId       string    `json:"user_id,omitempty"`
	GroupId      string    `json:"group_id,omitempty"`
	LastModified time.Time `json:"last_modified,omitempty"`
	CreatedAt    time.Time `json:"creation_at,omitempty"`
}

// GetJSON marshals the Group struct data into JSON bytes
func (t *Task) GetJSON() []byte {
	b, _ := json.Marshal(t)
	return b
}

// GetBodyString returns a string version of the Task model for session management
func (t *Task) GetBodyString() string {
	bodyStr := `{"name": "` + t.Name + `"}`
	return bodyStr
}

// GetID returns the Task ID
func (t *Task) GetID() string {
	return t.Id
}
