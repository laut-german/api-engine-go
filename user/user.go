package userroute

import (
	"net/http"
	"strconv"
	serverengine "web-engine-go/server-engine"
	"web-engine-go/utils"
)

type User struct {}

func New() *serverengine.Route {
	user := &User{}
	return &serverengine.Route{
		WithLogger: true,
		Handler:  user,
	}
}


func (u *User)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	switch head {
	case "":
		u.list(w, r)
	case "detail":
		head, r.URL.Path = utils.ShiftPath(r.URL.Path)
		id, err := strconv.Atoi(head)
		if err != nil {
			utils.Respond(w, r, http.StatusBadRequest, err)
			return
		}

		u.detail(w, r, id)
	default:
		utils.Respond(w,r, http.StatusNotFound, "user path not found")
	}
}

func (u *User) list(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, r, http.StatusOK, []int{1,2,3,4,5,6})
}

func (u *User) detail(w http.ResponseWriter, r *http.Request, id int) {
	utils.Respond(w, r, http.StatusOK, id)
}