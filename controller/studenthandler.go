package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/shuwenhe/shuwen-go-web/dao"
	"github.com/shuwenhe/shuwen-go-web/model"
	"github.com/shuwenhe/shuwen-go-web/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	num := r.FormValue("num")
	s, err := dao.Login(num)
	if err != nil {
		w.Write(utils.NewResult(utils.Fail, "student is not exist", nil))
		return
	}
	pass := r.FormValue("pass")
	if pass != s.Pass {
		w.Write(utils.NewResult(utils.Fail, "password is error,please try it again!", nil))
		return
	}
	data := model.Jwt{
		ID:   s.ID,
		Name: s.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}
	jwts := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, err := jwts.SignedString([]byte(`k`))
	if err != nil {
		w.Write(utils.NewResult(utils.Fail, "token fail,try it again!"))
		return
	}
	w.Write(utils.NewResult(utils.Success, "student is login success", token))
}

func Mid(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")
		if token == "" {
			w.Write(utils.NewResult(utils.Fail, "token is null", nil))
			return
		}
		d1 := &model.Jwt{}
		j, err := jwt.ParseWithClaims(token, d1, func(token *jwt.Token) (interface{}, error) {
			return []byte("k"), nil
		})
		if err != nil {
			w.Write(utils.NewResult(utils.Fail, "token is error", err.Error()))
		}
		if j.Valid { //
			next.ServeHTTP(w, r)
		} else {

		}
	})
}
