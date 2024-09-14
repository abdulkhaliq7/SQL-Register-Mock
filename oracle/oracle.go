package oracle

import (
	"github.com/abdulkhaliq7/db/sql"
	"fmt"
)

func init() {

	sql.Register("oracle", &Oracle{})
}

type Oracle struct {
}

func (o Oracle) Open(name string) {

	fmt.Printf("%v Database", name)
}