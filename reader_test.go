package reversereader

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func Test_reader_Read(t *testing.T) {
	type in struct {
		rs io.ReadSeeker
		p  []byte
	}
	type out struct {
		p   []byte
		n   int
		err error
	}

	cases := []struct {
		in  in
		exp out
	}{
		{
			in: in{
				rs: bytes.NewReader([]byte{0x01, 0x02, 0x03, 0x04}),
				p:  make([]byte, 0),
			},
			exp: out{
				p:   []byte{},
				n:   0,
				err: nil,
			},
		},
		{
			in: in{
				rs: bytes.NewReader([]byte{0x01, 0x02, 0x03, 0x04}),
				p:  make([]byte, 5),
			},
			exp: out{
				p:   []byte{0x04, 0x03, 0x02, 0x01, 0x00},
				n:   4,
				err: io.EOF,
			},
		},
		{
			in: in{
				rs: bytes.NewReader([]byte{0x01, 0x02, 0x03, 0x04}),
				p:  make([]byte, 4),
			},
			exp: out{
				p:   []byte{0x04, 0x03, 0x02, 0x01},
				n:   4,
				err: io.EOF,
			},
		},
		{
			in: in{
				rs: bytes.NewReader([]byte{0x01, 0x02, 0x03, 0x04}),
				p:  make([]byte, 3),
			},
			exp: out{
				p:   []byte{0x04, 0x03, 0x02},
				n:   3,
				err: nil,
			},
		},
		{
			in: in{
				rs: bytes.NewReader([]byte{}),
				p:  make([]byte, 1),
			},
			exp: out{
				p:   []byte{0x00},
				n:   0,
				err: io.EOF,
			},
		},
	}

	for i, c := range cases {
		r := NewReader(c.in.rs)

		n, err := r.Read(c.in.p)
		act := out{
			p:   c.in.p,
			n:   n,
			err: err,
		}
		exp := c.exp

		if !reflect.DeepEqual(act, exp) {
			t.Errorf("\n   index: %d\n  actual: %#v\nexpected: %#v\n", i, act, exp)
		}
	}
}
