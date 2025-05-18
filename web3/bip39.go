package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	bip32 "github.com/tyler-smith/go-bip32"
	bip39 "github.com/tyler-smith/go-bip39"
)

func main() {
	var pharases []string
	var addresses []string
	for i := 0; i < 50; i++ {
		// 生成 128 位熵（对应 12 个助记词）
		entropy, err := bip39.NewEntropy(128)
		if err != nil {
			log.Fatal(err)
		}

		// 生成助记词
		mnemonic, err := bip39.NewMnemonic(entropy)
		if err != nil {
			log.Fatal(err)
		}

		// 从助记词生成种子
		seed := bip39.NewSeed(mnemonic, "")

		// 从种子派生出 master key
		masterKey, err := bip32.NewMasterKey(seed)
		if err != nil {
			log.Fatal(err)
		}

		// BIP44 路径 m/44'/60'/0'/0/0
		purposeKey, _ := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
		coinTypeKey, _ := purposeKey.NewChildKey(bip32.FirstHardenedChild + 60)
		accountKey, _ := coinTypeKey.NewChildKey(bip32.FirstHardenedChild + 0)
		changeKey, _ := accountKey.NewChildKey(0)
		addressKey, _ := changeKey.NewChildKey(0)

		privateKey, err := crypto.ToECDSA(addressKey.Key)
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public().(*ecdsa.PublicKey)
		address := crypto.PubkeyToAddress(*publicKey).Hex()

		// 打印助记词 + 地址（每行）
		pharases = append(pharases, mnemonic)
		addresses = append(addresses, address)
	}
	for _, v := range pharases {
		fmt.Println(v)
	}
	for _, v := range addresses {
		fmt.Println(v)
	}

}
