package randompics

import (   
	"net/http"	
	"fmt" 
	"encoding/json"
	"appengine"
	"appengine/datastore"
	"math/rand"	
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/submit", submitImg)
}

func handler(w http.ResponseWriter, r *http.Request) {
	
	c := appengine.NewContext(r)

	query := datastore.NewQuery("Images").Ancestor(createKey(c))

	count, err := query.Count(c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	offset := rand.Intn(count)

	imgs := make([]Image, 0, 1)

	if _, err := query.Offset(offset).Limit(1).GetAll(c, &imgs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	json, err := json.Marshal(imgs)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	fmt.Fprintf(w, "%s", json)	
}