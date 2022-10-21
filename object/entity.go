package object

import (
	"encoding/json"
	"io"

	"github.com/b2b2b-pro/lib/torepo"
)

// Entity - организация, участвующая в системе долговых обязательств
type Entity struct {
	ID        int    `json:"id"`
	INN       string `json:"entity_inn"`
	KPP       string `json:"entity_kpp"`
	ShortName string `json:"short_name"`
	FullName  string `json:"full_name,omitempty"`
}

// ParseEntity - распарсить организацию из json
func ParseEntity(data io.Reader) (*Entity, error) {
	var err error

	frm := &Entity{}
	err = json.NewDecoder(data).Decode(frm)

	return frm, err
}

/*
GEntity конвертирует объект организация (Entity) В формат для передачи по gRPC
From Model to gRPC
*/
func GEntity(x *Entity) *torepo.Entity {
	return &torepo.Entity{
		Id:        int32(x.ID),
		Inn:       x.INN,
		Kpp:       x.KPP,
		ShortName: x.ShortName,
		FullName:  x.FullName,
	}
}

/*
MEntity конвертирует объект организация (Entity) ИЗ формата для передачи по gRPC
From gRPC to Model
*/
func MEntity(x *torepo.Entity) *Entity {
	return &Entity{
		ID:        int(x.Id),
		INN:       x.Inn,
		KPP:       x.Kpp,
		ShortName: x.ShortName,
		FullName:  x.FullName,
	}
}
