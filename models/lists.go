package models

// NewGroupsList initializes a new list of groups for rendering
func NewGroupsList(groups []*Group) *List {
	var listItems []*ListItem
	for _, gr := range groups {
		listItems = append(listItems, NewListItem("group", gr.Id, gr.Name))
	}
	return NewUnorderedList("groups", "labeled", listItems)
}

// NewLinkedGroupsList initializes a new linked list of groups for rendering
func NewLinkedGroupsList(groups []*Group) *List {
	var listItems []*ListItem
	for _, gr := range groups {
		gLink := NewLink("group", "", "/groups/"+gr.Id, gr.Name)
		listItems = append(listItems, NewLinkListItem("group", gr.Id, gLink))
	}
	return NewUnorderedList("groups", "labeled", listItems)
}
