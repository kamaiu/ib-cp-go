package gen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestOpenAPI(t *testing.T) {
	f, err := ioutil.ReadFile("doc.json")
	if err != nil {
		t.Fatal(err)
	}
	spec := &Spec{}
	err = json.Unmarshal(f, spec)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(spec)
}
