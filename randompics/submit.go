package randompics

import (
	"net/http"
	"net/url"
	"fmt" 
	"encoding/json"   
	"appengine"
	"appengine/datastore"
)

func submitImg(w http.ResponseWriter, r *http.Request) {

	urlStr := r.FormValue("url")
	url, error := url.Parse(urlStr)

	if error != nil {
		fmt.Fprintf(w, "Invalid url %s %s", urlStr, error)
		return
	}

	img := Image{url.String()}

	c := appengine.NewContext(r)

	//TODO Look for duplicates
	
	key := datastore.NewIncompleteKey(c, "Images", createKey(c))

	_, err := datastore.Put(c, key, &img)

	json, err := json.Marshal(img)

	if err != nil {
		//Do nothing
	}

	fmt.Fprintf(w, "You submitted %s", json)
}

func createKey(c appengine.Context) *datastore.Key {        
    return datastore.NewKey(c, "Images", "default_images", 0, nil)
}