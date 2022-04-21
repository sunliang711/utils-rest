package utils

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func Verify(msg_ []byte, sig_ string, publicKey_ string) (bool, error) {
	publicKey, err := hexutil.Decode("0x" + publicKey_)
	if err != nil {
		return false, err
	}
	publicKey__, err := crypto.UnmarshalPubkey(publicKey)
	if err != nil {
		return false, err
	}
	address := crypto.PubkeyToAddress(*publicKey__)

	return VerifyWithAddress(msg_, sig_, hexutil.Encode(address.Bytes()))
}

// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L404
// signHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calculated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

// func VerifySig(msg []byte, from, sigHex string) (bool, error) {
// 	fromAddr := common.HexToAddress(from)

// 	sig, err := hexutil.Decode(sigHex)
// 	if err != nil {
// 		return false, err
// 	}
// 	logrus.Debugf("sig: %v", sig)
// 	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
// 	if sig[63] != 27 || sig[63] != 28 {
// 		return false, nil
// 	}
// 	sig[63] -= 27

// 	pubKey, err := crypto.SigToPub(signHash(msg), sig)
// 	if err != nil {
// 		return false, err
// 	}

// 	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

// 	return fromAddr == recoveredAddr, nil
// }

func VerifyWithAddress(msg_ []byte, sig_ string, address_ string) (bool, error) {
	signature, err := hexutil.Decode(sig_)
	if err != nil {
		return false, err
	}
	pk, err := crypto.Ecrecover(signHash(msg_), signature)
	if err != nil {
		return false, err
	}

	publicKey, err := crypto.UnmarshalPubkey(pk)
	if err != nil {
		return false, err
	}
	address := crypto.PubkeyToAddress(*publicKey)
	address__, err := hexutil.Decode(address_)
	if err != nil {
		return false, err
	}
	return bytes.Equal(address.Bytes(), address__), nil

}

func VerifyWithAddress2(msg_ []byte, sig_ string, address_ string) (bool, error) {
	hash := signHash(msg_)
	signature, err := hexutil.Decode(sig_)
	if err != nil {
		return false, err
	}
	if len(signature) != 65 {
		return false, errors.New("invalid signature length")
	}
	signature[64] -= 27
	sigPublicKeyECDSA, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return false, err
	}

	sigAddress := crypto.PubkeyToAddress(*sigPublicKeyECDSA)

	address, err := hexutil.Decode(address_)
	if err != nil {
		return false, err
	}
	return bytes.Equal(sigAddress.Bytes(), address), nil
}
