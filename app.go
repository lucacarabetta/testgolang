package main

import	(
			"fmt"
			"net/http"
			"log"
		)

func httpRequests(w http.ResponseWriter, r *http.Request)	{
	r.ParseForm()
	fmt.Fprintf(w, r.URL.Path)
}

func main()	{
	http.HandleFunc("/", httpRequests)
	err := http.ListenAndServe(":8080", nil)
	if err != nil	{
		log.Fatal("ListenAndServe: ", err)
	}
}