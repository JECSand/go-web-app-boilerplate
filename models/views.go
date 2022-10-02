package models

import "fmt"

// IndexModel of dynamic data used for index view
type IndexModel struct {
	Title    string
	Name     string
	SubRoute string
	Route    string
	Heading  *Heading
	Auth     *Auth
	Form     *Form
}

// RegisterModel of dynamic data used for register view
type RegisterModel struct {
	Title    string
	Name     string
	SubRoute string
	Route    string
	Heading  *Heading
	Auth     *Auth
	Form     *Form
}

// LoginModel of dynamic data used for login view
type LoginModel struct {
	Title    string
	Variable string
	Name     string
	SubRoute string
	Route    string
	Heading  *Heading
	Auth     *Auth
	Form     *Form
}

// BuildRoute ...
func (lm *LoginModel) BuildRoute() {
	route := lm.Name + "/" + lm.SubRoute
	lm.Route = route
}

// AboutModel of dynamic data used for about view
type AboutModel struct {
	Title    string
	Variable string
	Name     string
	SubRoute string
	Route    string
	Heading  *Heading
	Auth     *Auth
}

// BuildRoute ...
func (am *AboutModel) BuildRoute() {
	route := am.Name + "/" + am.SubRoute
	am.Route = route
}

// VariableModel of dynamic data used for variable view
type VariableModel struct {
	Title    string
	Variable string
	Name     string
	SubRoute string
	Route    string
	Heading  *Heading
	Auth     *Auth
}

// BuildRoute ...
func (vm *VariableModel) BuildRoute() {
	route := vm.Name + "/" + vm.SubRoute
	vm.Route = route
}

// AdminModel ...
type AdminModel struct {
	Title         string
	Variable      string
	Name          string
	SubRoute      string
	Route         string
	Id            string
	Method        string
	Heading       *Heading
	CreateGroup   *Form
	CreateUser    *Form
	GroupSettings *Settings
	UserSettings  *Settings
	ListGroups    *List
	ListUsers     *List
	Users         []*User
	Groups        []*Group
	Auth          *Auth
	Alert         *Alert
	Status        bool
}

// BuildRoute ...
func (adm *AdminModel) BuildRoute() {
	route := adm.Name
	if adm.SubRoute != "" {
		route = route + "/" + adm.SubRoute
		if adm.Id != "" {
			route = "/" + adm.Id
		}
	}
	adm.Route = route
}

// Initialize a new Admin Page Data Model
func (adm *AdminModel) Initialize() {
	adm.BuildRoute()
}

// AccountModel ...
type AccountModel struct {
	Title    string
	Variable string
	Name     string
	SubRoute string
	Route    string
	Heading  *Heading
	Auth     *Auth
	User     *User
	Alert    *Alert
	Settings *Settings
	//SettingsForm *Form
}

// BuildRoute ...
func (acm *AccountModel) BuildRoute() {
	route := acm.Name + "/" + acm.SubRoute
	fmt.Println("subroutetest:", acm.SubRoute)
	acm.Route = route
	fmt.Println("routetest:", acm.Route)
}

// Initialize a new Account Page Data Model
func (acm *AccountModel) Initialize() {
	acm.BuildRoute()
}
