package views

import (
	"errors"
	"goblog/common"
	"goblog/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole

	pigeonholeRes, err := service.GetPostPigeonhole()
	if err != nil {
		pigeonhole.WriteError(w, errors.New("查询错误"))
		return
	}
	pigeonhole.WriteData(w, pigeonholeRes)
}
