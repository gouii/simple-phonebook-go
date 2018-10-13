package phonebook

import (
	"net/http"

	"github.com/gouii/simple-phonebook-go/context/phonebook"
	"github.com/mmuflih/go-httplib/httplib"
)

type GetListHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type getListHandler struct {
	uc phonebook.GetListUsecase
}

func NewGetListHandler(uc phonebook.GetListUsecase) GetListHandler {
	return &getListHandler{uc}
}

func (ah *getListHandler) Handle(w http.ResponseWriter, r *http.Request) {
	err, res := ah.uc.GetList()
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, res)
}
