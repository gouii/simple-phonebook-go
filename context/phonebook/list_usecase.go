package phonebook

import (
	"github.com/gouii/simple-phonebook-go/domain/repository"
)

type GetListUsecase interface {
	GetList() (error, GetListResponse)
}

type GetListResponse interface{}

type getListUc struct {
	pbRepo repository.PhonebookRepository
}

func NewGetListUsecase(pbRepo repository.PhonebookRepository) GetListUsecase {
	return &getListUc{pbRepo}
}

func (auc *getListUc) GetList() (error, GetListResponse) {
	err, o := auc.pbRepo.List()
	if err != nil {
		return err, nil
	}
	return nil, o
}
