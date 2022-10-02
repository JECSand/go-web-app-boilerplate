package models

/*
Heading Types
*/

// Heading ...
type Heading struct {
	Class    string
	Id       string
	Label    string
	Category string
}

// NewHeading ...
func NewHeading(label string, class string) *Heading {
	return &Heading{Class: class, Label: label, Category: "default"}
}

// NewColumnHeading ...
func NewColumnHeading(label string, class string) *Heading {
	return &Heading{Class: class, Label: label, Category: "column"}
}

/*
Alert Types
*/

// Alert ...
type Alert struct {
	Message    string
	ClickClose bool
	Category   string
}

// NewAlert constructs and returns a new Link
func NewAlert(message string, cc bool) *Alert {
	return &Alert{
		Message:    message,
		ClickClose: cc,
		Category:   "default",
	}
}

// NewSuccessAlert constructs and returns a new Link
func NewSuccessAlert(message string, cc bool) *Alert {
	return &Alert{
		Message:    message,
		ClickClose: cc,
		Category:   "success",
	}
}

// NewErrorAlert constructs and returns a new Link
func NewErrorAlert(message string, cc bool) *Alert {
	return &Alert{
		Message:    message,
		ClickClose: cc,
		Category:   "error",
	}
}

/*
Link Types
*/

// Link ...
type Link struct {
	Class    string
	Id       string
	Ref      string
	Label    string
	Break    bool
	Category string
}

// NewLink constructs and returns a new Link
func NewLink(class string, id string, ref string, label string, br bool) *Link {
	return &Link{
		Class:    class,
		Id:       id,
		Ref:      ref,
		Label:    label,
		Break:    br,
		Category: "default",
	}
}

/*
Button Types
*/

// Button ...
type Button struct {
	Name     string
	Class    string
	Id       string
	Type     string
	Label    string
	Category string
}

/*
Data Field Types
*/

// SelectOptions ...
type SelectOptions struct {
	Value    string
	Label    string
	Selected bool
}

// GetGroupSelectOptions ...
func GetGroupSelectOptions(groups []*Group) []*SelectOptions {
	var ops []*SelectOptions
	for _, g := range groups {
		ops = append(ops, &SelectOptions{Value: g.Id, Label: g.Name})
	}
	return ops
}

// GetRoleSelectOptions ...
func GetRoleSelectOptions() []*SelectOptions {
	return []*SelectOptions{{Value: "admin", Label: "Admin"}, {Value: "member", Label: "Member"}}
}

// InputField ...
type InputField struct {
	Name     string
	Label    string
	Class    string
	Id       string
	Type     string
	Value    string
	Multi    bool
	Options  []*SelectOptions
	Rows     string
	Cols     string
	Category string
}

// NewInput ...
func NewInput(name string, label string, class string, id string, iType string, val string) *InputField {
	return &InputField{
		Name:     name,
		Label:    label,
		Class:    class,
		Id:       id,
		Type:     iType,
		Value:    val,
		Multi:    false,
		Options:  nil,
		Rows:     "",
		Cols:     "",
		Category: "input",
	}
}

// NewButtonInput ...
func NewButtonInput(name string, label string, class string, id string, iType string, val string) *InputField {
	return &InputField{
		Name:     name,
		Label:    label,
		Class:    class,
		Id:       id,
		Type:     iType,
		Value:    val,
		Multi:    false,
		Options:  nil,
		Rows:     "",
		Cols:     "",
		Category: "submit",
	}
}

// NewSelectInput ...
func NewSelectInput(name string, label string, class string, id string, iType string, ops []*SelectOptions, m bool) *InputField {
	return &InputField{
		Name:     name,
		Label:    label,
		Class:    class,
		Id:       id,
		Type:     iType,
		Value:    "",
		Multi:    m,
		Options:  ops,
		Rows:     "",
		Cols:     "",
		Category: "select",
	}
}

// NewTextInput ...
func NewTextInput(name string, label string, class string, id string, rows string, cols string, val string) *InputField {
	return &InputField{
		Name:     name,
		Label:    label,
		Class:    class,
		Id:       id,
		Type:     "",
		Value:    val,
		Multi:    false,
		Options:  nil,
		Rows:     rows,
		Cols:     cols,
		Category: "text",
	}
}
