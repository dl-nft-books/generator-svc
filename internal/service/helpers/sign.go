package helpers

import (
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	signer "github.com/ethersphere/bee/pkg/crypto"
	"github.com/ethersphere/bee/pkg/crypto/eip712"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
)

const defaultAddress = "0x0000000000000000000000000000000000000000"

type SignatureParameters struct {
	R string `json:"r"`
	S string `json:"s"`
	V int    `json:"v"`
}

type SignInfo struct {
	ContractName    string
	ContractVersion string
	ContractAddress string
	TokenAddress    string
	Price           *big.Int
	ChainID         int64
}

func Sign(info *SignInfo, config *config.EthMinterConfig) (*SignatureParameters, error) {
	privateKey := config.PrivateKey

	endTimestamp := time.Now().Add(config.Expiration).Unix()

	if info.TokenAddress == "" {
		info.TokenAddress = defaultAddress
	}

	signature, err := signEIP712(privateKey, info, endTimestamp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign EIP712 hash")
	}

	return parseSignatureParameters(signature)
}

func signEIP712(privateKey *ecdsa.PrivateKey, info *SignInfo, endTimestamp int64) ([]byte, error) {
	data := &eip712.TypedData{
		Types: apitypes.Types{
			"Mint": []apitypes.Type{
				{Name: "paymentTokenAddress", Type: "address"},
				{Name: "paymentTokenPrice", Type: "uint256"},
				{Name: "endTimestamp", Type: "uint256"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "Mint",
		Domain: apitypes.TypedDataDomain{
			Name:              info.ContractName,
			Version:           info.ContractVersion,
			ChainId:           math.NewHexOrDecimal256(info.ChainID),
			VerifyingContract: info.ContractAddress,
		},
		Message: apitypes.TypedDataMessage{
			"paymentTokenAddress": info.TokenAddress,
			"paymentTokenPrice":   info.Price.String(),
			"endTimestamp":        math.NewHexOrDecimal256(endTimestamp),
		},
	}

	return signer.NewDefaultSigner(privateKey).SignTypedData(data)
}

func parseSignatureParameters(signature []byte) (*SignatureParameters, error) {
	if len(signature) != 65 {
		return nil, errors.New("bad signature")
	}

	params := SignatureParameters{}

	params.R = hexutil.Encode(signature[:32])
	params.S = hexutil.Encode(signature[32:64])
	params.V = int(signature[64])

	return &params, nil
}
