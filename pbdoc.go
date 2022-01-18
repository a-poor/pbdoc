package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Service struct {
	Name        string
	Description string
	Methods     []Method
}

type Method struct {
	Name         string
	Description  string
	Input        Message
	InputStream  bool
	Output       Message
	OutputStream bool
}

type Message struct {
	Description string
	Type        string
	Fields      []Field
}

type Field struct {
	Name        string
	Description string
	Type        string
	Required    bool
	Repeated    bool
}

func main() {
	http.HandleFunc("/pbdoc", func(w http.ResponseWriter, r *http.Request) {
		// Load the template
		raw, err := os.ReadFile("./static/index.html")
		if err != nil {
			log.Panic(err)
		}
		tmp, err := template.New("pbdoc").Parse(string(raw))
		if err != nil {
			log.Panic(err)
		}

		// Load the data
		buf, err := os.ReadFile("./example.json")
		if err != nil {
			log.Panic(err)
		}
		var data Service
		err = json.Unmarshal(buf, &data)
		if err != nil {
			log.Panic(err)
		}

		err = tmp.Execute(w, &data)
		if err != nil {
			log.Panic(err)
		}
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}

}
