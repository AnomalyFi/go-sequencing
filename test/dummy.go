package test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"sync"

	"github.com/rollkit/go-sequencing"
)

var ErrorRollupIdMismatch = errors.New("rollup id mismatch")

type TransactionQueue struct {
	queue []sequencing.Tx
	mu    sync.Mutex
}

// NewTransactionQueue creates a new TransactionQueue
func NewTransactionQueue() *TransactionQueue {
	return &TransactionQueue{
		queue: make([]sequencing.Tx, 0),
	}
}

// AddTransaction adds a new transaction to the queue
func (tq *TransactionQueue) AddTransaction(tx sequencing.Tx) {
	tq.mu.Lock()
	defer tq.mu.Unlock()
	tq.queue = append(tq.queue, tx)
}

// GetBatch extracts a batch of transactions from the queue
func (tq *TransactionQueue) GetNextBatch() []sequencing.Tx {
	tq.mu.Lock()
	defer tq.mu.Unlock()

	batchSize := len(tq.queue)

	batch := tq.queue[:batchSize]
	tq.queue = tq.queue[batchSize:]
	return batch
}

// DummySequencer is a dummy sequencer for testing that serves a single rollup
type DummySequencer struct {
	sequencing.RollupId

	tq            *TransactionQueue
	lastBatchHash []byte

	seenBatches map[string]struct{}
}

// SubmitRollupTransaction implements sequencing.Sequencer.
func (d *DummySequencer) SubmitRollupTransaction(ctx context.Context, rollupId []byte, tx []byte) error {
	if d.RollupId == nil {
		d.RollupId = rollupId
	} else {
		if !bytes.Equal(d.RollupId, rollupId) {
			return ErrorRollupIdMismatch
		}
	}
	d.tq.AddTransaction(tx)
	return nil
}

// GetNextBatch implements sequencing.Sequencer.
func (d *DummySequencer) GetNextBatch(ctx context.Context, lastBatch [][]byte) ([][]byte, error) {
	batch := d.tq.GetNextBatch()
	d.lastBatchHash = hashSHA256(batch[0])
	d.seenBatches[string(d.lastBatchHash)] = struct{}{}
	return batch, nil
}

// VerifyBatch implements sequencing.Sequencer.
func (d *DummySequencer) VerifyBatch(ctx context.Context, batch [][]byte) (bool, error) {
	hash := hashSHA256(batch[0])
	_, ok := d.seenBatches[string(hash)]
	return ok, nil
}

func NewDummySequencer() *DummySequencer {
	return &DummySequencer{
		tq: NewTransactionQueue(),
	}
}

func hashSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

var _ sequencing.Sequencer = &DummySequencer{}
