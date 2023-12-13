package methods

import (
	"bytes"
	"encoding/json"
	"github.com/chwjbn/go-softether-api/pkg"
)

type DeleteUserParams struct {
	HubName string `json:"HubName_str"`
	Name string `json:"Name_str"`
}

func (e DeleteUserParams) Tags() []string {
	return []string{
		"HubName_str",
		"Name_str",
	}
}

func newDeleteUserParams(hubName string,name string) *DeleteUserParams {
	return &DeleteUserParams{HubName: hubName,Name: name}
}


type DeleteUser struct {
	pkg.Base
	Params *DeleteUserParams `json:"params"`
}

func (e *DeleteUser) Name() string {
	return  e.Base.Name
}

func (e *DeleteUser) GetId() int {
	return e.Id
}

func (e *DeleteUser) SetId(id int) {
	e.Base.Id=id
}

func (e *DeleteUser) Parameter() pkg.Params {
	return e.Params
}

func (e *DeleteUser) Marshall() ([]byte, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

func NewDeleteUser(hubName string,name string) *DeleteUser {
	return &DeleteUser{
		Base:   pkg.NewBase("DeleteUser"),
		Params: newDeleteUserParams(hubName,name),
	}
}
