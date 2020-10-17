package model

// User struct
type User struct {
	ID       int64  `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Num      string `json:"num" form:"num"`
	Password string `json:"password" form:"password"`
	Logo     string `json:"img_path" form:"img_path"`
	Age      int64  `json:"age" form:"age"`
}
