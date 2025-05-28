package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

// 合约地址
var contractAddress = common.HexToAddress("0x78606Ea26275F6180745480FaF74E311586EB652")

// 合约 ABI
const contractABI = `[{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"data","type":"bytes"}],"name":"transfer","outputs":[],"payable":true,"stateMutability":"payable","type":"function"}]`

// 目标地址列表
var targets = []string{
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
	//arb from sepolia
	execute(false, "https://rpc.ankr.com/eth_sepolia/d58d3df73575c1439d66f04b0b524730f45de8e22704c0ade9d60f2c9f301c73", 11155111, "0xb5aadef97d81a77664fcc3f16bfe328ad6cec7ac", 9535)
	//monad	from sepolia
	go func() {
		execute(false, "https://rpc.ankr.com/eth_sepolia/d58d3df73575c1439d66f04b0b524730f45de8e22704c0ade9d60f2c9f301c73", 11155111, "233e416b0897e8f4796d89a84b5ade4365ed566c", 9596)
	}()
	//bsc  ok 收续费 0.01
	go func() {
		execute(true, "https://opbnb-testnet-rpc.bnbchain.org", 5611, "0xb5aadef97d81a77664fcc3f16bfe328ad6cec7ac", 9515)
	}()
	////sepolia  from arb  checked
	go func() {
		execute(false, "https://rpc.ankr.com/arbitrum_sepolia/d58d3df73575c1439d66f04b0b524730f45de8e22704c0ade9d60f2c9f301c73", 421614, "0xb5aadef97d81a77664fcc3f16bfe328ad6cec7ac", 9526)
	}()

	//base  ok from sepolia  checked
	go func() {
		execute(false, "https://rpc.ankr.com/eth_sepolia/d58d3df73575c1439d66f04b0b524730f45de8e22704c0ade9d60f2c9f301c73", 11155111, "233e416b0897e8f4796d89a84b5ade4365ed566c", 9521)
	}()

	time.Sleep(9 * time.Hour)
}

func execute(isBsc bool, rpcURL string, chainId uint, bridgeAddr string, chainNum uint) {
	// Sepolia 测试网 RPC 端点（需替换为你的实际端点）
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		panic(fmt.Sprintf("无法连接 RPC 端点: %v", err))
	}

	// 助记词
	mnemonic := "todo"
	privateKey, err := mnemonicToPrivateKey(mnemonic)
	if err != nil {
		panic(fmt.Sprintf("恢复私钥失败: %v", err))
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 检查余额
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		panic(fmt.Sprintf("获取余额失败: %v", err))
	}
	balanceEth := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("主地址 %s 的余额: %s ETH\n", fromAddress.Hex(), balanceEth.Text('f', 6))

	// 固定 to 地址
	fixedToAddress := common.HexToAddress(bridgeAddr)

	// 循环向每个目标地址发送交易
	for _, target := range targets {
		rand.New(rand.NewSource(time.Now().UnixNano()))

		minMinutes := 20
		maxMinutes := 40
		randomMinutes := minMinutes + rand.Intn(maxMinutes-minMinutes+1) // [20, 40]
		sleepDuration := time.Duration(randomMinutes) * time.Minute
		time.Sleep(sleepDuration)

		// 生成随机金额（0.1 到 0.15 ETH）
		var a, b float64
		if isBsc {
			a = 0.01
			b = 0.005
		} else {
			a = 0.1
			b = 0.05
		}
		amount := big.NewFloat(a + rand.Float64()*b) // 0.1 + [0, 0.05)
		amountWei := new(big.Int)
		amountWei, _ = amount.Mul(amount, big.NewFloat(1e18)).Int(amountWei)

		// 构建交易
		tx, err := buildTransaction(chainId, chainNum, client, privateKey, contractAddress, fixedToAddress, common.HexToAddress(target), amountWei)
		if err != nil {
			fmt.Printf("构建交易失败（目标地址 %s）: %v\n", target, err)
			continue
		}

		// 发送交易
		err = sendTransaction(client, privateKey, tx)
		if err != nil {
			fmt.Printf("%d发送交易失败（目标地址 %s）: %v\n", chainId, target, err)
			continue
		}

		fmt.Printf("成功发送交易，to: %s，data 中目标地址: %s，金额: %s ETH\n", fixedToAddress.Hex(), target, amount.Text('f', 6))

	}
	fmt.Printf("chainNum: %d\n done", chainNum)
}

// 从助记词恢复私钥
func mnemonicToPrivateKey(mnemonic string) (*ecdsa.PrivateKey, error) {
	// 生成种子
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, fmt.Errorf("生成种子失败: %v", err)
	}

	// 使用 BIP-32 生成主密钥
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, fmt.Errorf("生成主密钥失败: %v", err)
	}

	// 使用 BIP-44 路径派生密钥：m/44'/60'/0'/0/0
	path, err := accounts.ParseDerivationPath("m/44'/60'/0'/0/0")
	if err != nil {
		return nil, fmt.Errorf("解析派生路径失败: %v", err)
	}

	key := masterKey
	for _, n := range path {
		key, err = key.NewChildKey(uint32(n))
		if err != nil {
			return nil, fmt.Errorf("派生子密钥失败: %v", err)
		}
	}

	// 转换为 ECDSA 私钥
	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return nil, fmt.Errorf("转换为 ECDSA 私钥失败: %v", err)
	}

	return privateKey, nil
}

// 构建交易
func buildTransaction(chainId uint, chainNum uint, client *ethclient.Client, privateKey *ecdsa.PrivateKey, contractAddress, fixedToAddress, targetAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	// 获取建议的 Gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	// 解析 ABI
	abiObj, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, err
	}

	// 创建合约实例
	contract := bind.NewBoundContract(contractAddress, abiObj, client, client, client)

	// 动态生成 data 参数：c=9535&t=<目标地址>
	dataStr := fmt.Sprintf("c=%d&t=%s", chainNum, targetAddress.Hex())
	data := []byte(dataStr)

	// 设置交易参数
	opts := &bind.TransactOpts{
		From:     fromAddress,
		Nonce:    big.NewInt(int64(nonce)),
		Value:    amount,
		GasLimit: 60000, // 可根据需要调整
		GasPrice: gasPrice,
		Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(chainId))), privateKey) // 11155111 是 Sepolia 的链 ID
		},
	}

	// 调用 transfer 函数
	tx, err := contract.Transact(opts, "transfer", fixedToAddress, data)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// 发送交易
func sendTransaction(client *ethclient.Client, privateKey *ecdsa.PrivateKey, tx *types.Transaction) error {
	err := client.SendTransaction(context.Background(), tx)
	if err != nil {
		return err
	}
	return nil
}
