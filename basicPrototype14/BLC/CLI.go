package BLC

import (
	"fmt"
	"os"
	"flag"
	"log"
)

type CLI struct {
	BlockChain *BlockChain
}

func (cli *CLI)addBlock(data string)  {
	cli.BlockChain.AddBlockchain(data)
	fmt.Println("Success!")
}

func (cli *CLI)printchain()  {
	cli.BlockChain.PrintChain()
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\taddBlock -data DATA -- 交易数据")
	fmt.Println("\tprintchain -- 输出区块信息")

}
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli CLI) Run() {
	isValidArgs()
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	flageAddBlockdata := addBlockCmd.String("data", "","Block Data")

	switch os.Args[1] {
	case "addBlock":

		if(addBlockCmd.Parse(os.Args[2:])!=nil){
			log.Panic(addBlockCmd.Parse(os.Args[2:]))
		}
	case "printchain":
		if(printChainCmd.Parse(os.Args[2:])!=nil){
			log.Panic(printChainCmd.Parse(os.Args[2:]))
		}

	default:
		printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *flageAddBlockdata == "" {
			printUsage()
			os.Exit(1)
		}

		//fmt.Println(*flageAddBlockdata)
		cli.addBlock(*flageAddBlockdata)

	}

	if printChainCmd.Parsed(){
		cli.printchain()
	}


}
