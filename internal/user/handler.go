package user

import (
	"net/http"
	"restapi/internal/handlers"
	"restapi/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{} // helper

const (
	usersURL = "/users"
	userURL  = "/user/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	h.logger.Info("registering user handler")
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.PUT(userURL, h.UpdateUser)
	router.DELETE(userURL, h.DeleteUser)

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.WriteHeader(200)
	w.Write([]byte("this is list of users"))
}
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is get user"))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is partially update user"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is update user"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("this is create user"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is delete user"))
}
