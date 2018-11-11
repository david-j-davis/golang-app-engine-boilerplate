// [START gae_go111_app]

// Sample helloworld is an App Engine app.
package main

// [START import]
import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// [END import]
// [START main]

func main() {
	http.HandleFunc("/hello", helloHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	// [END setting_port]
}

// [END main]

// [START helloHandler]

// helloHandler responds to requests with our greeting.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello from the Go app!")
}

// [END helloHandler]
// [END gae_go111_app]