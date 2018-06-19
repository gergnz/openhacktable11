package main

import (
"encoding/json"
"fmt"
"html/template"
"log"
"net/http"
	"time"
	"io/ioutil"
	"bytes"
)

// http://13.75.219.113/list

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func root_handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "list: %s!", r.URL.Path[1:])

	var servers = GetServerInfo()
	tmpl := template.Must(template.ParseFiles("layout.html"))
	tmpl.Execute(w, servers)
//	fmt.Printf(tpl.String())
}

func add_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "add: %s!", r.URL.Path[1:])
}

func delete_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete: %s!", r.URL.Path[1:])
}



func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/add", add_handler)
	http.HandleFunc("/delete", delete_handler)
	http.HandleFunc("/list", root_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// [{"name": "hackerone", "endpoints": {"minecraft": "13.70.121.157:25565", "remote-console": "13.70.121.157:25575"}}]

type Endpoint struct {
	Minecraft string `json:"minecraft"`
	Rcon string `json:"remote-console"`
}

type Server struct {
	Name string `json:"name"`
	Endpoints Endpoint `json:"endpoints"`
}

func main3() {
	text := `[{"name": "hackerone", "endpoints": {"minecraft": "13.70.121.157:25565", "remote-console": "13.70.121.157:25575"}}]`
	var servers []Server
	json.Unmarshal([]byte(text), &servers)
	fmt.Printf("Servers : %+v", servers)
}


func GetServerInfo() [] Server {
	url := "http://13.75.219.113/list"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var servers []Server
	json.Unmarshal(body, &servers)
	fmt.Printf("Hello")
	fmt.Printf("Servers : %+v", servers)

	return servers
}

func main4() {
	GetServerInfo()
}

//{{.Endpoints.Minecraft}}, {{.Endpoints.Rcon}}

func test2() {
	var servers = GetServerInfo()
	tmpl := template.Must(template.ParseFiles("layout.html"))
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, servers)
	fmt.Printf(tpl.String())
}

func main5() {
	test2()
}