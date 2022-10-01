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
		panic(err)
	}
	gList := models.NewLinkedGroupsList(loadGroups)
	createForm := models.InitializePopupCreatGroupForm()
	model := models.AdminModel{
		Name:        "admin",
		Title:       "Admin Settings",
		Route:       "admin",
		SubRoute:    subRoute,
		Auth:        auth,
		Id:          updateId,
		Method:      "GET",
		CreateGroup: createForm,
		Groups:      loadGroups,
		ListGroups:  gList,
	}
	model.Initialize()
	// 3: TODO Render EDIT FORM based on subRoute (either groups or users in this scenario)
	if !auth.Authenticated {
		lModel := models.LoginModel{Title: "Login", Name: "login", Auth: auth}
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &lModel)
		return
	}
	p.manager.Viewer.RenderTemplate(w, "templates/admin.html", &model)
}

// GroupAdminPage renders the Admin Page
func (p *AdminController) GroupAdminPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	//fmt.Println("VarsTEST:", params.ByName("id"))
	subRoute := params.ByName("child")
	updateId := params.ByName("id")
	fmt.Println("ID Admin:", updateId)
	fmt.Println("subRoute1:", subRoute)
	loadGroups, err := p.groupService.GetMany(auth)
	if err != nil {
		panic(err)
	}
	gList := models.NewLinkedGroupsList(loadGroups)
	createForm := models.InitializePopupCreatGroupForm()
	model := models.AdminModel{
		Name:        "admin",
		Title:       "Group Settings",
		Route:       "admin",
		SubRoute:    subRoute,
		Auth:        auth,
		Id:          updateId,
		Method:      "GET",
		CreateGroup: createForm,
		Groups:      loadGroups,
		ListGroups:  gList,
	}
	model.Initialize()
	// 3: TODO Render EDIT FORM based on subRoute (either groups or users in this scenario)
	if !auth.Authenticated {
		lModel := models.LoginModel{Title: "Login", Name: "login", Auth: auth}
		p.manager.Viewer.RenderTemplate(w, "templates/login.html", &lModel)
		return
	}
	p.manager.Viewer.RenderTemplate(w, "templates/admin.html", &model)
}

// CreateGroupHandler creates a new user group
func (p *AdminController) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
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
	gList := models.NewLinkedGroupsList(loadGroups)
	createForm := models.InitializePopupCreatGroupForm()
	model := models.AdminModel{
		Name:        "admin",
		Title:       "Group Settings",
		Route:       "admin",
		SubRoute:    "",
		Auth:        auth,
		Id:          "",
		Method:      "GET",
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
	//http.Redirect(w, r, "/admin", 201)
	p.manager.Viewer.RenderTemplate(w, "templates/admin.html", &model)
	return
}
