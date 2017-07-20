package main

import	(
			"fmt"
			"net/http"
			"log"
			"os"
			"github.com/ethereum/go-ethereum/rpc"
			"time"
		)

type Block struct {
	Number *big.Int
}

func httpRequests(w http.ResponseWriter, r *http.Request)	{
	
	r.ParseForm()
	fmt.Fprintf(w, r.URL.Path +"<br>")

	switch r.URL.Path	{
		case "/exit":
			os.Exit(0)
		case "/test/blockchain":
			client, _ := rpc.Dial("ws://127.0.0.1:8485")
			subch := make(chan Block)
			// Ensure that subch receives the latest block.
	go func() {
		for i := 0; ; i++ {
			if i > 0 {
				time.Sleep(2 * time.Second)
			}
			subscribeBlocks(client, subch)
		}
	}()

	// Print events from the subscription as they arrive.
	for block := range subch {
		fmt.Println("latest block:", block.Number)
	}
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
	err := http.ListenAndServe(":8080", nil)
	if err != nil	{
		log.Fatal("ListenAndServe: ", err)
	}
}