package phonebook

import (
	"github.com/gouii/simple-phonebook-go/domain/model"
	"github.com/gouii/simple-phonebook-go/domain/repository"
)

type AddUsecase interface {
	Add(AddRequest) (error, AddResponse)
}

type AddResponse interface{}

type AddRequest interface {
	GetName() string
	GetPhone() string
}

type addUc struct {
	pbRepo repository.PhonebookRepository
}

func NewAddUsecase(pbRepo repository.PhonebookRepository) AddUsecase {
	return &addUc{pbRepo}
}

func (auc *addUc) Add(req AddRequest) (error, AddResponse) {
	o := auc.createObject(req)
	err, id := auc.pbRepo.Save(o)
	o.ID = id
	return err, o
}

func (auc *addUc) createObject(req AddRequest) *model.Phonebook {
	data := new(model.Phonebook)
	data.Name = req.GetName()
	data.Phone = req.GetPhone()
	return data
}
