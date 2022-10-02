package models

// NewList initializes a new list of groups for rendering
func NewList[T DataModel](m []T) *List {
	var listItems []*ListItem
	for _, gr := range m {
		listItems = append(listItems, NewListItem("group", gr.GetID(), gr.GetLabel()))
	}
	return NewUnorderedList("groups", "", listItems)
}

// NewLinkedList initializes a new linked list of DataModel for rendering
func NewLinkedList[T DataModel](m []T, delete bool) *List {
	var listItems []*ListItem
	for _, gr := range m {
		gLink := NewLink("pill", "", "/admin/"+gr.GetClass(true)+"/"+gr.GetID(), gr.GetLabel(), false)
		var ops []*ItemOption
		if delete {
			// NewDeleteOption(delForm *Form, btn *Button)
			defForm := InitializePopupDeleteForm(gr)
			btn := defForm.Popup
			defForm.Popup = nil
			delOp := NewDeleteOption(defForm, btn)
			ops = append(ops, delOp)
		}
		listItems = append(listItems, NewLinkListItem("pill", gr.GetID(), gLink, ops))
	}
	return NewUnorderedList("linked", "", listItems)
}
