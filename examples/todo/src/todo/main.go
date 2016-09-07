package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todo/store"
	"todo/views"
)

func log(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		views.Index(w, store.TodoStore.List())
		log("[INF]enter index page")
	})
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			views.Add(w)
			log("[INF]enter add page")
			return
		}
		finishAt, err := time.Parse("2006-01-02", r.FormValue("finat"))
		if err != nil {
			log("[ERR]add new todo fail: %s", err.Error())
			return
		}
		rec := store.TodoRec{
			Title:    r.FormValue("title"),
			Content:  r.FormValue("content"),
			Done:     false,
			CreateAt: time.Now(),
			FinishAt: finishAt,
		}
		store.TodoStore.Add(&rec)
		log("[INF]add new todo ok")
		http.Redirect(w, r, "/", 302)
	})
	http.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {
		recId, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			return
		}
		index, rec := store.TodoStore.FindById(recId)
		if index < 0 {
			return
		} else if r.Method == "GET" {
			views.Edit(w, rec)
		} else {
			finishAt, err := time.Parse("2006-01-02", r.FormValue("finat"))
			if err != nil {
				log("[ERR]add new todo fail: %s", err.Error())
				return
			}
			rec.Title = r.FormValue("title")
			rec.Content = r.FormValue("content")
			rec.FinishAt = finishAt
			log("[INF]edit todo ok")
			http.Redirect(w, r, "/", 302)
		}
	})
	log("Todo server started")
	http.ListenAndServe(":8080", nil)
}
