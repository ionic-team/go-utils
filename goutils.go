package goutils

import (
  "fmt"
  "runtime/debug"
  "os"
  "log"
  "net/http"
  "encoding/json"
  "html/template"
)

func CheckFatalErr(err error, msg string) {
  if err != nil {
    fmt.Fprintf(os.Stderr, msg)
    os.Exit(1)
  }
}

// Send a 400 - BadRequest error back to client
func Send400Json(w http.ResponseWriter, msg string) {
  w.WriteHeader(http.StatusBadRequest)
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w, "{\"error\": \"%s\"}", msg)
}

func Send500Json(w http.ResponseWriter, err error) {
  w.WriteHeader(http.StatusInternalServerError)
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w, "{\"error\": \"%s\"}", msg)
}

// Send a 500 - Server error error back to client
func Send500(w http.ResponseWriter, err error) {
  w.WriteHeader(http.StatusInternalServerError)
  fmt.Fprintf(os.Stderr, "Server error %s\n", err)
  debug.PrintStack()
  fmt.Fprintf(w, "Server error. We've been notified and will fix this, sorry!")
}

// Send a 404 - Not Found error back to client
func Send404(w http.ResponseWriter, err error) {
  w.Header().Set("Content-Type", "text/html")
  w.WriteHeader(http.StatusNotFound)
  if err != nil {
    //fmt.Fprintf(os.Stderr, "404 Not Found: %s\n", err)
  }

  errTemplate := template.Must(template.ParseFiles("templates/404.html"))

  errTemplate.Execute(w, map[string]string{
  })
}

func Send404Json(w http.ResponseWriter, msg string, err error) {
  w.WriteHeader(http.StatusNotFound)
  w.Header().Set("Content-Type", "application/json")
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error: %s", err)
  }
  fmt.Fprintf(w, "{\"error\": \"%s\"}", msg)
}

func JsonResponse(w http.ResponseWriter, t map[string]string, status int) {
  jsonObj, err := json.Marshal(t)
  if err != nil {
    log.Println("Unable to marshal response into JSON")
    Send500(w, err)
    return
  }

  w.WriteHeader(status)
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w, string(jsonObj))
}

func JsonResponseString(w http.ResponseWriter, json_string string, status int) {
  w.WriteHeader(status)
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(w, json_string)
}
