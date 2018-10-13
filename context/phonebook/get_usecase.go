package phonebook

import (
	"github.com/gouii/simple-phonebook-go/domain/repository"
)

type GetUsecase interface {
	Get(int64) (error, GetResponse)
}

type GetResponse interface{}

type getUc struct {
	pbRepo repository.PhonebookRepository
}

func NewGetUsecase(pbRepo repository.PhonebookRepository) GetUsecase {
	return &getUc{pbRepo}
}

func (auc *getUc) Get(id int64) (error, GetResponse) {
	err, o := auc.pbRepo.Find(id)
	if err != nil {
		return err, nil
	}
	return err, o
}
