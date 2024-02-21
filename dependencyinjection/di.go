package dependencyinjection

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func OsExample() {
	Greet(os.Stdout, "Charlie")
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func LocalhostExample() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreeterHandler)))
}
