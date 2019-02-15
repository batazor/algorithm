package circularBuffer

import "errors"

// WARNING: use https://golang.org/pkg/container/ring/
// source code: https://github.com/armon/circbuf/blob/master/circbuf.go

// Buffer implements a circular buffer. It is a fixed size,
// and new writes overwrite older data, such that for a buffer
// of size N, for any amount of writes, only the last N bytes
// are retained.
type Buffer struct {
	data        []byte
	size        int
	writeCursor int
	written     int
}

// New creates a new buffer of a given size.
func New(size int) (*Buffer, error) {
	if size <= 0 {
		return nil, errors.New("Incorrect size for new buffer. Need a value > 0")
	}

	return &Buffer{
		size: size,
		data: make([]byte, size),
	}, nil
}

// Write writes up to len(buf) bytes to the internal ring,
// overriding older data if necessary.
func (b *Buffer) Write(buf []byte) (int, error) {
	// account for total bytes written
	n := len(buf)
	b.written += n

	// if the buffer is larger than ours, then we only care
	// about the last size bytes anyways
	if n > b.size {
		buf = buf[n-b.size:]
	}

	// Copy in place
	remain := b.size - b.writeCursor
	copy(b.data[b.writeCursor:], buf)
	if n > remain {
		copy(b.data, buf[remain:])
	}

	// Update location of the cursor
	b.writeCursor = ((b.writeCursor + n) % b.size)
	return n, nil
}

// Size returns the size of the buffer
func (b *Buffer) Size() int64 {
	return b.size
}

// TotalWritten provides the total number of bytes written
func (b *Buffer) TotalWritten() int64 {
	return b.written
}

// Bytes provides a slice of the bytes written. This
// slice should not be written to.
func (b *Buffer) Bytes() []byte {
	switch {
	case b.written >= b.size && b.writeCursor == 0:
		return b.data
	case b.written > b.size:
		out := make([]byte, b.size)
		copy(out, b.data[b.writeCursor:])
		copy(out[b.size-b.writeCursor:], b.data[:b.writeCursor])
		return out
	default:
		return b.data[:b.writeCursor]
	}
}

// Reset resets the buffer so it has no content.
func (b *Buffer) Reset() {
	b.writeCursor = 0
	b.written = 0
}

// String returns the contents of the buffer as a string
func (b *Buffer) String() string {
	return string(b.Bytes())
}
