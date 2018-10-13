package phonebook

import (
	"github.com/gouii/simple-phonebook-go/domain/repository"
)

type EditUsecase interface {
	Edit(EditRequest) (error, EditResponse)
}

type EditResponse interface{}

type EditRequest interface {
	GetID() int64
	GetName() string
	GetPhone() string
}

type editUc struct {
	pbRepo repository.PhonebookRepository
}

func NewEditUsecase(pbRepo repository.PhonebookRepository) EditUsecase {
	return &editUc{pbRepo}
}

func (auc *editUc) Edit(req EditRequest) (error, EditResponse) {
	err, o := auc.pbRepo.Find(req.GetID())
	if err != nil {
		return err, nil
	}

	o.Name = req.GetName()
	o.Phone = req.GetPhone()

	err = auc.pbRepo.Update(o)
	return err, o
}
