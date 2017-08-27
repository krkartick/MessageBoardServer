package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
)

type HttpResp struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

type MessageBoard struct {
	Topic string `json:"Topic"`
	Body string `json:"Body"`
}

type MBoards []MessageBoard

func returnOneTopic(w http.ResponseWriter, r* http.Request){
	vars := mux.Vars(r)
	topicid := vars["id"]
	w.Header().Set("Content-Type","application/json")
	db := connect()
	defer db.Close()
	var allMessages []string
	/* Query DB for the given TOPIC */
	results, err := db.Query("SELECT content from MessageBoard where topicName = ?", topicid)
	if err == nil {
		allMessages = append(allMessages,fmt.Sprintf("Topic: %s",topicid))
	}
	count := 0
	/* Loop through all the results from DB */
	for results.Next() {
		var message MessageBoard
		err := results.Scan(&message.Body)
		if err != nil {
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Topic Not Found"})
		}
		count += 1
		allMessages = append(allMessages,message.Body)
	}
	if count == 0 {
		allMessages = append(allMessages,fmt.Sprintf("Content Not Found"))
	}
	/* Encode the Query response in Json format */
	json.NewEncoder(w).Encode(allMessages)
}

func checkRowCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println("DB Row scan failed")
		}
	}
	return count
}

func updateOneTopic(w http.ResponseWriter, r* http.Request){
	db := connect()
	defer db.Close()
	decoder := json.NewDecoder(r.Body)
	vars := mux.Vars(r)
	topicid := vars["id"]
	newtopic := "Content Created"
	var message MessageBoard
	decoder.Decode(&message)
    /* Check if the content is available in Body */
	if len(strings.TrimSpace(message.Body)) != 0 {
		fmt.Println("Post with data "+message.Body)
		fmt.Fprintf(w,"Update Topic Name: "+topicid)
		/* If the Body is empty, then create a dummy entry in the DB */
		stmt, _ :=db.Prepare("INSERT into MessageBoard (topicName, content, timestamp) values (?,?,NOW())")
		_, err := stmt.Exec(topicid,message.Body)
		if err != nil {
			fmt.Println("DB Insert failed")
		} else {
			fmt.Println("DB Insert SUCCESS")
		}
	} else {
		/* Check if the content is available in Body */
		fmt.Println("Post without data")
		fmt.Fprintf(w,"Insert Topic Name: "+topicid)
		DBRows, _ := db.Query("SELECT count(*) FROM MessageBoard where topicName = ?", topicid)
		if checkRowCount(DBRows) > 0 {
			fmt.Println("Topic already created for: "+topicid)

		} else {
			stmt, _ :=db.Prepare("INSERT into MessageBoard (topicName, content, timestamp) values (?,?,NOW())")
			_, err := stmt.Exec(topicid,newtopic)
			if err != nil {
				fmt.Println("DB Insert failed")
			} else {
				fmt.Println("DB Insert SUCCESS")
			}
		}
	}
}
func deleteOneTopic(w http.ResponseWriter, r* http.Request){
	db := connect()
	defer db.Close()
	vars := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	topicid := vars["id"]
        fmt.Fprintf(w,"Delete Topic Name: "+topicid)
	var message MessageBoard
	decoder.Decode(&message)
	/* Delete the given Topic from DB */
	stmt, _ := db.Prepare("DELETE FROM MessageBoard where topicName = ?")
	_,err := stmt.Exec(topicid)
	if err != nil {
		fmt.Println("Topic Deleted : "+topicid)
	} else {
		fmt.Println("Topic Deleted FAILED : "+topicid)
	}
}
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"! Welcome to Message Board !")
	fmt.Println(" Thank You !!! ")

}
func handleHttpRequests(){
	myHttpRouter := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		myHttpRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.FunName)
	}
	log.Fatal(http.ListenAndServe(":9000",myHttpRouter))

}
func main() {
	handleHttpRequests()
}
