// Copyright (c) 2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcrpcclient

import (
	"encoding/json"

	"github.com/mably/btcjson"
	"github.com/mably/btcwire"
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

	// Unmarshal the result as an uint64.
	var kernelStakeModifier uint64
	err = json.Unmarshal(res, &kernelStakeModifier)
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
func (c *Client) GetKernelStakeModifierAsync(blockHash *btcwire.ShaHash) FutureKernelStakeModifierResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	id := c.NextID()
	cmd, err := btcjson.NewGetKernelStakeModifierCmd(id, hash, false)
	if err != nil {
		return newFutureError(err)
	}

	return c.sendCmd(cmd)
}

// GetKernelStakeModifier returns a raw block from the server given its hash.
//
// See GetKernelStakeModifierVerbose to retrieve a data structure with information about the
// block instead.
func (c *Client) GetKernelStakeModifier(blockHash *btcwire.ShaHash) (uint64, error) {
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
func (c *Client) GetKernelStakeModifierVerboseAsync(blockHash *btcwire.ShaHash) FutureGetKernelStakeModifierVerboseResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	id := c.NextID()
	cmd, err := btcjson.NewGetKernelStakeModifierCmd(id, hash, true)
	if err != nil {
		return newFutureError(err)
	}

	return c.sendCmd(cmd)
}

// GetKernelStakeModifierVerbose returns a data structure from the server with information
// about a block given its hash.
//
// See GetKernelStakeModifier to retrieve a raw block instead.
func (c *Client) GetKernelStakeModifierVerbose(blockHash *btcwire.ShaHash) (*btcjson.KernelStakeModifierResult, error) {
	return c.GetKernelStakeModifierVerboseAsync(blockHash).Receive()
}
