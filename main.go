package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/nomadcoders/nomadcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// func main() {
// 	chain := blockchain.GetBlockchain()
// 	chain.AddBlock("Second Block")
// 	chain.AddBlock("Third Block")
// 	chain.AddBlock("Fourth Block")
// 	for _, block := range chain.AllBlocks() {
// 		fmt.Printf("Data: %s\n", block.Data)
// 		fmt.Printf("Hash: %s\n", block.Hash)
// 		fmt.Printf("Prev Hash: %s\n", block.PrevHash)
// 	}
// }
