package store

import (
	"encoding/json"

	"github.com/hashicorp/raft"
)

type fsmSnapshot struct {
	store map[string]string
}

// Persist writes the internal snapshot to the raft snapshot sink
func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	err := func() error {
		// Encode data.
		b, err := json.Marshal(f.store)
		if err != nil {
			return err
		}

		// Write data to sink.
		if _, err := sink.Write(b); err != nil {
			return err
		}

		// Close the sink.
		if err := sink.Close(); err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		sink.Cancel()
		return err
	}

	return nil
}

//  Release is a noop method in our implementation because we cloned the
//  store before we sent it in.  We could of blocked the store from
//  writes and then released it via this method instead.
func (f *fsmSnapshot) Release() {}
