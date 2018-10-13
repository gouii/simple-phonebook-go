package phonebook

import (
	"net/http"

	"github.com/gouii/simple-phonebook-go/context/phonebook"
	"github.com/mmuflih/go-httplib/httplib"
)

type AddHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type addHandler struct {
	uc phonebook.AddUsecase
	rr httplib.RequestReader
}

func NewAddHandler(uc phonebook.AddUsecase, rr httplib.RequestReader) AddHandler {
	return &addHandler{uc, rr}
}

func (ah *addHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := baseRequest{}
	ah.rr.GetJsonData(r, &req)
	err, res := ah.uc.Add(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, res)
}
