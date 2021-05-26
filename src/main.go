package main

import (
	"C"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Customer struct {
	Id        int64   `gorm:"primary key" json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	City      string  `json:"city"`
	Phone     string  `json:"phone"`
	Height    float64 `json:"height"`
	Gender    string  `json:"gender"`
	Password  string  `json:"password"`
	Married   bool    `json:"married"`
	Created   int64
	Updated   int64
}
type Output struct {
	Id        int64   `gorm:"primary key" json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	City      string  `json:"city"`
	Phone     string  `json:"phone"`
	Height    float64 `json:"height"`
	Gender    string  `json:"gender"`
	Password  string  `json:"password"`
	Married   bool    `json:"married"`
}

var db *gorm.DB
//getting records from list of Ids
func getSomeUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	var arr []int64
	err := json.NewDecoder(r.Body).Decode(&arr)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 0; i < len(arr); i++ {
		var res []Output
		db.Table("user").Select("id", "first_name", "last_name", "city", "phone", "height", "gender", "password", "married").Where("id = ?", arr[i]).Scan(&res)
		if res == nil {

			var str string
			str = "No such Id exits"
			json.NewEncoder(w).Encode(str)
		} else {
			json.NewEncoder(w).Encode(res)
		}
	}

}
//get user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	vars := mux.Vars(r)
	var res []Output
	db.Table("user").Select("id", "first_name", "last_name", "city", "phone", "height", "gender", "password", "married").Where("id = ?", vars["id"]).Scan(&res)
	if res == nil {

		var str string
		str = "No such Id exits"
		json.NewEncoder(w).Encode(str)
	} else {
		json.NewEncoder(w).Encode(res)
	}
}
//get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	var res []int64
	db.Table("user").Select("id").Scan(&res)
	if res == nil {
		var str string
		str = "No such Id exits"
		json.NewEncoder(w).Encode(str)
	} else {
		json.NewEncoder(w).Encode(res)
	}
}
//insert a new record into table
func insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	var custom Customer
	err := json.NewDecoder(r.Body).Decode(&custom)
	if err != nil {

		var str string
		str = "No input"
		json.NewEncoder(w).Encode(str)
	} else {
		id := custom.Id
		firstname := custom.FirstName
		lastname := custom.LastName
		city := custom.City
		phone := custom.Phone
		height := custom.Height
		gender := custom.Gender
		password := custom.Password

		married := custom.Married
		h := sha1.New()
		h.Write([]byte(password))
		bs := hex.EncodeToString(h.Sum(nil))

		created := time.Now().Unix()
		updated := time.Now().Unix()
		user := Customer{Id: id, FirstName: firstname, LastName: lastname, City: city, Phone: phone, Height: height, Gender: gender, Password: bs, Married: married, Created: created, Updated: updated}
		db.Create(&Customer{Id: id, FirstName: firstname, LastName: lastname, City: city, Phone: phone, Height: height, Gender: gender, Password: bs, Married: married, Created: created, Updated: updated})
		json.NewEncoder(w).Encode(user)
	}
}
//returns the name of created table
func (Customer) TableName() string {
	return "user"
}

func InitialMigration() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("customer.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect")
	}
	db.AutoMigrate(&Customer{})
	return db, nil
}
func main() {

	database, err := InitialMigration()
	db = database
	if err != nil {
		panic("Not connected to database!")
	}
	handlesRequests()
}
func handlesRequests() {
	myrouter := mux.NewRouter()
	myrouter.HandleFunc("/api/v1/user/gets", getSomeUser).Methods("POST")

	myrouter.HandleFunc("/api/v1/user/create", insert).Methods("POST")
	myrouter.HandleFunc("/api/v1/user/{id}", getUser).Methods("GET")
	myrouter.HandleFunc("/api/v1/user/fetch", getUsers).Methods("POST")

	fmt.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", myrouter))

}
