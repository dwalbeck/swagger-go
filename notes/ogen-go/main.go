package order_upload_ogen

import (
	"log"
	"net/http"

	api "order_upload_ogen/api"
)

func main() {
	// Create service instance.
	service := &petsService{
		pets: map[int64]petstore.Pet{},
	}
	// Create generated server.
	srv, err := api.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}

	// swagger-ui
	fs := http.FileServer(http.Dir("./ui"))
	http.Handle("/ui/", http.StripPrefix("/ui/", fs))

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
