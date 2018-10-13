package phonebook

import (
	"net/http"
	"strconv"

	"github.com/gouii/simple-phonebook-go/context/phonebook"
	"github.com/mmuflih/go-httplib/httplib"
)

type DeleteHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type deleteHandler struct {
	uc phonebook.DeleteUsecase
	rr httplib.RequestReader
}

func NewDeleteHandler(uc phonebook.DeleteUsecase, rr httplib.RequestReader) DeleteHandler {
	return &deleteHandler{uc, rr}
}

func (ah *deleteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(ah.rr.GetRouteParam(r, "id"))
	err, res := ah.uc.Delete(int64(id))
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, res)
}
