package object

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestParseObligation(t *testing.T) {
	str := "{\"id\":1,\"debtor_id\":11,\"creditor_id\":21,\"cost\":2,\"origin_id\":1,\"payment_date\":\"2021-10-30\"}"
	x1 := strings.NewReader(str)
	r1, err := ParseObligation(x1)
	if err != nil {
		t.Errorf("ParseObligation() return error: %v", err)
	}
	if r1.ID != 1 || r1.DebtorID != 11 {
		t.Errorf("Неправильно распарсил\n")
	}
	t.Logf("Result: %v\n", r1)

	b, err := json.Marshal(r1)
	if err != nil {
		t.Errorf("Marshal error: %b\n", err)
	}
	t.Logf("Marshal res: %v\n", b)

	x2 := strings.NewReader(string(b))
	r2, err := ParseObligation(x2)
	if err != nil {
		t.Errorf("ParseObligation() return error: %v", err)
	}
	if !reflect.DeepEqual(r1, r2) {
		t.Errorf("полученные структуры отличаются")
	}
}

/*
type Obligation struct {
	ID         int         `json:"id"`
	DebtorID   int         `json:"debtor_id"`   // должник
	CreditorID int         `json:"creditor_id"` // кому должны
	Cost       float32     `json:"cost"`        // стоимость
	OriginID   int         `json:"origin_id"`   // ,omitempty
	Date       PaymentDate `json:"payment_date"`
}
*/

/*
ParseObligation - разбирает полученный json в Obligation
{id:1,debtor_tin:"11",debtor_name:"рога",creditor_tin:"21", creditor_name: "копыта", cost: 20, origin_id:0}
*/
// func ParseObligation(data io.Reader) (*Obligation, error) {
