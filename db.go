package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "time"
)

var database *sql.DB

func init_db(){
    database, _ = sql.Open("sqlite3", db_path)
    statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS keys (key TEXT PRIMARY KEY, note TEXT, project TEXT, created TEXT, updated TEXT, active INT)")
    if err != nil {
        log.Fatal(err)
    }
    statement.Exec()
}

func add_key( keydata TrafiklabKey ) {
    statement, err := database.Prepare("INSERT INTO keys (key, note, project, created, updated, active) VALUES (?, ?, ?, ? ,?, ?)")
    if err != nil {
        log.Print(err)
    }
    statement.Exec(keydata.Key, keydata.Note, keydata.Project, keydata.CreatedDate, keydata.UpdatedDate, 1)
}

func delete_key(key_id string) {
    statement, err := database.Prepare("DELETE FROM keys WHERE key = ? LIMIT 1")
    if err != nil {
        log.Print(err)
    }
    statement.Exec()
}

func update_key(key_id string, keydata TrafiklabKey ) {
    statement, err := database.Prepare("UPDATE keys SET note = ?, updated = ? WHERE key = ? LIMIT 1")
    if err != nil {
        log.Print(err)
    }
    statement.Exec(keydata.Note,time.Now().UTC().Format(time.RFC3339), key_id)
}

func get_key(key string) TrafiklabKey {
    rows, err := database.Query("SELECT key, note, project, created, updated, active FROM keys where key == ? ", key)
    if err != nil {
        log.Print(err)
    }
    for rows.Next() {
        var key, note, project, created, updated string
        var active int
        rows.Scan(&key, &note, &project, &created, &updated, &active)
        return TrafiklabKey{
            Key:key,
            Note:note,
            Api:"oxygps",
            Profile:"One",
            Project:project,
            CreatedDate:created,
            UpdatedDate:updated,
            Active:active > 0,
        }
    }
    return TrafiklabKey{
        Key:"",
        Note:"",
        Api:"oxygps",
        Profile:"One",
        Project:"",
        CreatedDate:"",
        UpdatedDate:"",
        Active:false,
    }
}

func all_keys() TrafiklabKeys {
    var keysdata TrafiklabKeys
    var key, note, project, created, updated string
    rows, err := database.Query("SELECT key, note, project, created, updated FROM keys")
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        rows.Scan(&key, &note, &project, &created, &updated )
        keysdata = append(keysdata, TrafiklabKey{
                Key:key,
                Note:note,
                Api:"oxygps",
                Profile:"One",
                Project:project,
                CreatedDate:created,
                UpdatedDate:updated,
                Active:true,
            })
    }
    return keysdata
}