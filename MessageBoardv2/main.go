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

type WorkRequest struct {
	r* http.Request
	topicid string
	body string
}
var WorkerQueue chan chan WorkRequest

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 1000)

type MBoards []MessageBoard

type Worker struct {
  ID          int
  Work        chan WorkRequest
  WorkerQueue chan chan WorkRequest
  QuitChan    chan bool
}


// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
  // Create, and return the worker.
  worker := Worker{
    ID:          id,
    Work:        make(chan WorkRequest),
    WorkerQueue: workerQueue,
    QuitChan:    make(chan bool)}

  return worker
}

func returnOneTopic(w http.ResponseWriter, r* http.Request){
	vars := mux.Vars(r)
	topicid := vars["id"]
	fmt.Println("GET request for "+topicid)
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

func updateOneTopic(topicid string, body string){
	newtopic := "Content Created"
	db := connect()
	defer db.Close()
    /* Check if the content is available in Body */
	if len(strings.TrimSpace(body)) != 0 {
		fmt.Println("Post with data "+body)
		/* If the Body is empty, then create a dummy entry in the DB */
		stmt, _ :=db.Prepare("INSERT into MessageBoard (topicName, content, timestamp) values (?,?,NOW())")
		_, err := stmt.Exec(topicid,body)
		if err != nil {
			fmt.Println("DB Insert failed")
		} else {
			fmt.Println("DB Insert SUCCESS")
		}
	} else {
		/* Check if the content is available in Body */
		fmt.Println("Post without data")
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

func deleteOneTopic(topicid string, body string){
	db := connect()
	defer db.Close()
	/* Delete the given Topic from DB */
	stmt, _ := db.Prepare("DELETE FROM MessageBoard where topicName = ?")
	_,err := stmt.Exec(topicid)
	if err != nil {
		fmt.Println("Topic Deleted : "+topicid)
	} else {
		fmt.Println("Topic Deleted FAILED : "+topicid)
	}
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w *Worker) Start() {
	go func() {
		for {
		// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work

			select {
				case work := <-w.Work:
					// Receive a work request.
					if work.r.Method == "POST" {
						updateOneTopic(work.topicid, work.body)
					}
					if work.r.Method == "DELETE" {
						deleteOneTopic(work.topicid, work.body)
					}

				case <-w.QuitChan:
					// We have been asked to stop.
					fmt.Println("worker%d stopping\n", w.ID)
					return
        }
      }
    }()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
func webServiceRequests(w http.ResponseWriter, r *http.Request) {

	// Now, we take the delay, and the person's name, and make a WorkRequest out of them.
	vars := mux.Vars(r)
	topicid := vars["id"]
	decoder := json.NewDecoder(r.Body)
	var message MessageBoard
	decoder.Decode(&message)
	work := WorkRequest{r: r, topicid: topicid,body: message.Body }

	// Push the work onto the queue.
	WorkQueue <- work
	fmt.Println("Work request queued")

  return
}

func StartDispatcher(nworkers int) {
  // First, initialize the channel we are going to but the workers' work channels into.
  WorkerQueue = make(chan chan WorkRequest, nworkers)

  // Now, create all of our workers.
  for i := 0; i< nworkers; i++ {
	fmt.Println("Starting worker", i+1)
	worker := NewWorker(i+1, WorkerQueue)
	worker.Start()
  }

  go func() {
	for {
		select {
			case work := <-WorkQueue:
				fmt.Println("Received work requeust")
				go func() {
					worker := <-WorkerQueue
					fmt.Println("Dispatching work request")
					worker <- work
				}()
		}
	}
  }()
}

func handleHttpRequests(){
	myHttpRouter := mux.NewRouter().StrictSlash(true)

	StartDispatcher(1000)

	for _, route := range routes {
		myHttpRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.FunName)
	}


	log.Fatal(http.ListenAndServe(":9900",myHttpRouter))

}
func main() {
	handleHttpRequests()
}
