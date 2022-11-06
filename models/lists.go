package models

// NewList initializes a new list of groups for rendering
func NewList[T DataModel](m []T, class string, baseURL string, endpoint string, chkBox bool, chkRule string, scriptType string) *List {
	var listItems []*ListItem
	for _, gr := range m {
		var ops []*ItemOption
		if chkBox {
			reqURL := baseURL + gr.GetClass(true) + "/" + gr.GetID() + endpoint
			chkInput := NewCheckboxInput("", "pill checkbox", "check"+gr.GetID(), gr.GetBoolField(chkRule), reqURL, scriptType)
			chkOp := NewCheckOption("", chkInput)
			ops = append(ops, chkOp)
		}
		listItems = append(listItems, NewListItem("pill", gr.GetID(), gr.GetLabel(), ops))
	}
	return NewUnorderedList(class, "", listItems, nil)
}

// NewLinkedList initializes a new linked list of DataModel for rendering
func NewLinkedList[T DataModel](m []T, baseURL string, endpoint string, delete bool, search bool, chkBox bool, chkRule string, scriptType string) *List {
	var listItems []*ListItem
	listId, _ := GenerateUuid()
	for _, gr := range m {
		gLink := NewLink("pill", "", baseURL+gr.GetClass(true)+"/"+gr.GetID()+endpoint, gr.GetLabel(), false)
		var ops []*ItemOption
		if delete {
			defForm := InitializePopupDeleteForm(gr, baseURL)
			btn := defForm.Popup
			defForm.Popup = nil
			delOp := NewDeleteOption(defForm, btn)
			ops = append(ops, delOp)
		}
		if chkBox {
			reqURL := baseURL + gr.GetClass(true) + "/" + gr.GetID() + endpoint
			chkInput := NewCheckboxInput("", "pill checkbox", "check"+gr.GetID(), gr.GetBoolField(chkRule), reqURL, scriptType)
			chkOp := NewCheckOption("", chkInput)
			ops = append(ops, chkOp)
		}
		liClass := "pill"
		if gr.GetID() == "000000000000000000000000" {
			liClass += " template"
		}
		listItems = append(listItems, NewLinkListItem(liClass, gr.GetID(), gLink, ops))
	}
	if search {
		inId, _ := GenerateUuid()
		searchInput := NewSearchInput("linked filter", inId, listId, "Enter a name: ")
		return NewUnorderedList("linked", listId, listItems, searchInput)
	}
	return NewUnorderedList("linked", listId, listItems, nil)
}
