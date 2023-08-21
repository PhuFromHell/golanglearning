package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	// value := []float64{98, 93, 77, 82, 83}
	// fmt.Println(rangeForLoop(value))
	// time.Sleep(4 * time.Second)

	// if else condition
	// drivingAge(12)

	// witch case
	// fruitCheck("banana")

	// range for loop
	// fmt.Println(rangeForLoop(value))

	// vòng lặp vô hạn
	// infiniteLoop(0)

	// goto key word
	// gotoKeyWord()

	// đọc ghi files
	// fileInOut()

	// recover hoạt động như thế nào
	// fmt.Println(saveDevide(10, 0))
	// fmt.Println(saveDevide(10, 10))

	// http.HandleFunc("/", myHandler1)
	// http.HandleFunc("/detail", myHandler2)
	// http.ListenAndServe("/:8080", nil)

	// timeInGo()
	// timeTickers()

	router := mux.NewRouter()
	emp = append(emp, Employee{ID: "1", FistName: "hoang", LastName: "dinh phu", Address: &Address{City: "Dubai", State: "saudi alba"}})
	emp = append(emp, Employee{ID: "2", FistName: "huynh", LastName: " phu", Address: &Address{City: "Dubai", State: "saudi alba"}})

	router.HandleFunc("/employee", getEmployeeEndPoint).Methods("GET")
	router.HandleFunc("/employee/{id}", getEmpIdEndPoint).Methods("GET")
	router.HandleFunc("/employee/{id}", createEmployeeEndPoint).Methods("POST")
	router.HandleFunc("/employee/{id}", deleteEmployeeEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}

type Employee struct {
	ID       string   `json:"id,omitempty"`
	FistName string   `json:"fistname,omitempty"`
	LastName string   `json:"lastname,omitempty"`
	Address  *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var emp []Employee

func getEmpIdEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range emp {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}

func getEmployeeEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(emp)
}

func createEmployeeEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Employee
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	emp = append(emp, person)
	json.NewEncoder(w).Encode(emp)
}

func deleteEmployeeEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range emp {
		if item.ID == params["id"] {
			emp = append(emp[:index], emp[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(emp)
}

// time Tickers
func timeTickers() {
	tickerValue := time.NewTicker(time.Millisecond * 100)
	go func() {
		for t := range tickerValue.C {
			fmt.Println("ticker at", t)
		}

	}()

	time.Sleep(time.Millisecond * 2000)
	tickerValue.Stop()
	fmt.Println("sticker stopped")

}

// time
func timeInGo() {

	p := fmt.Println
	current_time := time.Now()
	secs := current_time.Unix()
	nanos := current_time.UnixNano()

	fmt.Println("current_time", current_time)

	millis := nanos / 1000000
	p("secs", secs)
	p("millis", millis)
	p("nanos", nanos)
	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))
}

// khái niệm recover hay try/catch
func saveDevide(val1, val2 int) int {
	// defer func() {
	// 	fmt.Println("recover", recover())
	// }()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			// val2 = val1 + val2
			// val1 = val2 - val1
			// val2 = val2 - val1
		}
	}()
	// panic("vppro")
	quotient := val1 / val2

	return quotient
}

// if else condition
func drivingAge(age int) {
	if age >= 18 {
		fmt.Println("you are old enough to drive!")
	} else {
		fmt.Println("you are not old enough to drive!")
	}
}

func fruitCheck(fruit string) {
	switch fruit {
	case "avocado":
		fmt.Println("đây là quả bơ!")
	case "banana":
		fmt.Println("đây là quả chuối!")
	case "apple":
		fmt.Println("đây là quả táo!")
	default:
		fmt.Println("đây không phải là trái cây hoặc quả gì mà tôi không biết!")

	}
}

// nomal for each
func nomalForLoop(value int) {
	for i := 0; i < value; i++ {
		fmt.Println(i)
	}
}

// range for each
func rangeForLoop(input []float64) float64 {
	total := 0.0
	for _, v := range input {
		total += v
	}
	return total / float64(len(input))
}

// vòng lặp vô hạn
func infiniteLoop(count int) {
	for {
		fmt.Println("Số lần lặp:", count)
		count++

		if count == 5 {
			break
		}
	}

}

// key word: goto
func gotoKeyWord() {
	i := 0

start:
	fmt.Println(i)
	i++

	if i < 10000 {
		goto start
	}
}

// file input output
func fileInOut() bool {
	file, err := os.Create("fileInOut.txt")
	if err != nil {
		log.Fatal(err)
		return false
	}

	file.WriteString("this is Hoàng Đình Phú")
	file.Close()

	stream, err := ioutil.ReadFile("fileInOut.txt")
	if err != nil {
		log.Fatal(err)
		return false
	}
	readString := string(stream)
	fmt.Println(readString + " ok")
	file.Close()
	return true

}

func myHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world\n")
}

func myHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello Hoàng A Phú\n")
}
