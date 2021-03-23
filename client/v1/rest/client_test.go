package rest

import (
	"encoding/json"
	"fmt"
	"testing"
)

func toString(v json.Marshaler) string {
	b, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestClient_Sso_Validate_GET(t *testing.T) {
	c := NewClient(DefaultUrl, true)
	r, err := c.Sso_Validate_GET()
	if err != nil {
		t.Fatal(err)
	}

	tickle, err := c.Tickle_POST()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(toString(r))
	fmt.Println(toString(tickle))
}
