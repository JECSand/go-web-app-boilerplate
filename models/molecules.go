package models

/*
Forms
*/

// Form structures a generic form
type Form struct {
	Name     string        `json:"name,omitempty"`
	Type     string        `json:"type,omitempty"`
	Class    string        `json:"class,omitempty"`
	Id       string        `json:"id,omitempty"`
	Fields   []*InputField `json:"fields,omitempty"`
	Buttons  []*Button     `json:"buttons,omitempty"`
	Popup    *Button       `json:"popup,omitempty"`
	Method   string        `json:"method,omitempty"`
	Action   string        `json:"action,omitempty"`
	Category string        `json:"category,omitempty"`
}

// InitializeForm for a new form
func InitializeForm(formMeta []string, fields []*InputField, buttons []*Button, popup *Button) *Form {
	newForm := Form{}
	// Field Vector String Array, this is order
	// Name, Class, Id, Type, Label, DefaultVal
	newForm.Name = formMeta[0]
	newForm.Type = formMeta[1]
	newForm.Class = formMeta[2]
	newForm.Id = formMeta[3]
	newForm.Method = formMeta[4]
	newForm.Action = formMeta[5]
	newForm.Fields = fields
	newForm.Buttons = buttons
	newForm.Category = formMeta[6]
	if popup != nil {
		newForm.Popup = popup
	}
	return &newForm
}

// InitializeSignInForm for a settings form
func InitializeSignInForm() *Form {
	// Field Vector String Array, this is order
	// NAME, TYPE, CLASS, ID, METHOD, ACTION
	formMeta := []string{"Sign In", "Auth", "form1", "form1", "POST", "/auth", "default"}
	// Name, Class, Id, Type, Label, DefaultVal
	emailField := NewInput("Email", "Email", "update", "email", "text", "")
	pwField := NewInput("Password", "Password", "update", "password", "password", "")
	fields := []*InputField{emailField, pwField}
	button := &Button{Name: "update", Class: "form1", Id: "form1", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return InitializeForm(formMeta, fields, buttons, nil)
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
	button := &Button{Name: "update", Class: "form1", Id: "form1", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return InitializeForm(formMeta, fields, buttons, nil)
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
	return InitializeForm(formMeta, fields, buttons, nil)
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
	return InitializeForm(formMeta, fields, buttons, opButton)
}
