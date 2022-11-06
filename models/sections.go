package models

// InitializeUserSettings ...
func InitializeUserSettings(user *User, admin bool) *Settings {
	SettingsForm := InitializeUserSettingsForm(user, admin)
	// 1. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	var links []*Link
	infoLink := NewLink("active", "", "/account/settings", "Change User Info", true)
	links = append(links, infoLink)
	if !admin {
		pwLink := NewLink("", "", "/account/password", "Change Password", true)
		links = append(links, pwLink)
	}
	optionsCol := NewLinkDiv("uneven columnOne", "", "", NewColumnHeading("Options", ""), links)
	// 2. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	unLink := NewLink("active", "", "", user.Username, true)
	rLink := NewLink("active", "", "", user.Role, true)
	infoCol := NewLinkDiv("uneven columnTwo", "", "", NewColumnHeading("Current Info", ""), []*Link{unLink, rLink})
	return NewSettings("", "", SettingsForm, optionsCol, infoCol)
}

// InitializeGroupSettings ...
func InitializeGroupSettings(group *Group, users []*User) *Settings {
	SettingsForm := InitializeGroupSettingsForm(group)
	// 1. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	uList := NewLinkedList(users, "/admin/", "", true, true, false, "", "")
	usersCol := NewListDiv("uneven columnOne", "", "", NewColumnHeading("Group Users", ""), uList)
	// 2. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	nameLink := NewLink("active", "", "", group.Name, true)
	infoCol := NewLinkDiv("uneven columnTwo", "", "", NewColumnHeading("Group Info", ""), []*Link{nameLink})
	return NewSettings("", "", SettingsForm, usersCol, infoCol)
}

// InitializeTaskOverview instantiates a task Overview Abstract	// TODO NEXT - START HERE, called in TASK CONTROLLER
func InitializeTaskOverview(inTasks []*Task, groupUsers *GroupUsersDTO, groups []*Group) *Overview {
	// Init group and user select filter drop down
	var filters []*List
	var overviewScripts []*Script
	userSelect := NewList(groupUsers.Users, "dropdown", "/data/", "/tasks", true, "*", "getTasks")
	filters = append(filters, userSelect)
	if len(groups) > 1 {
		groupSelect := NewList(groups, "dropdown", "/data/", "/tasks", true, groupUsers.Group.Id, "getTasks")
		filters = append(filters, groupSelect)
	}
	nsTasks, ipTasks, comTasks := SplitTasksByStatus(inTasks)
	// Init Not Started List DIV
	overviewScripts = append(overviewScripts, &Script{Category: "manageTasks"})
	overviewScripts = append(overviewScripts, &Script{Category: "postCheck", Axios: true})
	nsTasks = append(nsTasks, NewTask("000000000000000000000000", "TEMPLATE", NOTSTARTED)) // Template Task li for axios to use
	notStartedList := NewLinkedList(nsTasks, "/data/", "/check", false, true, true, "Completed", "postCheck")
	notStartedCol := NewListDiv("even columnOne", "", "", NewColumnHeading("Not Started", ""), notStartedList)
	// Init In Progress List DIV
	inProgressList := NewLinkedList(ipTasks, "/data/", "/check", false, true, true, "Completed", "postCheck")
	inProgressCol := NewListDiv("even columnTwo", "", "", NewColumnHeading("In Progress", ""), inProgressList)
	// Init In Completed List DIV
	completedList := NewLinkedList(comTasks, "/data/", "/check", false, true, true, "Completed", "postCheck")
	completedCol := NewListDiv("even columnThree", "", "", NewColumnHeading("Completed", ""), completedList)
	//	6) Init "drag and drop" jquery script
	//	7) Init and return NewTaskOverview
	return NewTasksOverview("tasksOverview", "", filters, notStartedCol, inProgressCol, completedCol, overviewScripts)
}
