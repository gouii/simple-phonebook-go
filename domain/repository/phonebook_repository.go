package repository

import (
	"database/sql"
	"fmt"

	"github.com/gouii/simple-phonebook-go/domain/model"
)

type PhonebookRepository interface {
	Save(*model.Phonebook) (error, int64)
	Update(*model.Phonebook) error
	Find(int64) (error, *model.Phonebook)
	List() (error, []*model.Phonebook)
	Delete(int64) error
}

type phonebookRepo struct {
	db *sql.DB
}

func NewPhonebookRepo(db *sql.DB) PhonebookRepository {
	return &phonebookRepo{db}
}

func (pbr *phonebookRepo) Save(data *model.Phonebook) (error, int64) {
	query := "INSERT INTO phonebooks (name, phone) VALUES (?, ?)"
	ps, err := pbr.db.Prepare(query)
	if err != nil {
		return err, 0
	}
	res, err := ps.Exec(data.Name, data.Phone)
	if err != nil {
		return err, 0
	}
	id, err := res.LastInsertId()
	ps.Close()
	return err, id
}

func (pbr *phonebookRepo) Update(data *model.Phonebook) error {
	query := "UPDATE phonebooks SET name = ?, phone = ? WHERE id = ? "
	ps, err := pbr.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = ps.Exec(data.Name, data.Phone, data.ID)
	if err != nil {
		return err
	}
	ps.Close()
	return err
}

func (pbr *phonebookRepo) Find(id int64) (error, *model.Phonebook) {
	query := "SELECT id, name, phone FROM phonebooks WHERE id = ?"
	row := pbr.db.QueryRow(query, id)
	o := new(model.Phonebook)
	err := row.Scan(&o.ID, &o.Name, &o.Phone)
	return err, o
}

func (pbr *phonebookRepo) List() (error, []*model.Phonebook) {
	var lst []*model.Phonebook
	query := "SELECT id, name, phone FROM phonebooks"
	rows, err := pbr.db.Query(query)
	if err != nil {
		return err, lst
	}
	for rows.Next() {
		o := new(model.Phonebook)
		err := rows.Scan(&o.ID, &o.Name, &o.Phone)
		if err != nil {
			fmt.Println("Error while populate data %v", err)
		}
		lst = append(lst, o)
	}
	rows.Close()
	return err, lst
}

func (pbr *phonebookRepo) Delete(id int64) error {
	query := "DELETE FROM phonebooks WHERE id = ? "
	ps, err := pbr.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = ps.Exec(id)
	if err != nil {
		return err
	}
	ps.Close()
	return err

}
