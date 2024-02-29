package add

import (
	"fmt"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {

	f, err := os.Create(fmt.Sprintf("handlers/%v%v.go", "user", "login"))
	if err != nil {
		t.Fatal(err)
		return
	}

	err = executeTemplate("handler.tmpl", f, service{
		ProjectName: "projectzero",
		LogicName:   "user",
		ServiceName: "login",
	})
	if err != nil {
		return
	}

}
