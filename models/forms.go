package models

import "net/http"

// InitializeSignInForm for a settings form
func InitializeSignInForm() *Form {
	// Field Vector String Array, this is order
	// NAME, TYPE, CLASS, ID, METHOD, ACTION
	formMeta := []string{"Sign In", "Auth", "form1", "form1", "POST", "/auth", "default"}
	// Name, Class, Id, Type, Label, DefaultVal
	emailField := NewInput("Email", "Email", "update", "email", "text", "")
	pwField := NewInput("Password", "Password", "update", "password", "password", "")
	fields := []*InputField{emailField, pwField}
	button := &Button{Name: "update", Class: "btn", Id: "form1", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return NewForm(formMeta, fields, buttons, nil)
}

// InitializeRegistrationForm for a settings form
func InitializeRegistrationForm() *Form {
	// Field Vector String Array, this is order
	// NAME, TYPE, CLASS, ID, METHOD, ACTION
	formMeta := []string{"Update", "User", "form1", "form1", "POST", "", "default"}
	// Name, Class, Id, Type, Label, DefaultVal
	unField := NewInput("User Name", "User Name", "update", "username", "text", "")
	pwField := NewInput("Password", "Password", "update", "password", "password", "")
	cpwField := NewInput("Password", "Confirm Password", "password", "cpassword", "password", "")
	fnField := NewInput("First Name", "First Name", "update", "first_name", "text", "")
	lnField := NewInput("Last Name", "Last Name", "update", "last_name", "text", "")
	emailField := NewInput("Email", "Email", "update", "email", "text", "")
	fields := []*InputField{unField, pwField, cpwField, fnField, lnField, emailField}
	button := &Button{Name: "update", Class: "btn", Id: "form1", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return NewForm(formMeta, fields, buttons, nil)
}

// InitializeSettingsForm for a settings form
func InitializeSettingsForm(user *User) *Form {
	// Field Vector String Array, this is order
	// NAME, TYPE, CLASS, ID, METHOD, ACTION
	formMeta := []string{"Update", "User", "form1", "form2", "PATCH", "", "default"}
	// Name, Class, Id, Type, Label, DefaultVal
	unField := NewInput("User Name", "User Name", "update", "username", "text", user.Username)
	fnField := NewInput("First Name", "First Name", "update", "first_name", "text", user.FirstName)
	lnField := NewInput("Last Name", "Last Name", "update", "last_name", "text", user.LastName)
	emailField := NewInput("Email", "Email", "update", "email", "text", user.Email)
	fields := []*InputField{unField, fnField, lnField, emailField}
	button := &Button{Name: "update", Class: "form1", Id: "form2", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return NewForm(formMeta, fields, buttons, nil)
}

// InitializePopupCreatUserForm for a settings form
func InitializePopupCreatUserForm() *Form {
	// Field Vector String Array, this is order
	// NAME, TYPE, CLASS, ID, METHOD, ACTION, CATEGORY
	formMeta := []string{"Update", "User", "form-container", "myForm", "POST", "/users", "popup"}
	// Name, Class, Id, Type, Label, DefaultVal
	unField := NewInput("User Name", "User Name", "update", "username", "text", "")
	pwField := NewInput("Password", "Password", "update", "password", "password", "")
	cpwField := NewInput("Password", "Confirm Password", "password", "cpassword", "password", "")
	fnField := NewInput("First Name", "First Name", "update", "first_name", "text", "")
	lnField := NewInput("Last Name", "Last Name", "update", "last_name", "text", "")
	emailField := NewInput("Email", "Email", "update", "email", "text", "")
	subButton := NewButtonInput("Submit", "", "btn", "", "submit", "Submit")
	fields := []*InputField{unField, pwField, cpwField, fnField, lnField, emailField, subButton}
	clsButton := &Button{Name: "update", Class: "btn cancel", Id: "myForm", Type: "button", Label: "Close", Category: "close-click"}
	opButton := &Button{Name: "update", Class: "open-button", Id: "myForm", Type: "button", Label: "Create User", Category: "open-click"}
	buttons := []*Button{clsButton}
	return NewForm(formMeta, fields, buttons, opButton)
}

/*
========================================= OLD =========================================
*/

// CreateUserForm ...
type CreateUserForm struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	UserName  string `json:"username,omitempty"`
	Role      string `json:"role,omitempty"`
	GroupUuid string `json:"groupuuid,omitempty"`
	Password  string `json:"password,omitempty"`
	CPassword string `json:"cpassword,omitempty"`
}

// LoadRequest ...
func (cuf *CreateUserForm) LoadRequest(r *http.Request) {
	cuf.FirstName = r.FormValue("first_name")
	cuf.LastName = r.FormValue("last_name")
	cuf.Email = r.FormValue("email")
	cuf.UserName = r.FormValue("username")
	cuf.Role = r.FormValue("role")
	cuf.GroupUuid = r.FormValue("group_uuid")
	cuf.Password = r.FormValue("password")
}

/*
// Create - Update user for info display
func (cuf *CreateUserForm) Create(auth *Auth) int {
	statusCode := 201
	// TODO Add User info to user struct below
	user := User{}
	// Other stuff maybe?
	user.Create(auth)
	// TODO - Capture response status and render a success or error message
	return statusCode
}
*/

// UpdateUserForm ...
type UpdateUserForm struct {
	Uuid      string `json:"uuid,omitempty"` // HIDDEN FIELD IN HTML MODEL
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	UserName  string `json:"username,omitempty"`
	Role      string `json:"role,omitempty"`
	GroupUuid string `json:"groupuuid,omitempty"`
	Password  string `json:"password,omitempty"`
	CPassword string `json:"cpassword,omitempty"`
	Form      Form   `json:"form,omitempty"`
}

/*
// Initialize ...
func (uuf *UpdateUserForm) Initialize(uuid string) {
	// Field Vector String Array, this is order
	// NAME, TYPE, CLASS, ID, METHOD, ACTION
	formMeta := []string{"Update", "User", "form1", "form2", "PATCH", ""}
	// Name, Class, Id, Type, Label, DefaultVal
	initStr := []string{"uuid", "update", "uuid", "text", "Uuid", uuid}
	fNameStr := []string{"first_name", "update", "name", "text", "First Name", ""}
	lNameStr := []string{"last_name", "update", "name", "text", "Last Name", ""}
	fieldStrs := [][]string{initStr, fNameStr, lNameStr}
	form := InitializeForm(formMeta, fieldStrs, Button{Name: "update", Class: "form1", Id: "form2", Type: "submit", Label: "Submit"})
	uuf.Form = form
}

*/

// LoadRequest ...
func (uuf *UpdateUserForm) LoadRequest(r *http.Request) {
	uuf.Uuid = r.FormValue("id")
	uuf.FirstName = r.FormValue("first_name")
	uuf.LastName = r.FormValue("last_name")
	uuf.Email = r.FormValue("email")
	uuf.UserName = r.FormValue("username")
	uuf.Role = r.FormValue("role")
	uuf.GroupUuid = r.FormValue("group_uuid")
	uuf.Password = r.FormValue("password")
}

// Load UpdateUserForm when AccountModel is Initialized
func (uuf *UpdateUserForm) Load(user *User) {
	uuf.Uuid = user.Id
	uuf.UserName = user.Username
	uuf.FirstName = user.FirstName
	uuf.LastName = user.LastName
	uuf.Email = user.Email
	uuf.GroupUuid = user.GroupId
	uuf.Role = user.Role
}

/*
// Update - Update user for info display
func (uuf *UpdateUserForm) Update(auth *Auth, uuid string) int {
	statusCode := 200
	if uuid == "" {
		uuid = auth.UserId
	}
	user := User{
		Id:        uuid,
		Username:  uuf.UserName,
		FirstName: uuf.FirstName,
		LastName:  uuf.LastName,
		Email:     uuf.Email,
		GroupId:   uuf.GroupUuid,
		Role:      uuf.Role,
		Password:  uuf.Password,
	}
	user.Update(auth, "Admin")
	// TODO - Capture response status and render a success or error message
	return statusCode
}
*/

// CreateGroupForm ...
type CreateGroupForm struct {
	Name string `json:"name,omitempty"`
}

// LoadRequest ...
func (cgf *CreateGroupForm) LoadRequest(r *http.Request) {
	cgf.Name = r.FormValue("name")
}

/*
// Create - Update user for info display
func (cgf *CreateGroupForm) Create(auth *Auth) int {
	statusCode := 201
	group := Group{
		Name: cgf.Name,
	}
	group.Create(auth)
	// TODO - Capture response status and render a success or error message
	return statusCode
}
*/

// UpdateGroupForm ...
type UpdateGroupForm struct {
	Uuid string `json:"uuid,omitempty"` // HIDDEN FIELD IN HTML MODEL
	Name string `json:"name,omitempty"`
}

// LoadRequest ...
func (ugf *UpdateGroupForm) LoadRequest(r *http.Request) {
	ugf.Uuid = r.FormValue("uuid")
	ugf.Name = r.FormValue("name")
}

// Load CreateGroupForm when AccountModel is Initialized
func (ugf *UpdateGroupForm) Load(group *Group) {
	ugf.Name = group.Name
}

/*
// Update - Update user for info display
func (ugf *UpdateGroupForm) Update(auth *Auth, uuid string) int {
	statusCode := 200
	if uuid == "" {
		uuid = auth.GroupId
	}
	group := Group{
		Id:   uuid,
		Name: ugf.Name,
	}
	group.Update(auth)
	// TODO - Capture response status and render a success or error message
	return statusCode
}
*/

// AccountSettingsForm ...
type AccountSettingsForm struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	UserName  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	CPassword string `json:"cpassword,omitempty"`
}

/*
// Update Account Settings
func (rm *AccountSettingsForm) Update() *Auth {
	auth := Auth{}
	auth.Register(rm.FirstName, rm.LastName, rm.Email, rm.UserName, rm.Password)
	return auth
}
*/

// AdminUsermenuForm ...
type AdminUsermenuForm struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	UserName  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	CPassword string `json:"cpassword,omitempty"`
}

/*
// Update admin Settings
func (aum *AdminUsermenuForm) Update() *Auth {
	auth := Auth{}
	auth.Register(aum.FirstName, aum.LastName, aum.Email, aum.UserName, aum.Password)
	return auth
}
*/

// UpdatePasswordForm ...
type UpdatePasswordForm struct {
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"newpassword,omitempty"`
	CPassword   string `json:"cpassword,omitempty"`
}

/* Change Account Settings
func (rm *UpdatePasswordForm) Update() Auth {
	auth := Auth{}
	auth.Register(rm.Password, rm.NewPassword)
	return auth
}
*/

// SettingsForm ...
type SettingsForm struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
}

// Load SettingsForm when AccountModel is Initialized
func (sm *SettingsForm) Load(user *User) {
	sm.Username = user.Username
	sm.FirstName = user.FirstName
	sm.LastName = user.LastName
	sm.Email = user.Email
}

/*
// UpdateSettings - Update user for info display
func (sm *SettingsForm) UpdateSettings(auth *Auth) int {
	statusCode := 200
	user := User{
		Id:        auth.UserId,
		Username:  sm.Username,
		FirstName: sm.FirstName,
		LastName:  sm.LastName,
		Email:     sm.Email,
	}
	user.Update(auth, "Settings")
	// TODO - Capture response status and render a success or error message
	return statusCode
}
*/
