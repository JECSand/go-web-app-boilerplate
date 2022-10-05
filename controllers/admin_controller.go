package controllers

import (
	"fmt"
	"github.com/JECSand/go-web-app-boilerplate/models"
	"github.com/JECSand/go-web-app-boilerplate/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// AdminController structures the set of app page views
type AdminController struct {
	manager      *ControllerManager
	userService  *services.UserService
	groupService *services.GroupService
}

// AdminPage renders the Admin Page
func (p *AdminController) AdminPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	//fmt.Println("VarsTEST:", params.ByName("id"))
	//vars returns and empty map array
	subRoute := params.ByName("child")
	updateId := params.ByName("id")
	fmt.Println("ID Admin:", updateId)
	//the subroute gives an index error and crashes the app if sURL[1]. If it is zero however it will work but it will never
	//work for create since that needs to be vars 1 for the if to catch it.
	fmt.Println("subRoute1:", subRoute)
	loadGroups, err := p.groupService.GetMany(auth)
	if err != nil {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	gList := models.NewLinkedList(loadGroups, true, true)
	createForm := models.InitializePopupCreateGroupForm()
	model := models.AdminModel{
		Name:        "admin",
		Title:       "Admin Settings",
		Route:       "admin",
		SubRoute:    subRoute,
		Auth:        auth,
		Id:          updateId,
		Method:      "GET",
		Heading:     models.NewHeading("Admin", "w3-wide text"),
		CreateGroup: createForm,
		Groups:      loadGroups,
		ListGroups:  gList,
	}
	model.Initialize()
	if !auth.Authenticated {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	p.manager.Viewer.RenderTemplate(w, "templates/admin.html", &model)
}

// AdminGroupsPage renders the Admin Page
func (p *AdminController) AdminGroupsPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	//fmt.Println("VarsTEST:", params.ByName("id"))
	//vars returns and empty map array
	subRoute := params.ByName("child")
	updateId := params.ByName("id")
	// Loop over header names
	del := r.URL.Query().Get("deleted")
	fmt.Println("\n\nID Admin Group: ", updateId)
	fmt.Println("subRoute1:", subRoute)
	loadGroups, err := p.groupService.GetMany(auth)
	if err != nil {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	gList := models.NewLinkedList(loadGroups, true, true)
	createForm := models.InitializePopupCreateGroupForm()
	model := models.AdminModel{
		Name:        "admin",
		Title:       "Admin Settings",
		Route:       "admin",
		SubRoute:    "groups",
		Auth:        auth,
		Id:          updateId,
		Method:      "GET",
		Heading:     models.NewHeading("Manage Groups", "w3-wide text"),
		CreateGroup: createForm,
		Groups:      loadGroups,
		ListGroups:  gList,
	}
	model.Initialize()
	if !auth.Authenticated {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	if del != "" {
		var alert *models.Alert
		if del == "yes" {
			alert = models.NewSuccessAlert("Group Deleted", true)
		} else {
			alert = models.NewErrorAlert("Error Deleting Group", true)
		}
		model.Alert = alert
		model.Status = true
	}
	p.manager.Viewer.RenderTemplate(w, "templates/admin.html", &model)
}

// AdminUserPage renders the Admin Page for a specific Group
func (p *AdminController) AdminUserPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	up := r.URL.Query().Get("updated")

	fmt.Println("VarsTEST:", params.ByName("id"))
	paramId := params.ByName("id")
	fmt.Println("ID Admin:", paramId)
	if !auth.RootAdmin && paramId != auth.GroupId { // if user not Root Admin scope to auth GroupID of user
		paramId = auth.GroupId
	}
	user, err := p.userService.Get(auth, &models.User{Id: paramId})
	if err != nil {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	userSettings := models.InitializeUserSettings(user, true)
	model := models.AdminModel{
		Name:         "admin",
		Title:        "User Settings",
		Route:        "admin",
		SubRoute:     "users",
		Auth:         auth,
		Id:           paramId,
		Method:       "GET",
		Heading:      models.NewHeading("User Settings", "w3-wide text"),
		UserSettings: userSettings,
	}
	model.Initialize()
	if !auth.Authenticated {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	if up != "" {
		var alert *models.Alert
		if up == "yes" {
			alert = models.NewSuccessAlert("User Updated", true)
		} else if up == "no" {
			alert = models.NewErrorAlert("Error Updating User", true)
		}
		model.Alert = alert
		model.Status = true
	}
	p.manager.Viewer.RenderTemplate(w, "templates/admin.html", &model)
}

// AdminGroupPage renders the Admin Page for a specific Group
func (p *AdminController) AdminGroupPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	up := r.URL.Query().Get("updated")
	del := r.URL.Query().Get("deleted")
	fmt.Println("VarsTEST:", params.ByName("id"))
	paramId := params.ByName("id")
	fmt.Println("ID Admin:", paramId)
	if !auth.RootAdmin && paramId != auth.GroupId { // if user not Root Admin scope to auth GroupID of user
		paramId = auth.GroupId
	}
	groupUsers, err := p.groupService.GetGroupUsers(auth, &models.Group{Id: paramId})
	if err != nil {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	groupSettings := models.InitializeGroupSettings(groupUsers.Group, groupUsers.Users)
	createForm := models.InitializePopupCreateUserForm([]*models.Group{groupUsers.Group}, true)
	model := models.AdminModel{
		Name:          "admin",
		Title:         "Group Settings",
		Route:         "admin",
		SubRoute:      "groups",
		Auth:          auth,
		Id:            paramId,
		Method:        "GET",
		Heading:       models.NewHeading("Group Admin", "w3-wide text"),
		GroupSettings: groupSettings,
		CreateUser:    createForm,
		Users:         groupUsers.Users,
	}
	model.Initialize()
	if !auth.Authenticated {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	if up != "" || del != "" {
		var alert *models.Alert
		if up == "yes" {
			alert = models.NewSuccessAlert("Group Updated", true)
		} else if up == "no" {
			alert = models.NewErrorAlert("Error Updating Group", true)
		}
		if del == "yes" {
			alert = models.NewSuccessAlert("User Deleted", true)
		} else if del == "no" {
			alert = models.NewErrorAlert("Error Deleting User", true)
		}
		model.Alert = alert
		model.Status = true
	}
	p.manager.Viewer.RenderTemplate(w, "templates/admin.html", &model)
}

// AdminCreateGroupHandler creates a new user group
func (p *AdminController) AdminCreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	var errMsg string
	auth, _ := p.manager.authCheck(r)
	group := &models.Group{
		Name: r.FormValue("name"),
	}
	var loadGroups []*models.Group
	group, err := p.groupService.Create(auth, group)
	if err != nil {
		errMsg = err.Error()
	}
	loadGroups, err = p.groupService.GetMany(auth)
	if err != nil && errMsg == "" {
		errMsg = err.Error()
	}
	gList := models.NewLinkedList(loadGroups, true, true)
	createForm := models.InitializePopupCreateGroupForm()
	model := models.AdminModel{
		Name:        "admin",
		Title:       "Group Settings",
		Route:       "admin",
		SubRoute:    "groups",
		Auth:        auth,
		Id:          "",
		Method:      "GET",
		Heading:     models.NewHeading("Manage Groups", "w3-wide text"),
		CreateGroup: createForm,
		Groups:      loadGroups,
		ListGroups:  gList,
		Status:      true,
	}
	var alert *models.Alert
	if errMsg != "" {
		alert = models.NewErrorAlert(errMsg, true)
	} else {
		alert = models.NewSuccessAlert(group.Name+" Created", true)
	}
	model.Alert = alert
	p.manager.Viewer.RenderTemplate(w, "templates/admin.html", &model)
	return
}

// AdminCreateUserHandler creates a new user group
func (p *AdminController) AdminCreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var errMsg string
	auth, _ := p.manager.authCheck(r)
	user := &models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Username:  r.FormValue("username"),
		Password:  r.FormValue("password"),
		GroupId:   r.FormValue("group_id"),
		Role:      r.FormValue("role"),
	}
	fmt.Println("\n\nCHECK NEW USR ROLE HERE: ", user.Role)
	if !auth.RootAdmin && user.GroupId != auth.GroupId { // if user not Root Admin scope to auth GroupID of user
		user.GroupId = auth.GroupId
	}
	paramId := user.GroupId
	user, err := p.userService.Create(auth, user)
	if err != nil {
		errMsg = err.Error()
	}
	groupUsers, err := p.groupService.GetGroupUsers(auth, &models.Group{Id: paramId})
	if err != nil {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	groupSettings := models.InitializeGroupSettings(groupUsers.Group, groupUsers.Users)
	createForm := models.InitializePopupCreateUserForm([]*models.Group{groupUsers.Group}, true)
	model := models.AdminModel{
		Name:          "admin",
		Title:         "Group Settings",
		Route:         "admin",
		SubRoute:      "groups",
		Auth:          auth,
		Id:            paramId,
		Method:        "GET",
		Heading:       models.NewHeading("Group Admin", "w3-wide text"),
		GroupSettings: groupSettings,
		CreateUser:    createForm,
		Users:         groupUsers.Users,
		Status:        true,
	}
	var alert *models.Alert
	if errMsg != "" {
		alert = models.NewErrorAlert(errMsg, true)
	} else {
		alert = models.NewSuccessAlert(user.Username+" Created", true)
	}
	model.Alert = alert
	//http.Redirect(w, r, "/admin", 201)
	p.manager.Viewer.RenderTemplate(w, "templates/admin.html", &model)
	return
}

// AdminDeleteUserHandler creates a new user group
func (p *AdminController) AdminDeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	delMsg := "yes"
	auth, _ := p.manager.authCheck(r)
	user := &models.User{
		Id: r.FormValue("id"),
	}
	user, err := p.userService.Get(auth, user)
	if err != nil {
		delMsg = "no"
	}
	if !auth.RootAdmin && user.GroupId != auth.GroupId { // if user not Root Admin end session
		http.Redirect(w, r, "/logout", 303)
		return
	}
	_, err = p.userService.Delete(auth, user)
	if err != nil {
		delMsg = "no"
	}
	http.Redirect(w, r, "/admin/groups/"+user.GroupId+"?deleted="+delMsg, 303)
	return
}

// AdminDeleteGroupHandler creates a new user group
func (p *AdminController) AdminDeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	delMsg := "yes"
	auth, _ := p.manager.authCheck(r)
	group := &models.Group{
		Id: r.FormValue("id"),
	}
	if !auth.RootAdmin { // if user not Root Admin end session
		http.Redirect(w, r, "/logout", 303)
		return
	}
	_, err := p.groupService.Delete(auth, group)
	if err != nil {
		delMsg = "no"
	}
	http.Redirect(w, r, "/admin/groups?deleted="+delMsg, 303)
	return
}

// AdminUpdateGroupHandler creates a new user group
func (p *AdminController) AdminUpdateGroupHandler(w http.ResponseWriter, r *http.Request) {
	upMsg := "yes"
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	paramId := params.ByName("id")
	group := &models.Group{
		Id:   paramId,
		Name: r.FormValue("name"),
	}
	if !auth.RootAdmin { // if user not Root Admin end session
		http.Redirect(w, r, "/logout", 303)
		return
	}
	_, err := p.groupService.Update(auth, group)
	if err != nil {
		upMsg = "no"
	}
	http.Redirect(w, r, "/admin/groups/"+paramId+"?updated="+upMsg, 303)
	return
}
