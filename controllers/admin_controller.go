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
	userService  *services.APIService[*models.User]
	groupService *services.APIService[*models.Group]
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
	model := models.AdminModel{
		Name:     "admin",
		Title:    "Admin Settings",
		Route:    "admin",
		SubRoute: subRoute,
		Auth:     auth,
		Id:       updateId,
		Method:   "GET",
	}
	model.Initialize()
	// 3: TODO Render EDIT FORM based on subRoute (either groups or users in this scenario)
	if !auth.Authenticated {
		lModel := models.LoginModel{Title: "Login", Name: "login", Auth: auth}
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &lModel)
		return
	}
	if model.SubRoute == "usermenu" {
		model.Title = "Admin Settings"
		model.Name = "Admin Settings"
	}

	p.manager.Viewer.RenderTemplate(w, "templates/admin.html", &model)
}

// AdminUserMenuHandler ...
/*
func (p *AdminController) AdminUserMenuHandler(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.manager.authCheck(r)
	lModel := models.LoginModel{Name: "Login", Title: "Login", Auth: auth}
	if r.Method != http.MethodPatch {
		p.viewer.RenderTemplate(w, "templates/login.html", &lModel)
		return
	}
	user := models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Username:  r.FormValue("username"),
	}
	settings := models.SettingsForm{}
	settings.Load(&user)
	//form := models.InitializeSettingsForm(user)
	status := settings.UpdateSettings(auth)
	model := models.AdminModel{Name: "admin", Title: "Admin Usermenu", SubRoute: "usermenu", Auth: auth}
	model.BuildRoute()
	if status != http.StatusOK {
		p.viewer.RenderTemplate(w, "templates/index.html", &model)
		return
	}
	//model.Form = form
	p.viewer.RenderTemplate(w, "templates/admin.html", &model)
}

// CreateUserForm managers user data
type CreateUserForm struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	UserName  string `json:"username,omitempty"`
	Role      string `json:"role,omitempty"`
	GroupUuid string `json:"groupuuid,omitempty"`
	Password  string `json:"password,omitempty"`
	CPassword string `json:"cpassword,omitempty"`
}

// CreateHandler creates a new user or group
func (p *AdminController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	subRoute := params.ByName("child")
	id := params.ByName("id")
	model := models.AdminModel{
		Name:     "admin",
		Route:    "admin",
		SubRoute: subRoute,
		Auth:     auth,
		Id:       id,
		Method:   "POST",
	}
	// Create New User
	if subRoute == "users" {
		createUserForm := models.CreateUserForm{}
		createUserForm.LoadRequest(r)
		model.Title = "Create New User"
		model.UserTable.CreateUserForm = createUserForm
		model.UserTable.CreateUserForm.Create(auth)
		fmt.Println("CREATE USER")
		// Create New Group
	} else if subRoute == "groups" {
		createGroupForm := models.CreateGroupForm{}
		createGroupForm.LoadRequest(r)
		model.Title = "Create New Group"
		model.GroupTable.CreateGroupForm = createGroupForm
		model.GroupTable.CreateGroupForm.Create(auth)
		fmt.Println("CREATE GROUP")
	}
	model.Id = ""
	model.Method = "GET"
	model.Initialize()
	p.viewer.RenderTemplate(w, "templates/admin.html", &model)
}

// UpdateHandler updates a user or group
func (p *AdminController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	subRoute := params.ByName("child")
	updateId := params.ByName("id")
	model := models.AdminModel{
		Name:     "admin",
		Route:    "admin",
		SubRoute: subRoute,
		Auth:     auth,
		Id:       updateId,
		Method:   "PATCH",
	}
	// Update User
	if subRoute == "users" && updateId != "" {
		updateUserForm := models.UpdateUserForm{}
		updateUserForm.LoadRequest(r)
		model.Title = "Update User: " + updateUserForm.UserName
		model.UserTable.UpdateUserForm = updateUserForm
		model.UserTable.UpdateUserForm.Update(auth, model.Id)
		fmt.Println("UPDATE USER")
		// Update Group
	} else if subRoute == "groups" && updateId != "" {
		updateGroupForm := models.UpdateGroupForm{}
		updateGroupForm.LoadRequest(r)
		model.Title = "Update Group: " + updateGroupForm.Name
		model.GroupTable.UpdateGroupForm = updateGroupForm
		model.GroupTable.UpdateGroupForm.Update(auth, model.Id)
		fmt.Println("UPDATE GROUP")
	}
	model.Id = ""
	model.Method = "GET"
	model.Initialize()
	p.viewer.RenderTemplate(w, "templates/admin.html", &model)
}

// DeleteHandler deletes a user or group
func (p *AdminController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	subRoute := params.ByName("child")
	delId := params.ByName("id")
	model := models.AdminModel{
		Name:     "admin",
		Route:    "admin",
		SubRoute: subRoute,
		Auth:     auth,
		Id:       delId,
		Method:   "DELETE",
	}
	// Delete User
	if subRoute == "users" && delId != "" {
		fmt.Println("DELETE USER")
		// Delete Group
	} else if subRoute == "groups" && delId != "" {
		fmt.Println("DELETE GROUP")
	}
	p.viewer.RenderTemplate(w, "templates/admin.html", &model)
}

// UserMenuHandler
func (p *AdminController) UsermenuPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	fmt.Println("VarsTEST:", params.ByName("child"))
	model := models.AdminModel{Title: "Admin", SubRoute: params.ByName("child"), Name: "admin", Auth: auth}
	model.Initialize()
	if !auth.Authenticated {
		lModel := models.LoginModel{Title: "Login", Name: "login", Auth: auth}
		p.viewer.RenderTemplate(w, "templates/login.html", &lModel)
		return
	}
	if model.SubRoute == "usermenu" {
		model.Title = "Admin Settings"
		model.Name = "Admin Settings"

	}
	p.viewer.RenderTemplate(w, "templates/account.html", &model)
}

func (p *AdminController) GroupmenuPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	fmt.Println("VarsTEST:", params.ByName("child"))
	model := models.AdminModel{Title: "Admin", SubRoute: params.ByName("child"), Name: "admin", Auth: auth}
	model.Initialize()
	if !auth.Authenticated {
		lModel := models.LoginModel{Title: "Login", Name: "login", Auth: auth}
		p.viewer.RenderTemplate(w, "templates/login.html", &lModel)
		return
	}
	if model.SubRoute == "groupmenu" {
		model.Title = "Admin Settings"
		model.Name = "Admin Settings"

	}
	p.viewer.RenderTemplate(w, "templates/account.html", &model)
}
*/
