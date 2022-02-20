package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type FinancialMove struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
	Type string `json:"type (income or expense)"`
	Comments string `json:"comments"`
	Category *Category `json:"category"`
}

type Category struct{
		ReferenceBook1 string `json:"referenceBook1"`
		ReferenceBook2 string `json:"referenceBook2"`
}

var financialMoves []FinancialMove

func getfinancialMoves(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(financialMoves)
}

func deleteFinancialMove(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	for index, item:=range financialMoves{

		if item.ID == params["id"]{
			financialMoves=append(financialMoves[:index],financialMoves[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(financialMoves)
}

func getFinancialMove(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, item := range financialMoves{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createFinancialMove(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var financialMove FinancialMove
	_=json.NewDecoder(r.Body).Decode(&financialMove)
	financialMove.ID = strconv.Itoa(rand.Intn(100000000))
	financialMoves=append(financialMoves,financialMove)
	json.NewEncoder(w).Encode(financialMove)
}

func updateFinancialMove(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index,item := range financialMoves{
		if item.ID == params["id"]{
			financialMoves = append(financialMoves[:index],financialMoves[index+1:]...)
			var financialMove FinancialMove
			_=json.NewDecoder(r.Body).Decode(&financialMove)
			financialMove.ID=params["id"]
			financialMoves=append(financialMoves,financialMove)
			json.NewEncoder(w).Encode(financialMove)
			return
		}
	}
}

func main(){
	r :=mux.NewRouter()

	financialMoves=append(financialMoves,FinancialMove{ID:"1",Name:"Move1",Date:"10.10.2021",Type:"income",Comments:"",Category: &Category{ReferenceBook1:"Book1",ReferenceBook2:"Book1"}})
	financialMoves=append(financialMoves,FinancialMove{ID:"2",Name:"Move2",Date:"12.10.2021",Type:"income",Comments:"",Category: &Category{ReferenceBook1:"Book2",ReferenceBook2:"Book2"}})
	r.HandleFunc("/financialMoves", getfinancialMoves).Methods("GET")
	r.HandleFunc("/financialMoves/{id}",getFinancialMove).Methods("GET")
	r.HandleFunc("/financialMoves", createFinancialMove).Methods("POST")
	r.HandleFunc("/financialMoves/{id}",updateFinancialMove).Methods("PUT")
	r.HandleFunc("/financialMoves/{id}",deleteFinancialMove).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}