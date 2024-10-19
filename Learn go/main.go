package main

import (
	"fmt"

	"maps"
	"slices"
	"time"
)

func add(a, b int) int {
	return a + b
}

func getLang() (string, string, string) {
	return "go", "programming language", "created by google"
}
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func counter() func() int {

	var count int = 0
	return func() int {
		count += 1
		return count
	}
}
func anyType(data ...interface{}) []interface{} {
	return data
}
func changeNum(number *int) {
	*number = 5
	print("In change num", *number)
}
func changeSlice(number *[]int) {
	*number = append(*number, 50)
	fmt.Println("In change num", number)
}

type customer struct {
	name    string
	age     int
	address string
}
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func (o *order) changeStatus(status string) {
	o.status = status
}
func (o *order) getAmount() float32 {
	return o.amount
}

func newOrder(id string, amount float32, status string, customer customer) *order {
	myorder := order{
		id:       id,
		amount:   amount,
		status:   status,
		customer: customer,
	}
	return &myorder
}
func main() {
	fmt.Println("Hello, world!")
	fmt.Println((1 + 2))
	fmt.Println(true)
	fmt.Println(10.6)
	// variables
	name := "majed"
	fmt.Println(name)
	var age int64 = 50
	fmt.Println(age)
	var isMarried bool = true
	fmt.Println(isMarried)
	// constant
	const id int = 20
	fmt.Println(id)
	const (
		port = 2000
		host = "localhost"
	)
	fmt.Println(port, host)

	// while not exists in go, we can use for loop to iterate over a range of numbers
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}
	//for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	for k := range 6 {
		fmt.Println(k)
	}
	// if else
	if 2 < 5 {
		fmt.Println("2 is less than 5")
	} else if 4 < 5 {
		fmt.Println("4 is less than 5")
	} else {
		fmt.Println("2 is not less than 5")
	}
	// switch
	x := 3
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
	default:
		fmt.Println("x is not 1, 2 or 3")
	}
	//type swich
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("i is an integer")
		case string:
			fmt.Println("i is a string")
		default:
			fmt.Println("i is of unknown type")
		}
	}
	whoAmI(10)

	// array
	var arr [5]int
	var ar = [5]int{1, 2, 3, 4, 5}
	num := [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println("num is", num)
	fmt.Println("ae is", ar)
	arr[0] = 1
	arr[1] = 2
	fmt.Println("arr is", arr)
	fmt.Println("arr is", len(arr))
	fmt.Println("arr is", cap(arr))

	// slice - dynamic array
	//  snum:= []int{}
	var mnum = make([]int, 0, 6)

	mnum = append(mnum, 10, 67)
	mnum = append(mnum, 20)
	mnum = append(mnum, 30)
	mnum = append(mnum, 40)
	mnum = append(mnum, 50)
	mnum = append(mnum, 50)
	mnum = append(mnum, 50)
	mnum2 := make([]int, len(mnum), cap(mnum))
	copy(mnum2, mnum)

	fmt.Println(mnum[2:3], mnum2[:3])
	fmt.Println(slices.Equal(mnum, mnum2))
	fmt.Println(slices.Delete(mnum, 0, 1))
	fmt.Println(mnum)
	// map
	m := make(map[string]string)
	m["name"] = "majed666"
	m["age"] = "30"
	fmt.Println(m["name"])
	delete(m, "name")
	clear(m)
	fmt.Println(m)
	m2 := map[string]string{"nameId5": "55"}
	m3 := map[string]string{"nameId5": "55"}

	m2["age"] = "60"
	fmt.Println()
	v, ok := m2["nameId5"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("not ok")
	}
	delete(m2, "age")
	clear(m2)
	fmt.Println(m2)
	fmt.Println(maps.Equal(m3, m2))

	//range
	// sum := 0
	// nums := []int{1, 2, 3, 4, 5}
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println(sum)
	m3 = map[string]string{"nameId5": "55"}
	for k, v := range m3 {
		fmt.Println(k, v)

	}
	for i, e := range "golang" {
		fmt.Println(i, string(e))
	}

	//function
	result := add(2, 5)
	fmt.Println(result)
	lang1, lang2, lang3 := getLang()
	fmt.Println(lang1, lang2, lang3)
	fmt.Println(getLang())

	//variatic function
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(n...))
	fmt.Println(anyType(1, "hello world", 5.0))
	//closure

	increment := counter()
	fmt.Println(increment())

	// pointer
	// number := 10
	// changeNum(&number)
	// fmt.Println("After changing number in main function", number)
	numberArr := []int{1, 2, 3, 4, 5}
	changeSlice(&numberArr)
	fmt.Println("After changing number in main function", numberArr)
	// struct 
	// myorder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "pending",
	// }
	// myorder2:= order{
	//     id:"1",
	//     amount: 50.00,
	//     status: "pending",
	//     createdAt: time.Now(),

	// }
	// myorder.createdAt = time.Now()
	// myorder.changeStatus("Paid")
	// fmt.Println("order struct", myorder)
	// fmt.Println("order amount", myorder.getAmount())
	newCustomer := customer{
		name:    "majed",
		age:     30,
		address: "USA",
	}
	myorder := newOrder("1", 50.00, "pending", newCustomer)
	fmt.Println("order struct", myorder)

	// another way struct
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language.name, language.isGood)

    // interface struct
    
}
