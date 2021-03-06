package prettyprint

import (
	"fmt"
	"testing"
	"time"
)

type Model struct {
	Firstname  string
	Lastname   string
	MiddleName string
	Age        *int
	Receipts   []struct {
		ItemID      int
		Date        time.Time
		Amount      float32
		Quantity    int
		Description string
	}
}

func TestPrettyPrint(t *testing.T) {
	age := 25
	md := &Model{
		Firstname:  "Alonso",
		Lastname:   "Laurent",
		MiddleName: "Edmund",
		Age:        &age,
	}

	str := Struct(md)

	fmt.Println(str)
}

func TestPrettyPrintJSON(t *testing.T) {

}
