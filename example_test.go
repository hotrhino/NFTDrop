package NFTDrop

import (
	"NFTDrop/internal/NFTDrop"
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"net/http"
	"testing"
)

func TestClaim(t *testing.T) {
	//1. init params (replace these values with real values)
	contractsAddress := "0x00D0A67D6E882fe34Eb33566Fb36d25d2C157f01"
	endPoint := "https://opbnb-testnet-rpc.bnbchain.org/"
	priKeyHex := "18...86" // private key in iteration
	merkleUrl := "https://x1-t01-api.legendofarcadia.io/PreSale/Proof?sale=44"

	limit := big.NewInt(1)
	price := big.NewInt(0)
	currency := common.HexToAddress("0x0000000000000000000000000000000000000000")

	//2. prepare user address
	priKey, err := crypto.HexToECDSA(priKeyHex)
	if err != nil {
		t.Fatal(err)
	}
	userAddress := crypto.PubkeyToAddress(priKey.PublicKey)

	//3. use http get to get merkle proof
	req, err := http.NewRequest("GET", merkleUrl, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("User-Address", userAddress.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("http get failed")
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var data MerkleProofResp
	err = json.Unmarshal(respData, &data)
	if err != nil {
		t.Fatal(err)
	}
	if data.Code != 0 {
		t.Fatal(data.Code)
	}

	// init cli
	rawRpc, err := ethclient.Dial(endPoint)
	if err != nil {
		t.Fatal(err)
	}
	cli, err := NFTDrop.NewNFTDrop(common.HexToAddress(contractsAddress), rawRpc)
	if err != nil {
		t.Fatal(err)
	}
	chainId, err := rawRpc.ChainID(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	//4. claim
	opt, err := bind.NewKeyedTransactorWithChainID(priKey, chainId)
	if err != nil {
		t.Fatal(err)
	}

	proof := &NFTDrop.IDropAllowlistProof{
		Proof:                  nil,
		QuantityLimitPerWallet: limit,
		PricePerToken:          price,
		Currency:               currency,
	}
	for _, path := range data.Data.Paths {
		proof.Proof = append(proof.Proof, common.HexToHash(path))
	}

	tx, err := cli.Claim(opt, userAddress, limit, currency, price, *proof, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx.Hash(), tx.Gas(), tx.GasPrice().String())
}
