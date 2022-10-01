package models

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
	Category string
}

// NewLink constructs and returns a new Link
func NewLink(class string, id string, ref string, label string) *Link {
	return &Link{
		Class:    class,
		Id:       id,
		Ref:      ref,
		Label:    label,
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
