package tasks

import (
	"encoding/json"
	"fmt"
)

type info struct {
	FetchTime string //`json:"FetchTime"`
	Symbol    string //`json:"Symbol"`
}

type fund struct {
	Name       string   //`json:"Name"`
	DataPoints []string //`json:"DataPoints"`
}

type combined struct {
	MyInfo *info `json:"myInfo"`
	MyFund *fund `json:"myFund"`
}

// func (c *Combined) MarshalJSON() ([]byte, error) {
// 	fmt.Printf("MarshalJSON() | c:\n%#v\n", c)

// 	jsonInfo, err := json.Marshal(c.myInfo)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf("MarshalJSON() | jsonInfo:\n%#v\n", string(jsonInfo))

// 	jsonFund, err := json.Marshal(c.myFund)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf("MarshalJSON() | jsonFund:\n%#v\n", string(jsonFund))

// 	bCombined := struct {
// 		bInfo []byte
// 		bFund []byte
// 	}{
// 		bInfo: jsonInfo,
// 		bFund: jsonFund,
// 	}
// 	fmt.Printf("\nMarshalJSON() | bCombined:\n%#v\n\n", bCombined)

// 	return json.Marshal(bCombined)
// }

func DoSomething() error {
	fmt.Println("DoSomething()")

	theInfo := new(info)
	theInfo.FetchTime = "2021-10-06T19:10:02"
	theInfo.Symbol = "1234567"

	fmt.Printf("theInfo initialized:\n%#v\n", theInfo)

	theFund := new(fund)
	theFund.Name = "Fund One"
	theFund.DataPoints = []string{"data1", "data2", "data3"}

	fmt.Printf("theFund initialized:\n%#v\n", theFund)

	theCombined := combined{theInfo, theFund}

	fmt.Printf("theCombined initialized:\n%#v\n", theCombined)

	jsonInfo, err := json.Marshal(&theInfo)
	if err != nil {
		return err
	}

	fmt.Printf("theInfo marshalled:\n%#v\n", string(jsonInfo))

	jsonFund, err := json.Marshal(theFund)
	if err != nil {
		return err
	}

	fmt.Printf("theFund marshalled:\n%#v\n", string(jsonFund))

	jsonCombined, err := json.Marshal(theCombined)
	if err != nil {
		return err
	}

	fmt.Printf("\ntheCombined marshalled:\n%#v\n", string(jsonCombined))

	return nil
}

var s string

func Modify(s *string) {
	*s = "modified"
}

func CauseModify() error {
	fmt.Printf("s: %s\n", s)
	Modify(&s)
	fmt.Printf("s: %s\n", s)

	return nil
}
