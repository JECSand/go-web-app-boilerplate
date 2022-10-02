package models

import "fmt"

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

// InitializeUserSettingsForm for a settings form
func InitializeUserSettingsForm(user *User) *Form {
	// Field Vector String Array, this is order
	// NAME, TYPE, CLASS, ID, METHOD, ACTION
	formMeta := []string{"Update", "User", "form1", "updateUser", "PATCH", "", "default"}
	// Name, Class, Id, Type, Label, DefaultVal
	fmt.Println("\n\nCHECK THIS USER BRO: ", user)
	unField := NewInput("User Name", "User Name", "update", "username", "text", user.Username)
	fnField := NewInput("First Name", "First Name", "update", "first_name", "text", user.FirstName)
	lnField := NewInput("Last Name", "Last Name", "update", "last_name", "text", user.LastName)
	emailField := NewInput("Email", "Email", "update", "email", "text", user.Email)
	fields := []*InputField{unField, fnField, lnField, emailField}
	button := &Button{Name: "update", Class: "btn", Id: "updateUser", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return NewForm(formMeta, fields, buttons, nil)
}

// InitializeGroupSettingsForm for a settings form
func InitializeGroupSettingsForm(group *Group) *Form {
	// Field Vector String Array, this is order
	// NAME, TYPE, CLASS, ID, METHOD, ACTION
	formMeta := []string{"Update", "Group", "form1", "updateGroup", "PATCH", "", "default"}
	// Name, Class, Id, Type, Label, DefaultVal
	nameField := NewInput("Group Name", "Group Name", "update", "name", "text", group.Name)
	fields := []*InputField{nameField}
	button := &Button{Name: "update", Class: "btn", Id: "updateGroup", Type: "submit", Label: "Submit", Category: "form"}
	buttons := []*Button{button}
	return NewForm(formMeta, fields, buttons, nil)
}

// InitializePopupCreateUserForm for a settings form
func InitializePopupCreateUserForm(availGroups []*Group, setRole bool) *Form {
	formAction := "/admin/users"
	formMeta := []string{"Create", "User", "form-container", "createUser", "POST"}
	// Name, Class, Id, Type, Label, DefaultVal
	unField := NewInput("User Name", "User Name", "update", "username", "text", "")
	pwField := NewInput("Password", "Password", "update", "password", "password", "")
	cpwField := NewInput("Password", "Confirm Password", "password", "cpassword", "password", "")
	fnField := NewInput("First Name", "First Name", "update", "first_name", "text", "")
	lnField := NewInput("Last Name", "Last Name", "update", "last_name", "text", "")
	emailField := NewInput("Email", "Email", "update", "email", "text", "")
	fields := []*InputField{unField, pwField, cpwField, fnField, lnField, emailField}
	if len(availGroups) > 1 {
		groupField := NewSelectInput("User Group", "User Group", "update", "group_id", "text", GetGroupSelectOptions(availGroups), false)
		fields = append(fields, groupField)
	} else if len(availGroups) == 1 {
		formAction = "/admin/groups/" + availGroups[0].Id
		groupField := NewInput("User Group", "", "update", "group_id", "hidden", availGroups[0].Id)
		fields = append(fields, groupField)
	}
	if setRole {
		groupField := NewSelectInput("User Role", "User Role", "update", "role", "text", GetRoleSelectOptions(), false)
		fields = append(fields, groupField)
	}
	subButton := NewButtonInput("Submit", "", "btn", "", "submit", "Submit")
	fields = append(fields, subButton)
	clsButton := &Button{Name: "update", Class: "btn cancel", Id: "createUser", Type: "button", Label: "Close", Category: "close-click"}
	opButton := &Button{Name: "update", Class: "open-button", Id: "createUser", Type: "button", Label: "Create User", Category: "open-click"}
	buttons := []*Button{clsButton}
	formMeta = append(formMeta, formAction)
	formMeta = append(formMeta, "popup")
	return NewForm(formMeta, fields, buttons, opButton)
}

// InitializePopupCreatGroupForm for a settings form
func InitializePopupCreatGroupForm() *Form {
	// Field Vector String Array, this is order
	// NAME, TYPE, CLASS, ID, METHOD, ACTION, CATEGORY
	formMeta := []string{"Create", "Group", "form-container", "createGroup", "POST", "/admin/groups", "popup"}
	// Name, Class, Id, Type, Label, DefaultVal
	unField := NewInput("Group Name", "Group Name", "update", "name", "text", "")
	subButton := NewButtonInput("Submit", "", "btn", "", "submit", "Submit")
	fields := []*InputField{unField, subButton}
	clsButton := &Button{Name: "update", Class: "btn cancel", Id: "createGroup", Type: "button", Label: "Close", Category: "close-click"}
	opButton := &Button{Name: "update", Class: "open-button", Id: "createGroup", Type: "button", Label: "Create Group", Category: "open-click"}
	buttons := []*Button{clsButton}
	return NewForm(formMeta, fields, buttons, opButton)
}
