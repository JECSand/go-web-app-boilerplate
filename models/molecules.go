package models

/*
ListItem Types
*/

// ItemOption ...
type ItemOption struct {
	Form     *Form
	Button   *Button
	Label    string
	Category string
}

// NewDeleteOption ...
func NewDeleteOption(delForm *Form, btn *Button) *ItemOption {
	return &ItemOption{
		Form:     delForm,
		Button:   btn,
		Label:    "",
		Category: "delete",
	}
}

// ListItem ...
type ListItem struct {
	Class    string
	Id       string
	Label    string
	Link     *Link
	Div      *Div
	Button   *Button
	Options  []*ItemOption
	Category string
}

// NewListItem ...
func NewListItem(class string, id string, label string) *ListItem {
	return &ListItem{
		Class:    class,
		Id:       id,
		Label:    label,
		Category: "default",
	}
}

// NewLinkListItem ...
func NewLinkListItem(class string, id string, link *Link, ops []*ItemOption) *ListItem {
	return &ListItem{
		Class:    class,
		Id:       id,
		Link:     link,
		Options:  ops,
		Category: "link",
	}
}

// NewDivListItem ...
func NewDivListItem(class string, id string, div *Div) *ListItem {
	return &ListItem{
		Class:    class,
		Id:       id,
		Div:      div,
		Category: "div",
	}
}

// NewButtonListItem ...
func NewButtonListItem(class string, id string, btn *Button) *ListItem {
	return &ListItem{
		Class:    class,
		Id:       id,
		Button:   btn,
		Category: "button",
	}
}

/*
Div Types
*/

// Div ...
type Div struct {
	Class    string
	Id       string
	Label    string
	Heading  *Heading
	Links    []*Link
	List     *List
	Category string
}

// NewLinkDiv ...
func NewLinkDiv(class string, id string, label string, head *Heading, links []*Link) *Div {
	return &Div{
		Class:    class,
		Id:       id,
		Label:    label,
		Heading:  head,
		Links:    links,
		Category: "links",
	}
}

// NewListDiv ...
func NewListDiv(class string, id string, label string, head *Heading, list *List) *Div {
	return &Div{
		Class:    class,
		Id:       id,
		Label:    label,
		Heading:  head,
		List:     list,
		Category: "list",
	}
}

/*
Table Types
*/

// TableRow ...
type TableRow struct {
	Class        string
	Id           string
	TableData    []*TableData
	TableHeaders []*TableHeader
	Category     string
}

// NewTableRow ...
func NewTableRow(class string, id string, td []*TableData) *TableRow {
	return &TableRow{
		Class:     class,
		Id:        id,
		TableData: td,
		Category:  "default",
	}
}

// NewTableHeaderRow ...
func NewTableHeaderRow(class string, id string, th []*TableHeader) *TableRow {
	return &TableRow{
		Class:        class,
		Id:           id,
		TableHeaders: th,
		Category:     "headers",
	}
}

// TableBody ...
type TableBody struct {
	Class     string
	Id        string
	TableRows []*TableRow
	Category  string
}

// NewTableBody ...
func NewTableBody(class string, id string, tr []*TableRow) *TableBody {
	return &TableBody{
		Class:     class,
		Id:        id,
		TableRows: tr,
		Category:  "default",
	}
}
