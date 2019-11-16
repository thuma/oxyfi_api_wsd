package main

import (
    "time"
    "fmt"
    "strconv"
    "encoding/json"
    "github.com/google/uuid"
    "strings"
    "net/http"
)

func rest_auth( next http.HandlerFunc ) http.HandlerFunc {
    return func( w http.ResponseWriter, r *http.Request) {
        user, pass, _ := r.BasicAuth()
        if (user == admin_username && pass == admin_password) {
            next(w, r)
        } else {
            w.Header().Set("WWW-Authenticate", `Basic realm="Trafiklav API 1.12 Auth"`)
            w.WriteHeader(401)
            w.Write([]byte("Unauthorised.\n"))
        }
    }
}

func ApiResonse(w http.ResponseWriter, data_to_json interface{} ){
    default_response.ResponseData = data_to_json
    json_data, _ := json.Marshal(default_response)
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(json_data))
}

func profilesHandler(w http.ResponseWriter, r *http.Request) {
    ApiResonse(w, default_profiles)
}

func adminKeysHandler(w http.ResponseWriter, r *http.Request) {
    ApiResonse(w, all_keys())
}

func keysHandler(w http.ResponseWriter, r *http.Request) {
    var this_key TrafiklabKey
    path := strings.Split(r.URL.Path, "/")
    key_id := path[len(path)-1]
    this_key = get_key(key_id)
    var created_time = time.Now().UTC().Format(time.RFC3339)
    if r.Method == http.MethodPost {
        decoder := json.NewDecoder(r.Body)
        var post_data NewKey
        err := decoder.Decode(&post_data)
        if err != nil {
            panic(err)
        }
        this_key.Key = strings.ReplaceAll(uuid.New().String(),"-","")
        this_key.Note = post_data.Note
        this_key.Project = post_data.Project.Id
        this_key.CreatedDate = created_time
        this_key.UpdatedDate = created_time
        this_key.Active = true
        add_key(this_key)
    } else if r.Method == http.MethodPut {
        decoder := json.NewDecoder(r.Body)
        var post_data TrafiklabKey
        err := decoder.Decode(&post_data)
        if err != nil {
            panic(err)
        }
        update_key(key_id, post_data)
        this_key = get_key(key_id)
    } else if r.Method == http.MethodDelete {
        delete_key(key_id)
    }
    ApiResonse(w, this_key)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Connected count: %s", strconv.Itoa(len(listeners)))
}