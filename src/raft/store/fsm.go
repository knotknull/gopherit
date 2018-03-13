package store

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/hashicorp/raft"
)

// fsm represents the underlying finite state machine
type fsm struct {
	mu   sync.RWMutex
	data map[string]string // The underlying data for the store
}

func NewFSM() *fsm {
	return &fsm{
		data: make(map[string]string),
	}
}

// Apply applies a Raft log entry to the key-value store.
func (f *fsm) Apply(l *raft.Log) interface{} {
	var c command

	if err := json.Unmarshal(l.Data, &c); err != nil {
		// Be sure to check your response to Apply.  Both Error() and Response()
		return fmt.Errorf("failed to unmarshal command: %s", err)
	}

	switch c.Command {
	case "set":
		f.applySet(c.Key, c.Value)
		return nil
	case "delete":
		f.applyDelete(c.Key)
		return nil
	default:
		return fmt.Errorf("invalid command: %s", c.Command)
	}
}

// Snapshot returns a snapshot of the key-value store.
func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	// We clone to avoid locking the store on a snapshot write.
	o := make(map[string]string)
	for k, v := range f.data {
		o[k] = v
	}
	return &fsmSnapshot{store: o}, nil
}

// Restore stores the key-value store to a previous state.
func (f *fsm) Restore(rc io.ReadCloser) error {
	o := make(map[string]string)
	if err := json.NewDecoder(rc).Decode(&o); err != nil {
		return err
	}

	// Set the state from the snapshot, no lock required according to
	// Hashicorp docs.
	f.data = o
	return nil
}

// Get will return a value from the state machine
func (f *fsm) Get(key string) string {
	f.mu.RLock()
	defer f.mu.RUnlock()

	return f.data[key]
}

func (f *fsm) applySet(key, value string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.data[key] = value
}

func (f *fsm) applyDelete(key string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	delete(f.data, key)
}
