package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	status := DBStatus()
	res := map[string]string{
		"mysql_status": status,
	}
	resJson, _ := json.Marshal(res)

	fmt.Fprintf(w, "[Database][Info]: %s", resJson)
}

func main() {
	http.HandleFunc("/healthy", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func DBStatus() string {
	var connectionString = fmt.Sprintf("root:toor@tcp(127.0.0.1:3306)/pcap?allowNativePasswords=true")

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return "ERROR"
	}

	err = db.Ping()
	if err != nil {
		return "ERROR"
	}

	return "OK"
}
