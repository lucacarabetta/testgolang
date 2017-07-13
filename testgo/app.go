package main

import	(
			"fmt"
			"bufio"
			"os"
		)

func main()	{
	for true	{
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		fmt.Printf(input)
	}
}