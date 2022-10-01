package models

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
