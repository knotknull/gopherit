package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	typeString = 1
	typeUint16 = 2
)

// TLV allows the ability to pack any message into a single binary representation
type TLV struct {
	Type   uint16
	Length uint16
	Value  []byte
}

// String is a utility helper that will print a human readable representation
// of the TLV struct
func (t *TLV) String() string {
	switch t.Type {
	case typeString:
		return fmt.Sprintf("string record: %s", t.Value)
	case typeUint16:
		i, e := t.Uint16()
		if e != nil {
			return e.Error()
		}
		return fmt.Sprintf("uint16 value %d", i)
	default:
		return fmt.Sprintf("unknown type: %d", t.Type)
	}
}

// Uint16 will return the value as a Uint16.  If the underlying value is not a
// Uint16, it will return an error.
func (t *TLV) Uint16() (uint16, error) {
	buf := bytes.NewBuffer(t.Value)
	var value uint16

	err := binary.Read(buf, binary.BigEndian, &value)
	if err != nil {
		return 0, fmt.Errorf("Invalid uint16 record:", err.Error())
	} else {
		return value, nil
	}
}

// Decode will take an `io.Reader` and parse out the `type`, `length`, and
// `value` from the provided reader.
func (t *TLV) Decode(r io.Reader) error {
	// Read in Type value
	err := binary.Read(r, binary.BigEndian, &t.Type)
	if err == io.EOF {
		return err
	} else if err != nil {
		return fmt.Errorf("error reading TLV type: %s", err.Error())
	}
	// Read in Record Length
	err = binary.Read(r, binary.BigEndian, &t.Length)
	if err != nil {
		return fmt.Errorf("error reading TLV length: %s", err.Error())
	}
	// Read in actual value
	t.Value = make([]byte, t.Length)
	_, err = io.ReadFull(r, t.Value)
	if err != nil {
		return fmt.Errorf("error reading TLV value: %s", err.Error())
		return nil
	}

	return nil
}

func main() {
	// pre-encoded value for example
	var example = "0001000d48656c6c6f2c20776f726c642e00020002002a"
	// ┌─────────┬─────────┬────────────────────────────────────────┬
	// │  TYPE   │ Length  │  Value                                 │
	// │  0001   │  000d   │  48656c6c6f2c20776f726c642e00020002002a│
	// └─────────┴─────────┴────────────────────────────────────────┴

	// decode the string from hex
	rawMessage, err := hex.DecodeString(example)
	if err != nil {
		fmt.Println("[!] couldn't decode message:", err.Error())
		return
	}
	message := bytes.NewBuffer(rawMessage)

	for {
		tlv := &TLV{}
		err := tlv.Decode(message)
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(tlv)
	}
}
