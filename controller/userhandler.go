package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/shuwenhe/shuwen-go-web/dao"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := dao.GetUsers()
	buf, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

func ViewGetUserByID(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/pages/getUserByID.html")
	w.Write(buf)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	user, _ := dao.GetUserByID(id)
	w.Header().Set("Content-Type", "application/json")
	buf, _ := json.Marshal(user)
	w.Write(buf)
}

func AddView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/pages/add.html")
	w.Write(buf)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var name = r.FormValue("kname")
	if name == "" {
		w.Write([]byte("user name is not null !"))
		return
	}
	num := r.Form.Get("knum")
	password := r.Form.Get("kpassword")
	logo := r.Form.Get("klogo")
	age := r.Form.Get("kage")
	intAge, _ := strconv.Atoi(age)
	err := dao.AddUser(name, num, password, logo, intAge)
	if err != nil {
		w.Write([]byte("Insert error!"))
	}
	w.Write([]byte("Insert success !"))
}

func JTest(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/static/script/jquery-1.7.2.js")
	w.Write(buf)
}

func DeleteUserByUserID(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("kid")
	userID, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte("delete err ,please try it again!"))
		return
	}
	err2 := dao.DeleteUserByUserID(userID)
	if err2 != nil {
		w.Write([]byte("delete err !"))
	}
	w.Write([]byte("delete success !"))
}

func ViewEdit(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/pages/edit.html")
	w.Write(buf)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var name = r.FormValue("kname")
	if name == "" {
		w.Write([]byte("user name is not null !"))
		return
	}
	num := r.Form.Get("knum")
	password := r.Form.Get("kpassword")
	logo := r.Form.Get("klogo")
	age := r.Form.Get("kage")
	intAge, _ := strconv.Atoi(age)
	id := r.Form.Get("kid")
	userID, _ := strconv.Atoi(id)
	err := dao.UpdateUserByID(name, num, password, logo, intAge, userID)
	if err != nil {
		w.Write([]byte("Update error!"))
	}
	w.Write([]byte("update success!"))
}
