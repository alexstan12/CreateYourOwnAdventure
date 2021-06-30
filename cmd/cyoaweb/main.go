package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alexstan12/cyoa"
)
func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web app on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)	
	}

	story, err := cyoa.JsonStory(f)
	if err !=nil {
		panic(err)
	}
	//fmt. Printf("%+v\n", story)
	//tpl:= template.Must(template.New("").Parse("Hello!"))
	//handler:= cyoa.NewHandler(story, cyoa.WithTemplate(tpl)) - usage for custom template
	handler:=cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port),handler))
}


