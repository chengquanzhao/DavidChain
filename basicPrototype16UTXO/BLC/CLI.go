package BLC

import (
	"fmt"
	"os"
	"flag"
	"log"
)

type CLI struct {
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\taddBlock -data DATA -- 交易数据")
	fmt.Println("\tprintchain -- 输出区块信息")
	fmt.Println("\tcreategensis -address DATA -- 交易数据")

}

func (cli *CLI) addBlock(txs []*Transaction) {
	if DBExists() == false {
		fmt.Println("数据不存在")
		os.Exit(1)
	}
	blockchain := BlockChainObject()
	defer blockchain.DB.Close()

	blockchain.AddBlockchain(txs)
	fmt.Println("Success!")
}

func (cli *CLI) printchain() {
	if DBExists() == false {
		fmt.Println("数据不存在")
		os.Exit(1)
	}
	blockchain := BlockChainObject()
	defer blockchain.DB.Close()

	blockchain.PrintChain()
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
	flageAddBlockdata := addBlockCmd.String("data", "", "Block Data")
	createBlockchainCmd := flag.NewFlagSet("creategensis", flag.ExitOnError)
	flagecreateBlockwithAddress := createBlockchainCmd.String("address", "", "创世创世区块的地址")
	switch os.Args[1] {
	case "addBlock":

		if (addBlockCmd.Parse(os.Args[2:]) != nil) {
			log.Panic(addBlockCmd.Parse(os.Args[2:]))
		}
	case "printchain":
		if (printChainCmd.Parse(os.Args[2:]) != nil) {
			log.Panic(printChainCmd.Parse(os.Args[2:]))
		}
	case "creategensis":

		if (createBlockchainCmd.Parse(os.Args[2:]) != nil) {
			log.Panic(createBlockchainCmd.Parse(os.Args[2:]))
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
		cli.addBlock([]*Transaction{})

	}

	if printChainCmd.Parsed() {
		cli.printchain()
	}

	if createBlockchainCmd.Parsed() {
		if *flagecreateBlockwithAddress== "" {

			fmt.Println("地址不能为空.....")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesblockchain(*flagecreateBlockwithAddress)
	}
}

func (cli *CLI) createGenesblockchain(address string) {
	CreateGenesisBlockChain(address)
}
