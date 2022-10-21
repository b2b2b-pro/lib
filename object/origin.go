package object

import (
	"encoding/json"
	"io"

	"github.com/b2b2b-pro/lib/torepo"
	"go.uber.org/zap"
)

// Origin - сущность, описывающая происхождение требования.
type Origin struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// ParseOrigin - разбирает полученный json в Origin
func ParseOrigin(data io.Reader) (*Origin, error) {
	var err error
	orgn := &Origin{}

	err = json.NewDecoder(data).Decode(orgn)
	zap.S().Debug("ParseOrigin decode data:\n", data, "\nResult:\n", orgn, "\nНаличие ошибки: ", err, "\n")

	return orgn, err
}

/*
GOrigin конвертирует объект организация (Origin) В формат для передачи по gRPC
From Model to gRPC
*/
func GOrigin(x *Origin) *torepo.Origin {
	return &torepo.Origin{
		Id:          int32(x.ID),
		Description: x.Description,
	}
}

/*
MOrigin конвертирует объект организация (Origin) ИЗ формата для передачи по gRPC
From gRPC to Model
*/
func MOrigin(x *torepo.Origin) *Origin {
	return &Origin{
		ID:          int(x.Id),
		Description: x.Description,
	}
}
