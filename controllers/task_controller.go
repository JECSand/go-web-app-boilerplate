package controllers

import (
	"encoding/json"
	"github.com/JECSand/go-web-app-boilerplate/models"
	"github.com/JECSand/go-web-app-boilerplate/services"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

// TaskController structures the set of app page views
type TaskController struct {
	manager      *ControllerManager
	taskService  *services.TaskService
	userService  *services.UserService
	groupService *services.GroupService
}

// TasksPage renders the Variable Page
func (p *TaskController) TasksPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	created := r.URL.Query().Get("created")
	groupUsers, err := p.groupService.GetGroupUsers(auth, &models.Group{Id: auth.GroupId})
	if err != nil {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	userTasks, err := p.taskService.GetMany(auth)
	userTasksList := models.NewLinkedList(userTasks, "/", true, true, true)
	userTasksList.Script = &models.Script{Category: "postCheck"}
	createForm := models.InitializePopupCreateTaskForm(groupUsers.Users)
	model := models.TaskModel{
		Title:         "Tasks",
		SubRoute:      params.ByName("child"),
		Name:          "tasks",
		Auth:          auth,
		CreateTask:    createForm,
		OverviewTasks: userTasksList,
	}
	model.BuildRoute()
	model.Heading = models.NewHeading("Tasks Overview", "w3-wide text")
	if !auth.Authenticated {
		http.Redirect(w, r, "/logout", 303)
		return
	}
	if created != "" {
		var alert *models.Alert
		if created == "yes" {
			alert = models.NewSuccessAlert("Task Created", true)
		} else if created == "no" {
			alert = models.NewErrorAlert("Error Creating Task", true)
		}
		model.Alert = alert
		model.Status = true
	}
	model.Auth = auth
	p.manager.Viewer.RenderTemplate(w, "templates/tasks.html", &model)
}

// CreateTaskHandler creates a new user task
func (p *TaskController) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	returnMsg := "yes"
	auth, _ := p.manager.authCheck(r)
	dt, err := models.ConvertToDateTime(r.FormValue("due"))
	if err != nil {
		returnMsg = "no"
		http.Redirect(w, r, "/tasks?created="+returnMsg, 303)
		return
	}
	task := &models.Task{
		Name:        r.FormValue("name"),
		Due:         dt, //	i.e. 2022-10-10T11:11
		Description: r.FormValue("description"),
		GroupId:     r.FormValue("group_id"),
		UserId:      r.FormValue("user_id"),
	}
	if !auth.RootAdmin && task.GroupId != auth.GroupId {
		task.GroupId = auth.GroupId
	}
	task, err = p.taskService.Create(auth, task)
	if err != nil {
		returnMsg = "no"
	}
	http.Redirect(w, r, "/tasks?created="+returnMsg, 303)
	return
}

// CompleteTaskHandler updates whether a task is completed or not
func (p *TaskController) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	auth, _ := p.manager.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			return
		}
		return
	}
	if err = r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			return
		}
		return
	}
	if err = json.Unmarshal(body, &t); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			return
		}
		return
	}
	updateId := params.ByName("id")
	t.Id = updateId
	task, err := p.taskService.Update(auth, &t)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(task); err != nil {
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(task); err != nil {
		return
	}
	return
}
