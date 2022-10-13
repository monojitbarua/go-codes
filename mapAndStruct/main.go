package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var m4 map[string]int //another way to declare map

type person struct {
	FirstName  string           `json:"firstName"`
	SecondName string           `json:"secondName"`
	Age        int              `json:"age"`
	contact    `json:"contact"` // embedded  struct
}
type contact struct {
	MobileNo string `json:"mobNo"`
}

// receiver type func
func (p person) fullName() string {
	return p.FirstName + " " + p.SecondName
}

func main() {

	m := make(map[string]int) //one way to create map
	m2 := map[string]int{}    //anonther way to create map
	m3 := map[string]int{     //anonther way to create map
		"X": 12,
		"Y": 14,
	}

	m["A"] = 1
	m["B"] = 2
	fmt.Println(m)
	fmt.Println(m["B"])
	fmt.Println("m2: ", m2)
	fmt.Println("m3: ", m3)
	fmt.Println("m4: ", m4)

	//
	//delete from map
	delete(m, "A")
	fmt.Println("after delete: ", m)

	//
	//check if value exist in map
	if v, e := m["B"]; e {
		fmt.Println(v, e)
	}

	//
	//iteration map
	for k, v := range m {
		fmt.Println(k, "-", v)
	}

	//
	//
	//struct
	p1 := person{FirstName: "monojit", SecondName: "barua", Age: 33, contact: contact{MobileNo: "87456987"}}
	fmt.Println("full struct print: ", p1)
	fmt.Println("specific field age print: ", p1.contact.MobileNo)
	fmt.Println("receiver func call with full name: ", p1.fullName())

	//
	//user defined type
	type customType int
	var f1 customType
	fmt.Printf("%T\n", f1)
	f1 = 95
	fmt.Println("custom field: ", f1)

	//
	//json marshal
	jsonByte, err := json.Marshal(p1)
	if err != nil {
		panic(1)
	}
	jsonStr := string(jsonByte)
	fmt.Println("json marshal: ", jsonStr)

	//
	//json un-marshal
	var p2 person
	unmarshalStr := `{"firstName":"John","secondName":"Doe","age":44,"contact":{"mobNo":"12345678"}}`
	jsonByteSliceForUnmarshal := []byte(unmarshalStr)
	anotherErr := json.Unmarshal(jsonByteSliceForUnmarshal, &p2)
	if anotherErr != nil {
		panic(1)
	}
	fmt.Println("json unmarshal: ", p2)

	//
	//json Encode
	fmt.Print("Json Encode: ")
	json.NewEncoder(os.Stdout).Encode(p2)

	//
	//json Decoder
	var p3 person
	decodeStr := `{"firstName":"Jerry","secondName":"Hartsfield","age":52,"contact":{"mobNo":"987451"}}`
	decoder := strings.NewReader(decodeStr)
	json.NewDecoder(decoder).Decode(&p3)
	fmt.Println("Json Encode: ", p3)
}
