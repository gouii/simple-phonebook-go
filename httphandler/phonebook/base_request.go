package phonebook

type baseRequest struct {
	ID    int64  `json:"-"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (br baseRequest) GetName() string {
	return br.Name
}

func (br baseRequest) GetPhone() string {
	return br.Phone
}

func (br baseRequest) GetID() int64 {
	return br.ID
}
