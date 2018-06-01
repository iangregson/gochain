package main

import (
	"flag"
	"github.com/iangregson/gochain/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "add a new block to chain")
	listFlag := flag.Bool("list", false, "get list of blocks on the chain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	client = proto.NewBlockchainClient(conn)

	if *addFlag {
		addBlock()
	}

	if *listFlag {
		listBlocks()
	}
}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatalf("unable to add block %v", err)
	}
	log.Printf("new block hash: %s\n", block.Hash)
}

func listBlocks() {
	bc, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("unable to get the blockchain %v", err)
	}
	log.Println("Blocks on the chain:")
	for _, b := range bc.Blocks {
		log.Printf("hash: %s, prev block hash: %s, data: %s", b.Hash, b.PrevBlockHash, b.Data)
	}
}
