package phonebook

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/gouii/simple-phonebook-go/context/phonebook"
	"github.com/gouii/simple-phonebook-go/domain/repository"
	"github.com/mmuflih/go-httplib/httplib"
)

func NewRoute(api *mux.Router, muxrr httplib.RequestReader, db *sql.DB) {
	pbRoute := api.PathPrefix("/phonebook").Subrouter()

	repo := repository.NewPhonebookRepo(db)

	addUc := phonebook.NewAddUsecase(repo)
	getUc := phonebook.NewGetUsecase(repo)
	getListUc := phonebook.NewGetListUsecase(repo)
	editUc := phonebook.NewEditUsecase(repo)
	deleteUc := phonebook.NewDeleteUsecase(repo)

	addH := NewAddHandler(addUc, muxrr)
	getH := NewGetHandler(getUc, muxrr)
	getListH := NewGetListHandler(getListUc)
	editH := NewEditHandler(editUc, muxrr)
	deleteH := NewDeleteHandler(deleteUc, muxrr)

	pbRoute.HandleFunc("", addH.Handle).Methods("POST")
	pbRoute.HandleFunc("", getListH.Handle).Methods("GET")
	pbRoute.HandleFunc("/{id}", getH.Handle).Methods("GET")
	pbRoute.HandleFunc("/{id}", editH.Handle).Methods("PUT")
	pbRoute.HandleFunc("/{id}", deleteH.Handle).Methods("DELETE")
}
