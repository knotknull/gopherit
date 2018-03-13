package main

import "github.com/hashicorp/memberlist"

type broadcast struct {
	msg    []byte
	notify chan struct{}
}

func (b *broadcast) Invalidates(other memberlist.Broadcast) bool {
	return false
}

func (b *broadcast) Message() []byte {
	return b.msg
}

// Finished will close the channel (if provided) to allow for asyc signal communication
func (b *broadcast) Finished() {
	if b.notify != nil {
		close(b.notify)
	}
}
