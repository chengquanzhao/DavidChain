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
	fmt.Println("\tsend -from FromSomeone -to ToSomeone -amount AMOUT -- 交易明细")
	fmt.Println("\tprintchain -- 输出区块信息")
	fmt.Println("\tcreategensis -address DATA -- 交易数据")

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
	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	flagFrom := sendBlockCmd.String("from", "", "转账地址....")

	flagTo := sendBlockCmd.String("to", "", "转账目的地址....")

	flagAmount := sendBlockCmd.String("amount", "", "转账金额....")

	createBlockchainCmd := flag.NewFlagSet("creategensis", flag.ExitOnError)
	flagecreateBlockwithAddress := createBlockchainCmd.String("address", "", "创世创世区块的地址")
	switch os.Args[1] {
	case "send":

		if (sendBlockCmd.Parse(os.Args[2:]) != nil) {
			log.Panic(sendBlockCmd.Parse(os.Args[2:]))
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

	if sendBlockCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || "" == *flagAmount {
			printUsage()
			os.Exit(1)
		}

		cli.send(JsonToArry(*flagFrom), JsonToArry(*flagTo), JsonToArry(*flagAmount))

	}

	if printChainCmd.Parsed() {
		cli.printchain()
	}

	if createBlockchainCmd.Parsed() {
		if *flagecreateBlockwithAddress == "" {

			fmt.Println("地址不能为空.....")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesblockchain(*flagecreateBlockwithAddress)
	}
}

func (cli *CLI) createGenesblockchain(address string) {
	blockChain := CreateGenesisBlockChain(address)
	defer blockChain.DB.Close()
}

func (cli *CLI) send(from []string, to []string, amount []string) {
	if DBExists() == false {
		fmt.Println(" 数据不存在")
		os.Exit(1)
	}
	blockChain := BlockChainObject()
	blockChain.MineNewBlock(from, to, amount)
	defer blockChain.DB.Close()
}
