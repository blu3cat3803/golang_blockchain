package main

import (
	"fmt"
	"bufio"
	"os"
	bc "mypackage/blockchain"
)

func Help() {
	fmt.Println("There are 3 operations:")
	fmt.Println("Type 1 for adding a new Block")
	fmt.Println("Type 2 for printing the Blockchain")
	fmt.Println("Type 3 for exiting")
}

func main() {
	fmt.Println("Welcome to our Blockchain project.")
	fmt.Println("Enter h for help")

	var (
		op string
	)

	NewBlockchain := bc.CreateBlockchain()
	for true {
		fmt.Scanln(&op)
		if op == "h" {
			fmt.Println("Printing the help")
			Help()
		} else if op == "1" {
			fmt.Println("Entering your data:")
			reader := bufio.NewReader(os.Stdin)
			data, _, _:= reader.ReadLine()
			NewBlockchain.AddBlock(data)
		} else if op == "2"{
				for _ , block := range NewBlockchain.Blocks {
					fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
			        fmt.Printf("Data: %s\n", block.Data)
			        fmt.Printf("Hash: %x\n", block.Hash)
					//pow := bc.NewProofOfWork(block)
			        //fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
			        fmt.Println()
				}
		} else if op == "3"{
				break
		} else {
			fmt.Println("Please Enter h, 1, 2, 3")
		}
	}

}