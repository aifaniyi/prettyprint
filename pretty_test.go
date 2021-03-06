package prettyprint

import (
	"strings"
	"testing"
)

type Model struct {
	Firstname  string
	Lastname   string
	MiddleName string
	Age        *int
	Receipt    *Receipt
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
	}

	str := Struct(md)

	if !strings.Contains(str, "Alonso") {
		t.Fatalf("output must contain %s", md.Firstname)
	}

	if !strings.Contains(str, "5000001") {
		t.Fatalf("output must contain %d", md.Receipt.ItemID)
	}
}

func TestPrettyPrintJSON(t *testing.T) {

}
