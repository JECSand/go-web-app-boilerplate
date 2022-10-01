package models

/*
ListItem Types
*/

// ListItem ...
type ListItem struct {
	Class    string
	Id       string
	Label    string
	Link     *Link
	Div      *Div
	Button   *Button
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
func NewLinkListItem(class string, id string, link *Link) *ListItem {
	return &ListItem{
		Class:    class,
		Id:       id,
		Link:     link,
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
	Links    []*Link
	Category string
}

// NewLinkDiv ...
func NewLinkDiv(class string, id string, label string, links []*Link) *Div {
	return &Div{
		Class:    class,
		Id:       id,
		Label:    label,
		Links:    links,
		Category: "links",
	}
}
