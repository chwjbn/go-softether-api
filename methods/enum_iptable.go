package methods

import (
	"bytes"
	"encoding/json"
	"github.com/chwjbn/go-softether-api/pkg"
)

type EnumIpTableParams struct {
	HubName string `json:"HubName_str"`
}

func (e EnumIpTableParams) Tags() []string {
	return []string{
		"HubName_str",
	}
}

func newEnumIpTableParams(name string) *EnumIpTableParams {
	return &EnumIpTableParams{HubName: name}
}


type EnumIpTable struct {
	pkg.Base
	Params *EnumIpTableParams `json:"params"`
}

func (e *EnumIpTable) Name() string {
	return  e.Base.Name
}

func (e *EnumIpTable) GetId() int {
	return e.Id
}

func (e *EnumIpTable) SetId(id int) {
	e.Base.Id=id
}

func (e *EnumIpTable) Parameter() pkg.Params {
	return e.Params
}

func (e *EnumIpTable) Marshall() ([]byte, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	res := bytes.Replace(data, []byte("null"), []byte("{}"), -1)
	return res, nil
}

func NewEnumIpTable(name string) *EnumIpTable {
	return &EnumIpTable{
		Base:   pkg.NewBase("EnumIpTable"),
		Params: newEnumIpTableParams(name),
	}
}
