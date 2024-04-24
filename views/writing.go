package views

import (
	"errors"
	"goblog/common"
	"goblog/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing

	wr, err := service.Writing()
	if err != nil {
		writing.WriteError(w, errors.New("写作页面请求失败!"))
	}

	writing.WriteData(w, wr)
}
