package uevent

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// Uevent represents a single uevent event.
type Uevent struct {
	Header string

	// default uevent keys according to kobject_uevent.c
	Action    string
	Devpath   string
	Subsystem string
	Seqnum    string

	// All keys of this uevent
	Keys map[string]string
}

type Decoder struct {
	r *bufio.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{bufio.NewReader(r)}
}

func (d *Decoder) Decode() (*Uevent, error) {
	ev := &Uevent{
		Keys: map[string]string{},
	}

loop:
	for {
		s, err := d.r.ReadString(0x00)
		if err != nil {
			return nil, err
		}

		switch {
		case strings.Contains(s, "@"):
			ev.Header = s
		case strings.Contains(s, "="):
			kv := strings.Split(s, "=")

			if len(kv) != 2 {
				return nil, errors.New("error reading event: unknown format")
			}

			k, v := kv[0], kv[1]

			ev.Keys[k] = v

			switch k {
			case "ACTION":
				ev.Action = v
			case "DEVPATH":
				ev.Devpath = v
			case "SUBSYSTEM":
				ev.Subsystem = v
			case "SEQNUM":
				ev.Seqnum = v
				break loop
			}
		}
	}

	return ev, nil
}
