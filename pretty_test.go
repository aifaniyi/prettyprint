package prettyprint

import (
	"fmt"
	"strings"
	"testing"
)

type Model struct {
	Firstname    string
	Lastname     string
	MiddleName   string
	Age          *int
	Receipt      *Receipt
	PastReceipts []Receipt
}

type Receipt struct {
	ItemID      int
	Amount      float32
	Quantity    int
	Description string
}

func TestPrettyPrint(t *testing.T) {
	age := 25
	md := &Model{
		Firstname:  "Alonso",
		Lastname:   "Laurent",
		MiddleName: "Edmund",
		Age:        &age,
		Receipt: &Receipt{
			ItemID:   5000001,
			Amount:   50.2,
			Quantity: 3,
		},
		PastReceipts: []Receipt{
			{ItemID: 5000002, Amount: 10.2, Quantity: 1},
			{ItemID: 5000003, Amount: 1.2, Quantity: 1},
			{ItemID: 5000042, Amount: 17.25, Quantity: 3},
		},
	}

	str := Struct(md)

	if !strings.Contains(str, "Alonso") {
		t.Fatalf("output must contain %s", md.Firstname)
	}

	if !strings.Contains(str, "5000001") {
		t.Fatalf("output must contain %d", md.Receipt.ItemID)
	}

	fmt.Printf(str)
}

func TestPrettyPrintJSON(t *testing.T) {

}
