// Copyright (c) 2014-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcrpcclient

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/ppcsuite/ppcd/btcjson"
	"github.com/ppcsuite/ppcd/wire"
)

// FutureKernelStakeModifierResult is a future promise to deliver the result of a
// GetKernelStakeModifierAsync RPC invocation (or an applicable error).
type FutureKernelStakeModifierResult chan *response

// Receive waits for the response promised by the future and returns the raw
// block requested from the server given its hash.
func (r FutureKernelStakeModifierResult) Receive() (uint64, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal the result as a string.
	var kernelStakeModifierStr string
	err = json.Unmarshal(res, &kernelStakeModifierStr)
	if err != nil {
		return 0, err
	}
	var kernelStakeModifier uint64
	kernelStakeModifier, err = strconv.ParseUint(kernelStakeModifierStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return kernelStakeModifier, nil
}

// GetKernelStakeModifierAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetKernelStakeModifier for the blocking version and more details.
func (c *Client) GetKernelStakeModifierAsync(blockHash *wire.ShaHash) FutureKernelStakeModifierResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	verbose := false
	cmd := btcjson.NewGetKernelStakeModifierCmd(hash, &verbose)

	return c.sendCmd(cmd)
}

// GetKernelStakeModifier returns a raw block from the server given its hash.
//
// See GetKernelStakeModifierVerbose to retrieve a data structure with information about the
// block instead.
func (c *Client) GetKernelStakeModifier(blockHash *wire.ShaHash) (uint64, error) {
	return c.GetKernelStakeModifierAsync(blockHash).Receive()
}

// FutureGetKernelStakeModifierVerboseResult is a future promise to deliver the result of a
// GetKernelStakeModifierVerboseAsync RPC invocation (or an applicable error).
type FutureGetKernelStakeModifierVerboseResult chan *response

// Receive waits for the response promised by the future and returns the data
// structure from the server with information about the requested block.
func (r FutureGetKernelStakeModifierVerboseResult) Receive() (*btcjson.KernelStakeModifierResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the raw result into a KernelStakeModifierResult.
	var ksmResult btcjson.KernelStakeModifierResult
	err = json.Unmarshal(res, &ksmResult)
	if err != nil {
		return nil, err
	}
	return &ksmResult, nil
}

// GetKernelStakeModifierVerboseAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See GetKernelStakeModifierVerbose for the blocking version and more details.
func (c *Client) GetKernelStakeModifierVerboseAsync(blockHash *wire.ShaHash) FutureGetKernelStakeModifierVerboseResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	verbose := true
	cmd := btcjson.NewGetKernelStakeModifierCmd(hash, &verbose)

	return c.sendCmd(cmd)
}

// GetKernelStakeModifierVerbose returns a data structure from the server with information
// about a block given its hash.
//
// See GetKernelStakeModifier to retrieve a raw block instead.
func (c *Client) GetKernelStakeModifierVerbose(blockHash *wire.ShaHash) (*btcjson.KernelStakeModifierResult, error) {
	return c.GetKernelStakeModifierVerboseAsync(blockHash).Receive()
}

// FutureNextRequiredTargetResult is a future promise to deliver the result of a
// GetNextRequiredTargetAsync RPC invocation (or an applicable error).
type FutureNextRequiredTargetResult chan *response

// Receive waits for the response promised by the future and returns the raw
// block requested from the server given its hash.
func (r FutureNextRequiredTargetResult) Receive() (uint32, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal the result as a string.
	var nextRequiredTargetStr string
	err = json.Unmarshal(res, &nextRequiredTargetStr)
	if err != nil {
		return 0, err
	}
	var nextRequiredTarget uint64
	nextRequiredTarget, err = strconv.ParseUint(nextRequiredTargetStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(nextRequiredTarget), nil
}

// GetNextRequiredTargetAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetNextRequiredTarget for the blocking version and more details.
func (c *Client) GetNextRequiredTargetAsync(proofOfStake bool) FutureNextRequiredTargetResult {

	verbose := false
	cmd := btcjson.NewGetNextRequiredTargetCmd(&proofOfStake, &verbose)

	return c.sendCmd(cmd)
}

// GetNextRequiredTarget returns a raw block from the server given its hash.
//
// See GetNextRequiredTargetVerbose to retrieve a data structure with information about the
// block instead.
func (c *Client) GetNextRequiredTarget(proofOfStake bool) (uint32, error) {
	return c.GetNextRequiredTargetAsync(proofOfStake).Receive()
}

// FutureGetNextRequiredTargetVerboseResult is a future promise to deliver the result of a
// GetNextRequiredTargetVerboseAsync RPC invocation (or an applicable error).
type FutureGetNextRequiredTargetVerboseResult chan *response

// Receive waits for the response promised by the future and returns the data
// structure from the server with information about the requested block.
func (r FutureGetNextRequiredTargetVerboseResult) Receive() (*btcjson.NextRequiredTargetResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the raw result into a NextRequiredTargetResult.
	var nrtResult btcjson.NextRequiredTargetResult
	err = json.Unmarshal(res, &nrtResult)
	if err != nil {
		return nil, err
	}
	return &nrtResult, nil
}

// GetNextRequiredTargetVerboseAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See GetNextRequiredTargetVerbose for the blocking version and more details.
func (c *Client) GetNextRequiredTargetVerboseAsync(proofOfStake bool) FutureGetNextRequiredTargetVerboseResult {

	verbose := true
	cmd := btcjson.NewGetNextRequiredTargetCmd(&proofOfStake, &verbose)

	return c.sendCmd(cmd)
}

// GetNextRequiredTargetVerbose returns... TODO(mably)
//
// See GetNextRequiredTarget to retrieve a raw block instead.
func (c *Client) GetNextRequiredTargetVerbose(proofOfStake bool) (*btcjson.NextRequiredTargetResult, error) {
	return c.GetNextRequiredTargetVerboseAsync(proofOfStake).Receive()
}

// FutureLastProofOfWorkRewardResult is a future promise to deliver the result of a
// GetNextRequiredTargetAsync RPC invocation (or an applicable error).
type FutureLastProofOfWorkRewardResult chan *response

// Receive waits for the response promised by the future and returns the raw
// block requested from the server given its hash.
func (r FutureLastProofOfWorkRewardResult) Receive() (int64, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal the raw result into a LastProofOfWorkRewardResult.
	var lpowrResult btcjson.LastProofOfWorkRewardResult
	err = json.Unmarshal(res, &lpowrResult)
	if err != nil {
		return 0, err
	}

	return int64(lpowrResult.Subsidy), nil
}

// GetLastProofOfWorkRewardAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetLastProofOfWorkReward for the blocking version and more details.
func (c *Client) GetLastProofOfWorkRewardAsync() FutureLastProofOfWorkRewardResult {

	cmd := btcjson.NewGetLastProofOfWorkRewardCmd()

	return c.sendCmd(cmd)
}

// GetLastProofOfWorkReward returns a raw block from the server given its hash.
func (c *Client) GetLastProofOfWorkReward() (int64, error) {
	return c.GetLastProofOfWorkRewardAsync().Receive()
}

// FutureSendCoinStakeTransactionResult is a future promise to deliver the result
// of a SendCoinStakeTransactionAsync RPC invocation (or an applicable error).
type FutureSendCoinStakeTransactionResult chan *response

// Receive waits for the response promised by the future and returns the result
// of submitting the encoded transaction to the server which then relays it to
// the network.
func (r FutureSendCoinStakeTransactionResult) Receive() (*wire.ShaHash, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var txHashStr string
	err = json.Unmarshal(res, &txHashStr)
	if err != nil {
		return nil, err
	}

	return wire.NewShaHashFromStr(txHashStr)
}

// SendCoinStakeTransactionAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See SendCoinStakeTransaction for the blocking version and more details.
func (c *Client) SendCoinStakeTransactionAsync(tx *wire.MsgTx) FutureSendCoinStakeTransactionResult {
	txHex := ""
	if tx != nil {
		// Serialize the transaction and convert to hex string.
		buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
		if err := tx.Serialize(buf); err != nil {
			return newFutureError(err)
		}
		txHex = hex.EncodeToString(buf.Bytes())
	}

	cmd := btcjson.NewSendCoinStakeTransactionCmd(txHex)
	return c.sendCmd(cmd)
}

// SendCoinStakeTransaction submits the encoded transaction to the server which will
// then relay it to the network.
func (c *Client) SendCoinStakeTransaction(tx *wire.MsgTx) (*wire.ShaHash, error) {
	return c.SendCoinStakeTransactionAsync(tx).Receive()
}
