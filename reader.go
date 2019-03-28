// Package reversereader provides basic interfaces to read.
// It traverse an io.Reader as a backward stream.
package reversereader

import "io"

// NewReader returns a io.Reader for reading underlying io.ReadSeeker backwards.
func NewReader(rs io.ReadSeeker) io.Reader {
	return &reader{
		rs:  rs,
		rem: -1,
	}
}

type reader struct {
	rs  io.ReadSeeker
	rem int64
}

// Read satisfies the io.Reader interface.
func (r *reader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	if r.rem < 0 {
		// Set remaining length
		rem, err := r.rs.Seek(0, io.SeekEnd)
		if err != nil {
			return 0, err
		}
		r.rem = rem
	}

	// offset         r.rem
	//   |              |
	//   |<-- len(p) -->|
	//   |              |
	//   |              |
	//   |              |
	//   v              v
	//        ############################## r.rs
	//   [              ] p
	//        [         ] q

	q := p[0:]
	offset := r.rem - int64(len(p))

	if offset < 0 {
		q = p[-offset:]
		offset = 0
	}

	if _, err := r.rs.Seek(offset, io.SeekStart); err != nil {
		return 0, err
	}

	if _, err := io.ReadFull(r.rs, q); err != nil {
		return 0, err
	}

	// Reverse the contents of p
	for i := 0; i < len(p)/2; i++ {
		j := len(p) - 1 - i
		p[i], p[j] = p[j], p[i]
	}

	r.rem = offset

	if r.rem == 0 {
		return len(q), io.EOF
	}

	return len(q), nil
}
