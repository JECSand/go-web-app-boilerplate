package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/JECSand/fetch"
	"github.com/JECSand/go-web-app-boilerplate/models"
	"io"
	"net/http"
	"os"
)

// dataModel is an abstraction of the db model types
type dataModel interface {
	GetJSON() []byte
	GetID() string
}

// APIRequest managers an async request to the API to return data model(s)
type APIRequest struct {
	AuthToken string
}

// getAuthHeaders builds a fetch usable slice of request headers
func (dr *APIRequest) getAuthHeaders() [][]string {
	reqHeaders := fetch.JSONDefaultHeaders()
	if dr.AuthToken != "" {
		return fetch.AppendHeaders(reqHeaders, []string{"Auth-Token", dr.AuthToken})
	}
	return reqHeaders
}

// Get returns one or more dataModel
func (dr *APIRequest) Get(endpoint string) (*fetch.Fetch, error) {
	f, err := fetch.NewFetch(os.Getenv("API_HOST")+endpoint, "GET", dr.getAuthHeaders(), nil)
	if err != nil {
		return f, err
	}
	err = f.Execute("")
	return f, err
}

// Post a new dataModel
func (dr *APIRequest) Post(endpoint string, bodyContents []byte) (*fetch.Fetch, error) {
	body := bytes.NewBuffer(bodyContents)
	f, err := fetch.NewFetch(endpoint, "POST", dr.getAuthHeaders(), body)
	if err != nil {
		return f, err
	}
	err = f.Execute("")
	return f, err
}

// Patch an existing data model
func (dr *APIRequest) Patch(endpoint string, bodyContents []byte) (*fetch.Fetch, error) {
	body := bytes.NewBuffer(bodyContents)
	f, err := fetch.NewFetch(os.Getenv("API_HOST")+endpoint, "PATCH", dr.getAuthHeaders(), body)
	if err != nil {
		return f, err
	}
	err = f.Execute("")
	return f, err
}

// Delete a dataModel
func (dr *APIRequest) Delete(endpoint string) (*fetch.Fetch, error) {
	f, err := fetch.NewFetch(os.Getenv("API_HOST")+endpoint, "DELETE", dr.getAuthHeaders(), nil)
	if err != nil {
		return f, err
	}
	err = f.Execute("")
	return f, err
}

// NewAPIRequest initializes and returns a new APIRequest struct
func NewAPIRequest(authToken string) *APIRequest {
	return &APIRequest{AuthToken: authToken}
}

// APIService is a Generic type struct for organizing dataModel methods
type APIService[T dataModel] struct {
	endpoint string
}

// loadModels loads returned json data into a dataModel
func (api *APIService[T]) loadModel(resp *http.Response) (T, error) {
	var m T
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		return m, err
	}
	if err = resp.Body.Close(); err != nil {
		return m, err
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

// loadModels loads returned json data into a slice of dataModel
func (api *APIService[T]) loadModels(resp *http.Response) ([]T, error) {
	var m []T
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1048576))
	fmt.Println("body:", body)
	if err != nil {
		return m, err
	}
	if err = resp.Body.Close(); err != nil {
		return m, err
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

// GetMany returns a slice of dataModels
func (api *APIService[T]) GetMany(auth *models.Auth) ([]T, error) {
	var m []T
	newReq := NewAPIRequest(auth.AuthToken)
	f, err := newReq.Get(api.endpoint)
	if err != nil {
		return m, err
	}
	f.Resolve()
	return api.loadModels(f.Res)
}

// Get a dataModel
func (api *APIService[T]) Get(auth *models.Auth, filter T) (T, error) {
	var m T
	newReq := NewAPIRequest(auth.AuthToken)
	f, err := newReq.Get(api.endpoint + "/" + filter.GetID())
	if err != nil {
		return m, err
	}
	f.Resolve()
	return api.loadModel(f.Res)
}

// Create a dataModel
func (api *APIService[T]) Create(auth *models.Auth, data T) (T, error) {
	var m T
	newReq := NewAPIRequest(auth.AuthToken)
	f, err := newReq.Post(api.endpoint, data.GetJSON())
	if err != nil {
		return m, err
	}
	f.Resolve()
	return api.loadModel(f.Res)
}

// Update a dataModel
func (api *APIService[T]) Update(auth *models.Auth, data T) (T, error) {
	var m T
	newReq := NewAPIRequest(auth.AuthToken)
	f, err := newReq.Patch(api.endpoint+"/"+data.GetID(), data.GetJSON())
	if err != nil {
		return m, err
	}
	f.Resolve()
	return api.loadModel(f.Res)
}

// Delete a dataModel
func (api *APIService[T]) Delete(auth *models.Auth, filter T) (T, error) {
	var m T
	newReq := NewAPIRequest(auth.AuthToken)
	f, err := newReq.Delete(api.endpoint + "/" + filter.GetID())
	if err != nil {
		return m, err
	}
	f.Resolve()
	return api.loadModel(f.Res)
}

// NewUserService initializes and returns a new APIHandler
func NewUserService() *APIService[*models.User] {
	return &APIService[*models.User]{
		endpoint: os.Getenv("API_HOST") + "/users",
	}
}

// NewGroupService initializes and returns a new APIHandler
func NewGroupService() *APIService[*models.Group] {
	return &APIService[*models.Group]{
		endpoint: os.Getenv("API_HOST") + "/groups",
	}
}

// NewTaskService initializes and returns a new APIHandler
func NewTaskService() *APIService[*models.Task] {
	return &APIService[*models.Task]{
		endpoint: os.Getenv("API_HOST") + "/tasks",
	}
}
