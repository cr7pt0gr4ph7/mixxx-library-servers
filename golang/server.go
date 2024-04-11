package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// We need the following features:
	//    - Search inside of whole database
	//    - Search inside of specific track collection (aka. playlist aka. crate aka. ...)
	//      => Both of these should probably use the same endpoint
	//    - Get list of track collections (aka. playlists aka. crates aka. ...)
	//
	// And maybe:
	//    - Edit track metadata
	http.HandleFunc("/tracks", func(w http.ResponseWriter, r *http.Request) {
		collectionId := r.FormValue("collection") // Optional. If not specified, search whole database.
		searchQuery := r.FormValue("search")      // Optional. If not specified, get all tracks.

		log.Println("GET TRACKS", "collection:", collectionId, "search:", searchQuery)
		io.WriteString(w, `{"tracks":[]}`)
	})

	http.HandleFunc("/collections", func(w http.ResponseWriter, r *http.Request) {
		collectionId := r.FormValue("id") // The id of the collection

		log.Println("GET COLLECTIONS", collectionId)
		io.WriteString(w, `{"collections":[]}`)
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
