# prettyprint

A golang library to pretty print golang structs using reflections.

### Installation
```bash
go get github.com/aifaniyi/prettyprint
```

### Example

```go
package main

import (
	"fmt"

	"github.com/aifaniyi/prettyprint"
)

type model struct {
    Name string
    email *string    
}

func main(){
    age := 25
	md := &Model{
		Firstname:  "Alonso",
		Lastname:   "Laurent",
		MiddleName: "Edmund",
		Age:        &age,
		Receipts: &Receipt{
			ItemID:   5000001,
			Amount:   50.2,
			Quantity: 3,
            Description: "sample receipt data",
        },
	}

	str := prettyprint.Struct(md)

	fmt.Println(str)
    /* would print
        {
        "Firstname": "Alonso"
        "Lastname": "Laurent"
        "MiddleName": "Edmund"
        "Age": 25
            "Receipts":    {
                "ItemID": 5000001
                "Amount": 50.2
                "Quantity": 3
                "Description": ""
            }
        }
    */
}
```
