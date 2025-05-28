package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
	"math/big"
	"strings"

	//"github.com/ethereum/go-ethereum/accounts"
	//"github.com/ethereum/go-ethereum/accounts/abi"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/crypto"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 目标地址列表
var accountList = []string{
	"0xf02D957658D8C836c9240545122BE6D168713Db6",
	"0xA16C56e87ff0cf44B1D843DbcA11eb7F348839A7",
	"0x117A4D82908F5d495943A588DA671d90909f8107",
	"0xEBAA0B86355C830790a49D149CcB80ad81a52266",
	"0x329ADD75074Aa6f2189D4BD32b2eF663C854D88A",
	"0x832f0AA31cebfd43904EEa8Bc0451fa72A3b8Ad9",
	"0xfff7Df895dA2AcB212235fA940eDbcaFd0a3C56f",
	"0x0100505Cb84444022317C3E663DFA3F7b80Dee86",
	"0xAF2139fd81623172D6e8Dd7496D07a766f4a1143",
	"0x0B665EC35117cFa88293F40Ba2f46218955b7a7C",
	"0x5A8b7fFAFb90cf20e14C862F7E44FB8A182A5B2F",
	"0xA977E4A7B803FAEEa8DD0557457c142D91B7FB64",
	"0x34b8147a6B3e8b8a7C29A71008Fc9d623a90490A",
	"0xdc1a9090B2DBC9d99E381f41B877c36c3BC9d2d0",
	"0xa4214a6D6386aF0F319255F60Eed83dEdb58B1b4",
}

func main() {
	sepoliaClient, _ := ethclient.Dial("https://rpc.ankr.com/eth_sepolia/d58d3df73575c1439d66f04b0b524730f45de8e22704c0ade9d60f2c9f301c73")
	baseSepoliaClient, _ := ethclient.Dial("https://rpc.ankr.com/base_sepolia/d58d3df73575c1439d66f04b0b524730f45de8e22704c0ade9d60f2c9f301c73")
	arbSepoliaClient, _ := ethclient.Dial("https://rpc.ankr.com/arbitrum_sepolia/d58d3df73575c1439d66f04b0b524730f45de8e22704c0ade9d60f2c9f301c73")
	monadSepoliaClient, _ := ethclient.Dial("https://testnet-rpc.monad.xyz")

	for _, element := range accountList {
		//sepolia
		balance := checkBalance(sepoliaClient, element)
		fmt.Println("address:", element, "balance:", balance.Int64(), "chain:eth_sepolia")
		//base sepolia
		balance = checkBalance(baseSepoliaClient, element)
		fmt.Println("address:", element, "balance:", balance.Int64(), "chain:base_sepolia")
		//arb sepolia
		balance = checkBalance(arbSepoliaClient, element)
		fmt.Println("address:", element, "balance:", balance.Int64(), "chain:arbitrum_sepolia")
		//monad sepolia
		balance = checkETHTokenBalance(monadSepoliaClient, element)
		fmt.Println("address:", element, "balance:", balance.Int64(), "chain:monad")
	}

}
func checkBalance(client *ethclient.Client, address string) big.Int {

	balanceAt, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		fmt.Println(err)
	}

	return *balanceAt

}
func checkETHTokenBalance(client *ethclient.Client, address string) big.Int {
	//调用（0x836047a99e11f376522b447bffb6e3495dd0637c）（erc20）的balanceof方法
	contractAddress := common.HexToAddress("0x836047a99e11f376522b447bffb6e3495dd0637c")
	userAddress := common.HexToAddress(address)
	// ERC20 ABI（标准）
	erc20ABI := `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],
	"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],
	"type":"function"}]`

	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		log.Fatal("ABI解析失败:", err)
	}
	// 构造data
	data, err := parsedABI.Pack("balanceOf", userAddress)
	if err != nil {
		log.Fatal("ABI调用打包失败:", err)
	}

	// 调用合约
	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Fatal("合约调用失败:", err)
	}

	// 解析返回值
	var balance = new(big.Int)
	balance.SetBytes(result)

	return *balance

}
