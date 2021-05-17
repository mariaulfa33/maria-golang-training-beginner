package lib

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type helloWorld struct {
	foo    string
	bar    int
	xendit Xendit
}

type Xendit struct {
	Location  string
	Valuation float32
}

type HelloWorldJson struct {
	Foo    string `json:"foo"`
	Bar    int    `json:"bar"`
	Xendit Xendit `json:"xendit"`
}
type XenditJson struct {
	Location  string  `json:"location"`
	Valuation float32 `json:"valuation"`
}

func DataStruct1() {
	world := helloWorld{
		foo: "foo",
		bar: 10,
		xendit: Xendit{
			Location:  "Worldwide",
			Valuation: 1000,
		},
	}
	fmt.Println(world)
}

func DataStruct2() {
	world2 := make(map[string]helloWorld)
	world2["Earth"] = helloWorld{
		foo: "foo",
		bar: 10,
		xendit: Xendit{
			Location:  "Worldwide",
			Valuation: 1000,
		},
	}

	world2["Mars"] = helloWorld{
		foo: "foo",
		bar: 10,
		xendit: Xendit{
			Location:  "Worldwide",
			Valuation: 1000,
		},
	}

	fmt.Println(world2)

}

func DataStruct3() {
	data := string(`
    {
        "foo": "foo",
        "bar": "bar",
        "xendit": {
            "location": "Worldwide",
            "valuation": 1000
        }
    }
    `)
	var hello HelloWorldJson
	json.Unmarshal([]byte(data), &hello)

	fmt.Println(hello)
}

/**
===================================
EXCERCISE
===================================
*/

type Employee struct {
	Name           string   `json:"name"`
	Entity         string   `json:"entity"`
	EmployeeNumber int      `json:"employee_number"`
	Salary         float32  `json:"salary"`
	Domicile       Domicile `json:"xendit"`
}

type Domicile struct {
	Country  string `json:"country"`
	IsRemote bool   `json:"is_remote"`
}

func DataStructExcercise() {
	// TODO: task #1 - give me a skeleton!
	data := string(`
{
	"name": "Golang",
	"entity": "Xendit",
	"employee_number": 10,
	"salary": 1.5,
	"domicile": {
		"country": "ID",
		"is_remote": true
	}
}`)
	var employee Employee

	if err := json.Unmarshal([]byte(data), &employee); err != nil {
		fmt.Println(err)

		fmt.Println("Task #1 failed!")
		return
	}
	// TODO: task #2 - I am a legal employee, include me into your database!
	database := make(map[string]Employee)
	database["Golang"] = employee
	if !reflect.DeepEqual(database["Golang"], employee) {
		fmt.Println("Task #2 failed!")
		return
	}
	fmt.Println("All Passed!")

	return
}
