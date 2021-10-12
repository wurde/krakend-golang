package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func PostServer() {
	fmt.Printf("Starting Post server at port 5040\n")
	r := mux.NewRouter()

	// /
	// curl -i http://localhost:5040
	r.Handle("/", http.RedirectHandler("/posts", http.StatusMovedPermanently))

	// /posts
	r.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		// GET /posts
		// curl -i http://localhost:5040/posts
		if r.Method == "GET" {
			res := make(map[string]interface{})
			posts := make([]map[string]string, 2)
			posts[0] = map[string]string{"id": "1", "title": "Lorem Ipsum1"}
			posts[1] = map[string]string{"id": "2", "title": "Lorem Ipsum2"}
			res["posts"] = posts
			jsonRes, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
		}

		// POST /posts
		// curl -i -X POST --data-raw '{"title":"Lorem Ipsum2"}' http://localhost:5040/posts
		if r.Method == "POST" {
			var m map[string]string
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err.Error())
			}
			json.Unmarshal(body, &m)
			fmt.Printf("%s", m)

			res := map[string]string{"message": "Successfully created post."}
			jsonRes, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}

			w.WriteHeader(http.StatusCreated)
			w.Write(jsonRes)
		}
	})

	// /posts/{id}
	r.HandleFunc("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			fmt.Println("id is missing in parameters")
		}

		// GET /posts/{id}
		// curl -i http://localhost:5040/posts/1
		if r.Method == "GET" {
			fmt.Printf(`GET /posts/%s`, id)

			res := map[string]string{"id": "1", "title": "Lorem Ipsum1"}
			jsonRes, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
		}

		// PATCH /posts/{id}
		// curl -i -X PATCH --data-raw '{"title":"Lorem Ipsum3"}' http://localhost:5040/posts/1
		if r.Method == "PATCH" {
			fmt.Printf(`PATCH /posts/%s`, id)

			res := map[string]string{"id": "1", "title": "Lorem Ipsum3"}
			jsonRes, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
		}

		// DELETE /posts/{id}
		// curl -i -X DELETE http://localhost:5040/posts/1
		if r.Method == "DELETE" {
			fmt.Printf(`DELETE /posts/%s`, id)

			res := map[string]string{"message": "Successfully deleted post."}
			jsonRes, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
		}
	})

	log.Fatal(http.ListenAndServe(":5040", r))
}

func CommentServer() {
	fmt.Printf("Starting Comment server at port 5050\n")
	r := mux.NewRouter()

	// /
	// curl -i http://localhost:5050
	r.Handle("/", http.RedirectHandler("/comments", http.StatusMovedPermanently))

	// /comments
	r.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		// GET /comments
		// curl -i http://localhost:5050/comments
		if r.Method == "GET" {
			res := make(map[string]interface{})
			comments := make([]map[string]string, 2)
			comments[0] = map[string]string{"id": "1", "postId": "1", "message": "First comment!!!"}
			comments[1] = map[string]string{"id": "2", "postId": "1", "message": "Beat me by 1ns."}
			res["comments"] = comments
			jsonRes, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
		}

		// POST /comments
		// curl -i -X POST --data-raw '{"postId": 2, "message":"Hello?"}' http://localhost:5050/comments
		if r.Method == "POST" {
			var m map[string]string
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err.Error())
			}
			json.Unmarshal(body, &m)
			fmt.Printf("%s", m)

			res := map[string]string{"message": "Successfully created comment."}
			jsonRes, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}

			w.WriteHeader(http.StatusCreated)
			w.Write(jsonRes)
		}
	})

	// /comments/{id}
	r.HandleFunc("/comments/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			fmt.Println("id is missing in parameters")
		}

		// GET /comments/{id}
		// curl -i http://localhost:5050/comments/1
		if r.Method == "GET" {
			fmt.Printf(`GET /comments/%s`, id)

			res := map[string]string{"id": "1", "postId": "1", "message": "First comment!!!"}
			jsonRes, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
		}

		// PATCH /comments/{id}
		// curl -i -X PATCH --data-raw '{"message":"[redacted]"}' http://localhost:5050/comments/1
		if r.Method == "PATCH" {
			fmt.Printf(`PATCH /comments/%s`, id)

			res := map[string]string{"id": "1", "postId": "1", "message": "[redacted]"}
			jsonRes, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
		}

		// DELETE /comments/{id}
		// curl -i -X DELETE http://localhost:5050/comments/1
		if r.Method == "DELETE" {
			fmt.Printf(`DELETE /comments/%s`, id)

			res := map[string]string{"message": "Successfully deleted comment."}
			jsonRes, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
		}
	})

	log.Fatal(http.ListenAndServe(":5050", r))
}

func main() {
	go PostServer()
	CommentServer()
}
