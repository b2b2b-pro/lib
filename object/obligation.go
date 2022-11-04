package object

import (
	"encoding/json"
	"io"

	"github.com/b2b2b-pro/lib/torepo"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.uber.org/zap"
)

// Obligation - денежное обязательство
type Obligation struct {
	ID         int         `json:"id"`
	DebtorID   int         `json:"debtor_id"`    // должник
	CreditorID int         `json:"creditor_id"`  // кому должны
	Cost       int64       `json:"cost"`         // сумма
	OriginID   int         `json:"origin_id"`    // происхождение долгового обязательства
	Date       PaymentDate `json:"payment_date"` // дата наступления денежного обязательства
}

// ParseObligation - разбирает полученный json в Obligation
func ParseObligation(data io.Reader) (*Obligation, error) {
	var err error

	o := &Obligation{}

	err = json.NewDecoder(data).Decode(o)
	if o.OriginID == 0 {
		o.OriginID = 1
	}
	zap.S().Debug("ParseObligation Decode data:\n", data, "\nResult:\n", o, "\nНаличие ошибки: ", err, "\n")
	return o, err
}

/*
GObligation конвертирует объект обязательство (Obligation) В формат для передачи по gRPC
From Model to gRPC
*/
func GObligation(tkn string, x *Obligation) *torepo.Obligation {
	return &torepo.Obligation{
		Tkn:        tkn,
		Id:         int32(x.ID),
		Dbtorid:    int32(x.DebtorID),
		Creditorid: int32(x.CreditorID),
		Cost:       x.Cost,
		Originid:   int32(x.OriginID),
		Date:       timestamppb.New(x.Date.Get()),
	}
}

/*
MObligation конвертирует объект организация (Obligation) ИЗ формата для передачи по gRPC
From gRPC to Model
*/
func MObligation(x *torepo.Obligation) *Obligation {
	if x.Date == nil {
		x.Date = timestamppb.Now()
	}
	return &Obligation{
		ID:         int(x.Id),
		DebtorID:   int(x.Dbtorid),
		CreditorID: int(x.Creditorid),
		Cost:       x.Cost,
		OriginID:   int(x.Originid),
		Date:       NewPaymentDate(x.Date.AsTime()),
	}
}
