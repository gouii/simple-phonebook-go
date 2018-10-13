package phonebook

import (
	"github.com/gouii/simple-phonebook-go/domain/repository"
)

type DeleteUsecase interface {
	Delete(int64) (error, DeleteResponse)
}

type DeleteResponse interface{}

type deleteUc struct {
	pbRepo repository.PhonebookRepository
}

func NewDeleteUsecase(pbRepo repository.PhonebookRepository) DeleteUsecase {
	return &deleteUc{pbRepo}
}

func (auc *deleteUc) Delete(id int64) (error, DeleteResponse) {
	err, o := auc.pbRepo.Find(id)
	if err != nil {
		return err, nil
	}
	err = auc.pbRepo.Delete(id)
	return err, o
}
