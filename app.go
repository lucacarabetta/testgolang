package main

import	(
			"fmt"
			"net/http"
			"log"
		)

func httpRequests(w http.ResponseWriter, r *http.Request)	{
	r.ParseForm()
	fmt.Println(w, r.Form)
}

func main()	{
	http.HandleFunc("/", httpRequests)
	err := http.ListenAndServe(":8080", nil)
	if err != nil	{
		log.Fatal("ListenAndServe: ", err)
	}
}