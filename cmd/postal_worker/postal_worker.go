package main

import (
	"encoding/json"
	"github.com/m0cchi/postal_worker/lib/model"
	"io/ioutil"
	"log"
	"net/http"
)

// HandleAPI is sugoi
func HandleAPI(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println(err)
		return
	}

	// Unmarshal
	var postalMatter model.PostalMatter
	err = json.Unmarshal(b, &postalMatter)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(postalMatter)
}

func main() {
	log.Println("start postal workerd")

	http.HandleFunc("/api/register", HandleAPI)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
