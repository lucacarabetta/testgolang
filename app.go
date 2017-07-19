package main

import	(
			"fmt"
			"net/http"
			"log"
			"os"
		)

func httpRequests(w http.ResponseWriter, r *http.Request)	{
	
	r.ParseForm()
	fmt.Fprintf(w, r.URL.Path)

	switch r.URL.Path	{
		case "/exit":
			os.Exit(0)
		case "/send/ether":
			fmt.Fprintf(w, "Sending Ether from"+ eth.coinbase +"to"+ eth.accounts[1])
			eth.sendTransaction({from:eth.coinbase, to:eth.accounts[1], value: web3.toWei(0.05, "ether")})
	}
}

func connectToBlockchain()	{

}

func printBlockchain()	{

}

func sendValue()	{
}

func main()	{
	http.HandleFunc("/", httpRequests)
	err := http.ListenAndServe(":9990", nil)
	if err != nil	{
		log.Fatal("ListenAndServe: ", err)
	}
}