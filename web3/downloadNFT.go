package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//download nft pngs from Opensea-Azuki
	//define params : tokenIds
	tokenId := 5534
	totalPicNum := 50
	client := &http.Client{}
	for i := 0; i < totalPicNum; i++ {
		//request for ipfs uri

		reqBody := "{\n  \"operationName\": \"ItemViewModalQuery\",\n  \"query\": \"query ItemViewModalQuery($identifier: ItemIdentifierInput!) {\\n  itemByIdentifier(identifier: $identifier) {\\n    __typename\\n    ... on Item {\\n      id\\n      name\\n      ...ItemView\\n      __typename\\n    }\\n  }\\n}\\nfragment ItemView on Item {\\n  chain {\\n    identifier\\n    arch\\n    __typename\\n  }\\n  contractAddress\\n  isFungible\\n  tokenId\\n  enforcement {\\n    isDisabled\\n    __typename\\n  }\\n  version\\n  ...ItemAbout\\n  ...ItemStats\\n  ...ItemDetails\\n  ...ItemSocial\\n  ...ItemAction\\n  ...itemUrl\\n  ...ItemTabs\\n  ...ItemOrders\\n  ...ItemMetadataChips\\n  ...ItemTitle\\n  ...ItemPageMedia\\n  ...itemIdentifier\\n  __typename\\n}\\nfragment ItemStats on Item {\\n  tokenId\\n  isFungible\\n  totalSupply\\n  rarity {\\n    rank\\n    __typename\\n  }\\n  bestOffer {\\n    pricePerItem {\\n      usd\\n      ...TokenPrice\\n      __typename\\n    }\\n    __typename\\n  }\\n  bestListing {\\n    pricePerItem {\\n      usd\\n      ...TokenPrice\\n      __typename\\n    }\\n    __typename\\n  }\\n  collection {\\n    id\\n    slug\\n    floorPrice {\\n      pricePerItem {\\n        usd\\n        ...TokenPrice\\n        __typename\\n      }\\n      __typename\\n    }\\n    __typename\\n  }\\n  lastSale {\\n    ...TokenPrice\\n    __typename\\n  }\\n  ...isItemRarityDisabled\\n  ...RarityTooltip\\n  __typename\\n}\\nfragment TokenPrice on Price {\\n  usd\\n  token {\\n    unit\\n    symbol\\n    contractAddress\\n    chain {\\n      identifier\\n      __typename\\n    }\\n    __typename\\n  }\\n  __typename\\n}\\nfragment isItemRarityDisabled on Item {\\n  collection {\\n    id\\n    slug\\n    __typename\\n  }\\n  __typename\\n}\\nfragment RarityTooltip on Item {\\n  rarity {\\n    category\\n    rank\\n    totalSupply\\n    __typename\\n  }\\n  ...isItemRarityDisabled\\n  __typename\\n}\\nfragment ItemAbout on Item {\\n  id\\n  tokenId\\n  tokenUri\\n  contractAddress\\n  chain {\\n    name\\n    identifier\\n    arch\\n    __typename\\n  }\\n  standard\\n  description\\n  collection {\\n    description\\n    __typename\\n  }\\n  details {\\n    name\\n    value\\n    __typename\\n  }\\n  collection {\\n    ...CollectionOwner\\n    owner {\\n      displayName\\n      __typename\\n    }\\n    __typename\\n  }\\n  __typename\\n}\\nfragment CollectionOwner on Collection {\\n  owner {\\n    displayName\\n    isVerified\\n    address\\n    ...profileUrl\\n    ...ProfilePreviewTooltip\\n    __typename\\n  }\\n  standard\\n  __typename\\n}\\nfragment profileUrl on ProfileIdentifier {\\n  address\\n  __typename\\n}\\nfragment ProfilePreviewTooltip on ProfileIdentifier {\\n  address\\n  ...ProfilePreviewTooltipContent\\n  __typename\\n}\\nfragment ProfilePreviewTooltipContent on ProfileIdentifier {\\n  address\\n  __typename\\n}\\nfragment ItemDetails on Item {\\n  name\\n  contractAddress\\n  tokenId\\n  chain {\\n    identifier\\n    __typename\\n  }\\n  owner {\\n    address\\n    ...AccountLockup\\n    ...profileUrl\\n    ...ProfilePreviewTooltip\\n    __typename\\n  }\\n  isFungible\\n  ...ItemCollection\\n  ...ItemSocial\\n  ...itemUrl\\n  collection {\\n    isVerified\\n    __typename\\n  }\\n  ...FavoriteItemButton\\n  ...ItemDetailsDropdown\\n  __typename\\n}\\nfragment ItemCollection on Item {\\n  collection {\\n    slug\\n    imageUrl\\n    isVerified\\n    name\\n    ...collectionUrl\\n    __typename\\n  }\\n  __typename\\n}\\nfragment collectionUrl on CollectionIdentifier {\\n  slug\\n  __typename\\n}\\nfragment ItemSocial on Item {\\n  chain {\\n    identifier\\n    __typename\\n  }\\n  contractAddress\\n  tokenId\\n  externalUrl\\n  collection {\\n    externalUrl\\n    discordUrl\\n    instagramUsername\\n    twitterUsername\\n    __typename\\n  }\\n  __typename\\n}\\nfragment AccountLockup on ProfileIdentifier {\\n  address\\n  displayName\\n  imageUrl\\n  ...profileUrl\\n  __typename\\n}\\nfragment itemUrl on ItemIdentifier {\\n  chain {\\n    identifier\\n    arch\\n    __typename\\n  }\\n  tokenId\\n  contractAddress\\n  __typename\\n}\\nfragment FavoriteItemButton on Item {\\n  ...itemIdentifier\\n  __typename\\n}\\nfragment itemIdentifier on ItemIdentifier {\\n  chain {\\n    identifier\\n    __typename\\n  }\\n  tokenId\\n  contractAddress\\n  __typename\\n}\\nfragment ItemDetailsDropdown on Item {\\n  ...ItemDetailsDropdownRefreshMetadata\\n  ...ItemDetailsDropdownViewOriginalMedia\\n  ...itemIdentifier\\n  __typename\\n}\\nfragment ItemDetailsDropdownRefreshMetadata on Item {\\n  ...useItemRefreshMetadata\\n  __typename\\n}\\nfragment useItemRefreshMetadata on ItemIdentifier {\\n  ...itemIdentifier\\n  __typename\\n}\\nfragment ItemDetailsDropdownViewOriginalMedia on Item {\\n  originalImageUrl\\n  originalAnimationUrl\\n  __typename\\n}\\nfragment ItemAction on Item {\\n  id\\n  isFungible\\n  version\\n  ...itemIdentifier\\n  __typename\\n}\\nfragment ItemTabs on Item {\\n  isFungible\\n  collection {\\n    slug\\n    __typename\\n  }\\n  __typename\\n}\\nfragment ItemOrders on Item {\\n  contractAddress\\n  tokenId\\n  collection {\\n    slug\\n    __typename\\n  }\\n  chain {\\n    identifier\\n    __typename\\n  }\\n  isFungible\\n  ...ItemOrdersDepthChart\\n  ...ItemOrdersFeed\\n  __typename\\n}\\nfragment ItemOrdersDepthChart on Item {\\n  contractAddress\\n  tokenId\\n  chain {\\n    identifier\\n    __typename\\n  }\\n  __typename\\n}\\nfragment ItemOrdersFeed on Item {\\n  tokenId\\n  collection {\\n    slug\\n    __typename\\n  }\\n  ...ItemOffersTable\\n  __typename\\n}\\nfragment ItemOffersTable on Item {\\n  isFungible\\n  ...itemIdentifier\\n  ...useAcceptOffers\\n  __typename\\n}\\nfragment useAcceptOffers on Item {\\n  chain {\\n    identifier\\n    arch\\n    __typename\\n  }\\n  contractAddress\\n  tokenId\\n  collection {\\n    isTradingDisabled\\n    __typename\\n  }\\n  bestOffer {\\n    pricePerItem {\\n      token {\\n        unit\\n        address\\n        __typename\\n      }\\n      __typename\\n    }\\n    maker {\\n      address\\n      __typename\\n    }\\n    __typename\\n  }\\n  enforcement {\\n    isCompromised\\n    __typename\\n  }\\n  __typename\\n}\\nfragment ItemMetadataChips on Item {\\n  ...ItemMetadataChip\\n  __typename\\n}\\nfragment ItemMetadataChip on Item {\\n  ...ItemChainChip\\n  ...ItemRarityChip\\n  ...ItemTokenIdChip\\n  ...ItemStandardChip\\n  ...ItemTopOfferChip\\n  ...ItemOwnersChip\\n  __typename\\n}\\nfragment ItemChainChip on Item {\\n  chain {\\n    ...ChainChip\\n    __typename\\n  }\\n  __typename\\n}\\nfragment ChainChip on Chain {\\n  identifier\\n  name\\n  __typename\\n}\\nfragment ItemRarityChip on Item {\\n  rarity {\\n    rank\\n    category\\n    __typename\\n  }\\n  __typename\\n}\\nfragment ItemTokenIdChip on Item {\\n  tokenId\\n  __typename\\n}\\nfragment ItemStandardChip on Item {\\n  standard\\n  __typename\\n}\\nfragment ItemTopOfferChip on Item {\\n  bestOffer {\\n    pricePerItem {\\n      ...TokenPrice\\n      __typename\\n    }\\n    __typename\\n  }\\n  enforcement {\\n    isCompromised\\n    __typename\\n  }\\n  ...EnforcementBadge\\n  __typename\\n}\\nfragment EnforcementBadge on EnforcedEntity {\\n  __typename\\n  enforcement {\\n    isCompromised\\n    isDisabled\\n    isOwnershipDisputed\\n    __typename\\n  }\\n}\\nfragment ItemOwnersChip on Item {\\n  ...ItemOwnersModal\\n  ...ItemOwnersCount\\n  isFungible\\n  __typename\\n}\\nfragment ItemOwnersModal on Item {\\n  ...ItemOwnersModalContent\\n  __typename\\n}\\nfragment ItemOwnersModalContent on Item {\\n  ...ItemOwnersTable\\n  __typename\\n}\\nfragment ItemOwnersTable on Item {\\n  tokenId\\n  chain {\\n    identifier\\n    __typename\\n  }\\n  contractAddress\\n  __typename\\n}\\nfragment ItemOwnersCount on Item {\\n  tokenId\\n  chain {\\n    identifier\\n    __typename\\n  }\\n  contractAddress\\n  __typename\\n}\\nfragment ItemTitle on Item {\\n  name\\n  isFavorite\\n  ...EnforcementBadge\\n  __typename\\n}\\nfragment ItemPageMedia on Item {\\n  ...ItemMedia\\n  __typename\\n}\\nfragment ItemMedia on Item {\\n  imageUrl\\n  animationUrl\\n  backgroundColor\\n  collection {\\n    imageUrl\\n    __typename\\n  }\\n  __typename\\n}\",\n  \"variables\": {\n    \"identifier\": {\n      \"chain\": \"ethereum\",\n      \"contractAddress\": \"0xed5af388653567af2f388e6224dc7c4b3241c544\",\n      \"tokenId\": \"" + strconv.Itoa(tokenId) + "\"\n    }\n  }\n}"
		req, err := http.NewRequest("POST", "https://gql.opensea.io/graphql", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Origin", "https://opensea.io")
		req.Header.Set("Referer", "https://opensea.io/")
		// 有时需要加上下面这一行才能让 Cloudflare 放行
		req.Header.Set("x-requested-with", "XMLHttpRequest")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("err:", err)
			panic(err)
		}

		var respMap map[string]any

		bodyBytes, err := io.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, &respMap)
		//获取tokenUri
		// 假设 respMap 是 map[string]interface{}
		tokenUri, _ := getString(respMap, "data", "itemByIdentifier", "tokenUri")
		fmt.Println(tokenUri)

		//获取image url
		uriResp, _ := http.Get(tokenUri)
		bodyBytes, _ = io.ReadAll(uriResp.Body)
		json.Unmarshal(bodyBytes, &respMap)
		uriImage, _ := getString(respMap, "image")
		fmt.Println(uriImage)
		withoutPrefix := strings.TrimPrefix(uriImage, "ipfs://")
		fmt.Println(withoutPrefix)
		pictureURI := "https://ipfs.io/ipfs/" + withoutPrefix
		err = downloadFile(pictureURI, "D:\\files\\web3\\lumao\\pics\\pic"+strconv.Itoa(i)+".png")
		if err != nil {
			fmt.Println("err:", err)
		}
		tokenId++
		time.Sleep(2 * time.Second)
	}

}

func getString(m map[string]interface{}, keys ...string) (string, bool) {
	var cur interface{} = m
	for _, k := range keys {
		m2, ok := cur.(map[string]interface{})
		if !ok {
			return "", false
		}
		cur, ok = m2[k]
		if !ok {
			return "", false
		}
	}
	s, ok := cur.(string)
	return s, ok
}

// downloadFile 从给定 URL 下载内容，并保存到本地路径 filepath
func downloadFile(url string, filepath string) error {
	// 1. 发起 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http.Get 错误: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("服务器返回状态: %s", resp.Status)
	}

	// 2. 在本地创建文件
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("os.Create 错误: %w", err)
	}
	defer out.Close()

	// 3. 将响应 body 写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("io.Copy 错误: %w", err)
	}

	return nil
}
