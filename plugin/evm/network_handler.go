// (c) 2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"

	"github.com/luxdefi/node/codec"
	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/coreth/ethdb"
	"github.com/luxdefi/coreth/metrics"
	"github.com/luxdefi/coreth/plugin/evm/message"
	syncHandlers "github.com/luxdefi/coreth/sync/handlers"
	syncStats "github.com/luxdefi/coreth/sync/handlers/stats"
	"github.com/luxdefi/coreth/trie"
	"github.com/luxdefi/coreth/warp"
	warpHandlers "github.com/luxdefi/coreth/warp/handlers"
)

var _ message.RequestHandler = &networkHandler{}

type networkHandler struct {
	stateTrieLeafsRequestHandler  *syncHandlers.LeafsRequestHandler
	atomicTrieLeafsRequestHandler *syncHandlers.LeafsRequestHandler
	blockRequestHandler           *syncHandlers.BlockRequestHandler
	codeRequestHandler            *syncHandlers.CodeRequestHandler
	signatureRequestHandler       *warpHandlers.SignatureRequestHandler
}

// newNetworkHandler constructs the handler for serving network requests.
func newNetworkHandler(
	provider syncHandlers.SyncDataProvider,
	diskDB ethdb.KeyValueReader,
	evmTrieDB *trie.Database,
	atomicTrieDB *trie.Database,
	warpBackend warp.Backend,
	networkCodec codec.Manager,
) message.RequestHandler {
	syncStats := syncStats.NewHandlerStats(metrics.Enabled)
	return &networkHandler{
		stateTrieLeafsRequestHandler:  syncHandlers.NewLeafsRequestHandler(evmTrieDB, provider, networkCodec, syncStats),
		atomicTrieLeafsRequestHandler: syncHandlers.NewLeafsRequestHandler(atomicTrieDB, nil, networkCodec, syncStats),
		blockRequestHandler:           syncHandlers.NewBlockRequestHandler(provider, networkCodec, syncStats),
		codeRequestHandler:            syncHandlers.NewCodeRequestHandler(diskDB, networkCodec, syncStats),
		signatureRequestHandler:       warpHandlers.NewSignatureRequestHandler(warpBackend, networkCodec),
	}
}

func (n networkHandler) HandleStateTrieLeafsRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, leafsRequest message.LeafsRequest) ([]byte, error) {
	return n.stateTrieLeafsRequestHandler.OnLeafsRequest(ctx, nodeID, requestID, leafsRequest)
}

func (n networkHandler) HandleAtomicTrieLeafsRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, leafsRequest message.LeafsRequest) ([]byte, error) {
	return n.atomicTrieLeafsRequestHandler.OnLeafsRequest(ctx, nodeID, requestID, leafsRequest)
}

func (n networkHandler) HandleBlockRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, blockRequest message.BlockRequest) ([]byte, error) {
	return n.blockRequestHandler.OnBlockRequest(ctx, nodeID, requestID, blockRequest)
}

func (n networkHandler) HandleCodeRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, codeRequest message.CodeRequest) ([]byte, error) {
	return n.codeRequestHandler.OnCodeRequest(ctx, nodeID, requestID, codeRequest)
}

func (n networkHandler) HandleMessageSignatureRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, messageSignatureRequest message.MessageSignatureRequest) ([]byte, error) {
	return n.signatureRequestHandler.OnMessageSignatureRequest(ctx, nodeID, requestID, messageSignatureRequest)
}

func (n networkHandler) HandleBlockSignatureRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, blockSignatureRequest message.BlockSignatureRequest) ([]byte, error) {
	return n.signatureRequestHandler.OnBlockSignatureRequest(ctx, nodeID, requestID, blockSignatureRequest)
}
