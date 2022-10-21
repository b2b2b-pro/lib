package object

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// PaymentDate вынесен в отдельный тип для библиотеки, разбирающей и собирающей Json
type PaymentDate struct {
	d time.Time
}

func (t *PaymentDate) Get() time.Time {
	return t.d
}

func NewPaymentDate(d time.Time) PaymentDate {
	return PaymentDate{d: d}
}

func (t *PaymentDate) String() string {
	return fmt.Sprintf("%d-%d-%d", t.d.Year(), t.d.Month(), t.d.Day())
}

func (t *PaymentDate) UnmarshalJSON(b []byte) error {
	const shortForm = "2006-01-02"
	v := string(b[1 : len(b)-1])

	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}

	t.d = parse

	return nil
}

func (t *PaymentDate) Short() time.Time {
	return t.d
}

func (t *PaymentDate) MarshalJSON() ([]byte, error) {
	v := fmt.Sprintf("%d-%d-%d", t.d.Year(), t.d.Month(), t.d.Day())

	return json.Marshal(v)
}

func (t *PaymentDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "2006-01-02"
	var v string
	var err error

	err = d.DecodeElement(&v, &start)
	if err != nil {
		zap.S().Debugf("DecodeElement error: %v\n", err)
	}

	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}

	*t = PaymentDate{parse}

	return nil
}
