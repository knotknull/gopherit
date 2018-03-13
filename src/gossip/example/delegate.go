package main

import "encoding/json"

type delegate struct {
	cache *Cache
}

func NewDelegate(c *Cache) *delegate {
	return &delegate{
		cache: c,
	}
}

type update struct {
	Action string // add, del
	Data   map[string]string
}

// NodeMeta is used to retrieve meta-data about the current node
// when broadcasting an alive message. It's length is limited to
// the given byte size. This metadata is available in the Node structure.
func (d *delegate) NodeMeta(limit int) []byte {
	// not currently implemented
	return []byte{}
}

// GetBroadcasts is called when user data messages can be broadcast.
// It can return a list of buffers to send. Each buffer should assume an
// overhead as provided with a limit on the total byte size allowed.
// The total byte size of the resulting data to send must not exceed
// the limit. Care should be taken that this method does not block,
// since doing so would block the entire UDP packet receive loop.
func (d *delegate) GetBroadcasts(overhead, limit int) [][]byte {
	// using the implementation provided from `TransmitLimitedQueue`
	return broadcasts.GetBroadcasts(overhead, limit)
}

// NotifyMsg is called when a user-data message is received.
// Care should be taken that this method does not block, since doing
// so would block the entire UDP packet receive loop. Additionally, the byte
// slice may be modified after the call returns, so it should be copied if needed
func (d *delegate) NotifyMsg(b []byte) {
	if len(b) == 0 {
		// no data
		return
	}

	switch b[0] {
	case 'd': // data
		var updates []*update
		// unpack the message
		if err := json.Unmarshal(b[1:], &updates); err != nil {
			return
		}
		for _, u := range updates {
			for k, v := range u.Data {
				switch u.Action {
				case "add":
					d.cache.Put(k, v)
				case "del":
					d.cache.Delete(k)
				}
			}
		}
	}
}

// LocalState is used for a TCP Push/Pull. This is sent to
// the remote side in addition to the membership information. Any
// data can be sent here. See MergeRemoteState as well. The `join`
// boolean indicates this is for a join instead of a push/pull.
func (d *delegate) LocalState(join bool) []byte {
	m := d.cache.Clone()
	b, _ := json.Marshal(m)
	return b
}

// MergeRemoteState is invoked after a TCP Push/Pull. This is the state
// received from the remote side and is the result of the remote side's
// LocalState call. The 'join' boolean indicates this is for a join instead of
// a push/pull.
func (d *delegate) MergeRemoteState(buf []byte, join bool) {
	if len(buf) == 0 {
		return
	}
	if !join {
		return
	}
	var m map[string]string
	if err := json.Unmarshal(buf, &m); err != nil {
		return
	}
	for k, v := range m {
		d.cache.Put(k, v)
	}
}
