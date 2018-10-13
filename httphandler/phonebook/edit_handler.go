package phonebook

import (
	"net/http"
	"strconv"

	"github.com/gouii/simple-phonebook-go/context/phonebook"
	"github.com/mmuflih/go-httplib/httplib"
)

type EditHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type editHandler struct {
	uc phonebook.EditUsecase
	rr httplib.RequestReader
}

func NewEditHandler(uc phonebook.EditUsecase, rr httplib.RequestReader) EditHandler {
	return &editHandler{uc, rr}
}

func (ah *editHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := baseRequest{}
	ah.rr.GetJsonData(r, &req)
	id, _ := strconv.Atoi(ah.rr.GetRouteParam(r, "id"))
	req.ID = int64(id)
	err, res := ah.uc.Edit(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, res)
}
