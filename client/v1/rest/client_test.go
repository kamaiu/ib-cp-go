package rest

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
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

	contract, err := c.Iserver_Contract_Conid_Info_GET("12087792")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(toString(contract))

	contract, err = c.Iserver_Contract_Conid_Info_GET("470728476")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(toString(contract))

	snapshot, snapshot400, err := c.Iserver_Marketdata_Snapshot_GET("470728476", time.Now().Add(-(time.Hour*0)).Unix()*1000, "31,55,70,71,84,85,86,87,88,6457,7285,7308,7309,7310,7311,7633")
	if err != nil {
		t.Fatal(err)
	}
	if snapshot400 != nil {
		t.Fatal(snapshot400)
	}
	fmt.Println(toString(snapshot))

	scannerParams, err := c.Iserver_Scanner_Params_GET()
	if err != nil {
		t.Fatal(err)
	}
	_ = scannerParams
	//fmt.Println(toString(scannerParams))

	strikes, strikes500, err := c.Iserver_Secdef_Strikes_GET("107113386", "OPT", "APR21", "SMART")
	if err != nil {
		t.Fatal(err)
	}
	if strikes500 != nil {
		t.Fatal(strikes500)
	}
	fmt.Println(toString(strikes))

	strikeInfo, strikeInfo500, err := c.Iserver_Secdef_Info_GET("107113386", "OPT", "APR21", "SMART", "300", "C")
	if err != nil {
		t.Fatal(err)
	}
	if strikeInfo500 != nil {
		t.Fatal(strikes500)
	}
	fmt.Println(toString(strikeInfo))
}

func TestSnapshotJSON(t *testing.T) {
	data := `
[{"conid":470728476,"server_id":"q0","_updated":1616548969282,"6119":"q0","70":"5.69","87":"10.2K","87_raw":10200.0,"31":"2.61","71":"2.45","6509":"ZBd","7308":"0.271","85":"1","7309":"0.022","7311":"0.143","84":"2.60","86":"2.67","7310":"-0.282","88":"1","6508":"&serviceID1=1042&serviceID2=661&serviceID3=122&serviceID4=123&serviceID5=1049&serviceID6=1048&serviceID7=1047&serviceID8=203&serviceID9=775&serviceID10=215&serviceID11=1046&serviceID12=204&serviceID13=205&serviceID14=216&serviceID15=1078&serviceID16=206&serviceID17=108&serviceID18=6666&serviceID19=109","55":"FB"}]
`
	_ = data
}
