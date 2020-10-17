package main

import (
	"net/http"

	"github.com/shuwenhe/shuwen-go-web/controller"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/index", controller.Index)
	http.HandleFunc("/list", controller.List)
	http.HandleFunc("/api/getUsers", controller.GetUsers)
	http.HandleFunc("/viewGetUserByID", controller.ViewGetUserByID)
	http.HandleFunc("/api/getUserByID", controller.GetUserByID)
	http.HandleFunc("/add", controller.AddView)
	http.HandleFunc("/edit", controller.ViewEdit)
	http.HandleFunc("/jTest", controller.JTest)

	http.HandleFunc("/api/user/add", controller.AddUser)
	http.HandleFunc("/api/user/del", controller.DeleteUserByUserID)
	http.HandleFunc("/api/user/edit", controller.UpdateUserByID)

	http.HandleFunc("/api/class/get", controller.GetClassByID)
	http.HandleFunc("/api/class/add", controller.AddClass)
	http.HandleFunc("/api/class/upd", controller.UpdateClass)
	http.HandleFunc("/api/class/del", controller.DeleteClassByID)

	http.HandleFunc("/login", controller.Login)
	http.Handle("/api/class/all", controller.Mid(http.HandlerFunc(controller.GetClasses)))

	http.ListenAndServe(":8080", nil)
}
