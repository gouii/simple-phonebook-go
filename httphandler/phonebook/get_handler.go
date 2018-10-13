package phonebook

import (
	"net/http"
	"strconv"

	"github.com/gouii/simple-phonebook-go/context/phonebook"
	"github.com/mmuflih/go-httplib/httplib"
)

type GetHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type getHandler struct {
	uc phonebook.GetUsecase
	rr httplib.RequestReader
}

func NewGetHandler(uc phonebook.GetUsecase, rr httplib.RequestReader) GetHandler {
	return &getHandler{uc, rr}
}

func (ah *getHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(ah.rr.GetRouteParam(r, "id"))
	err, res := ah.uc.Get(int64(id))
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, res)
}
