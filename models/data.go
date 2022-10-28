package models

import (
	"encoding/json"
	"errors"
	"time"
)

// DTOModel is an abstraction of the db model types
type DTOModel interface {
	Validate() error
	GetJSON() []byte
}

// DataModel is an abstraction of the db model types
type DataModel interface {
	GetJSON() []byte
	GetID() string
	GetLabel() string
	GetBoolField(fType string) bool
	GetClass(pl bool) string
}

// DataModels is an abstraction of the db model types
type DataModels interface {
	GetJSON() []byte
	Count() int
}

/*
================ Group Model ==================
*/

// Group is a root struct that is used to store the json encoded data for/from a mongodb group doc.
type Group struct {
	Id           string    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	RootAdmin    bool      `json:"root_admin,omitempty"`
	LastModified time.Time `json:"last_modified,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
}

// Validate checks the data in the DTO for issues
func (g *Group) Validate() error {
	return nil
}

// GetJSON marshals the Group struct data into JSON bytes
func (g *Group) GetJSON() []byte {
	b, _ := json.Marshal(g)
	return b
}

// GetID returns the Group ID
func (g *Group) GetID() string {
	return g.Id
}

// GetLabel returns the Group label
func (g *Group) GetLabel() string {
	return g.Name
}

// GetBoolField returns a bool value stored in Group
func (g *Group) GetBoolField(fType string) bool {
	if fType == "RootAdmin" {
		return g.RootAdmin
	}
	return false
}

// GetClass returns the Class string
func (g *Group) GetClass(pl bool) string {
	if pl {
		return "groups"
	}
	return "group"
}

/*
================ User DTOs ==================
*/

// User is a root struct that is used to store the json encoded data for/from a mongodb user doc.
type User struct {
	Id           string    `json:"id,omitempty"`
	Username     string    `json:"username,omitempty"`
	Password     string    `json:"password,omitempty"`
	FirstName    string    `json:"firstname,omitempty"`
	LastName     string    `json:"lastname,omitempty"`
	Email        string    `json:"email,omitempty"`
	Role         string    `json:"role,omitempty"`
	RootAdmin    bool      `json:"root_admin,omitempty"`
	GroupId      string    `json:"group_id,omitempty"`
	LastModified time.Time `json:"last_modified,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
}

// Validate checks the data in the DTO for issues
func (u *User) Validate() error {
	return nil
}

// GetJSON marshals the User struct data into JSON bytes
func (u *User) GetJSON() []byte {
	b, _ := json.Marshal(u)
	return b
}

// GetID returns the User ID
func (u *User) GetID() string {
	return u.Id
}

// GetLabel returns the User label
func (u *User) GetLabel() string {
	return u.Email
}

// GetBoolField returns a bool value stored in User
func (u *User) GetBoolField(fType string) bool {
	if fType == "RootAdmin" {
		return u.RootAdmin
	}
	return false
}

// GetClass returns the Class string
func (u *User) GetClass(pl bool) string {
	if pl {
		return "users"
	}
	return "user"
}

/*
================ Task DTOs ==================
*/

type TaskStatus string

const (
	NOTSTARTED TaskStatus = "NOT_STARTED"
	INPROGRESS TaskStatus = "IN_PROGRESS"
	COMPLETED  TaskStatus = "COMPLETED"
)

// Task is a root struct that is used to store the json encoded data for/from a mongodb todos doc.
type Task struct {
	Id           string     `json:"id,omitempty"`
	Name         string     `json:"name,omitempty"`
	Status       TaskStatus `json:"status,omitempty"`
	Due          time.Time  `json:"due,omitempty"`
	Description  string     `json:"description,omitempty"`
	UserId       string     `json:"user_id,omitempty"`
	GroupId      string     `json:"group_id,omitempty"`
	LastModified time.Time  `json:"last_modified,omitempty"`
	CreatedAt    time.Time  `json:"creation_at,omitempty"`
	DeletedAt    time.Time  `json:"deleted_at,omitempty"`
}

// Validate checks the data in the DTO for issues
func (t *Task) Validate() error {
	return nil
}

// GetJSON marshals the Group struct data into JSON bytes
func (t *Task) GetJSON() []byte {
	b, _ := json.Marshal(t)
	return b
}

// GetID returns the Task ID
func (t *Task) GetID() string {
	return t.Id
}

// GetLabel returns the Task label
func (t *Task) GetLabel() string {
	return t.Name
}

// GetBoolField returns a bool value stored in Task
func (t *Task) GetBoolField(fType string) bool {
	if fType == "Completed" && t.Status == COMPLETED {
		return true
	}
	return false
}

// GetClass returns the Class string
func (t *Task) GetClass(pl bool) string {
	if pl {
		return "tasks"
	}
	return "task"
}

/*
================ User DTOs ==================
*/

// Users structures a slice of User
type Users struct {
	Items []*User `json:"users"`
}

// GetJSON checks the data in the DTO for issues
func (d *Users) GetJSON() []byte {
	b, _ := json.Marshal(d)
	return b
}

// Validate checks the data in the DTO for issues
func (d *Users) Validate() error {
	if len(d.Items) == 0 {
		return errors.New("empty")
	}
	return nil
}

// Count checks the data in the DTO for issues
func (d *Users) Count() int {
	return len(d.Items)
}

/*
================ Group DTOs ==================
*/

// Groups is used when returning a slice of Group
type Groups struct {
	Items []*Group `json:"groups"`
}

// GetJSON checks the data in the DTO for issues
func (d *Groups) GetJSON() []byte {
	b, _ := json.Marshal(d)
	return b
}

// Validate checks the data in the DTO for issues
func (d *Groups) Validate() error {
	if len(d.Items) == 0 {
		return errors.New("empty")
	}
	return nil
}

// Count checks the data in the DTO for issues
func (d *Groups) Count() int {
	return len(d.Items)
}

// GroupUsersDTO is used when returning a group with its associated users
type GroupUsersDTO struct {
	Group *Group  `json:"group"`
	Users []*User `json:"users"`
}

// GetJSON checks the data in the DTO for issues
func (d *GroupUsersDTO) GetJSON() []byte {
	b, _ := json.Marshal(d)
	return b
}

// Validate checks the data in the DTO for issues
func (d *GroupUsersDTO) Validate() error {
	if len(d.Users) == 0 {
		return errors.New("empty")
	}
	return nil
}

// Count checks the data in the DTO for issues
func (d *GroupUsersDTO) Count() int {
	return len(d.Users)
}

/*
================ Task DTOs ==================
*/

// Tasks is used when returning a slice of Task
type Tasks struct {
	Items []*Task `json:"tasks"`
}

// GetJSON checks the data in the DTO for issues
func (d *Tasks) GetJSON() []byte {
	b, _ := json.Marshal(d)
	return b
}

// Validate checks the data in the DTO for issues
func (d *Tasks) Validate() error {
	if len(d.Items) == 0 {
		return errors.New("empty")
	}
	return nil
}

// Count checks the data in the DTO for issues
func (d *Tasks) Count() int {
	return len(d.Items)
}

/*
================ Auth DTOs ==================
*/

// UpdatePassword is used when updating a user password
type UpdatePassword struct {
	NewPassword     string `json:"new_password"`
	CurrentPassword string `json:"current_password"`
}

// GetJSON checks the data in the DTO for issues
func (d *UpdatePassword) GetJSON() []byte {
	b, _ := json.Marshal(d)
	return b
}

// Validate checks the data in the DTO for issues
func (d *UpdatePassword) Validate() error {
	if d.CurrentPassword == d.NewPassword {
		return errors.New("passwords cannot match")
	}
	return nil
}
