package instagram

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "appengine"
    "appengine/urlfetch"
)

func Fetch(w http.ResponseWriter, r *http.Request) {
    var accessToken = "332912852.1c5d48a.2497d322676b40fe91ffb96ed6729eca"
	c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    resp, err := client.Get("https://api.instagram.com/v1/tags/search?q=relax&access_token=" + accessToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    body, err := ioutil.ReadAll(resp.Body);
    fmt.Fprintf(w, "%s", body)
}