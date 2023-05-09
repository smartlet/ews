package kits

import (
	"io"
	"sync"
)

var (
	DefaultBufferMinimum = 2048
	DefaultBufferMaximum = DefaultBufferMinimum * 10
)

type Buffer struct {
	off int
	len int
	cap int
	buf []byte
}

func NewBuffer(c int) *Buffer {
	return &Buffer{
		off: 0,
		len: 0,
		cap: c,
		buf: make([]byte, c),
	}
}

func (b *Buffer) GetBody() (io.ReadCloser, error) {
	b.Seek(0, io.SeekStart)
	return b, nil
}

func (b *Buffer) Len() int {
	return b.len
}

func (b *Buffer) Cap() int {
	return b.cap
}

func (b *Buffer) Data() []byte {
	return b.buf[:b.len]
}

func (b *Buffer) Bytes() []byte {
	return b.buf[b.off:b.len]
}

func (b *Buffer) Reset() {
	b.off = 0
	b.len = 0
}

func (b *Buffer) Read(p []byte) (int, error) {
	if b.off >= b.len {
		return 0, io.EOF
	}
	n := copy(p, b.buf[b.off:b.len])
	b.off += n
	return n, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
	n := len(p)
	if b.len+n > b.cap {
		buf := b.buf
		b.cap += b.cap + n
		b.buf = make([]byte, b.cap)
		copy(b.buf, buf)
	}
	b.len += copy(b.buf[b.len:], p)
	return n, nil
}

func (b *Buffer) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		b.off = int(offset)
	case io.SeekCurrent:
		b.off += int(offset)
	case io.SeekEnd:
		b.off = b.len - int(offset)
	}
	return int64(b.off), nil
}

func (b *Buffer) Close() error {
	return nil
}

var _ io.Reader = (*Buffer)(nil)
var _ io.Writer = (*Buffer)(nil)
var _ io.Seeker = (*Buffer)(nil)
var _ io.Closer = (*Buffer)(nil)

var pool = sync.Pool{
	New: func() any {
		return NewBuffer(DefaultBufferMinimum)
	},
}

func BorrowBuffer() *Buffer {
	buf := pool.Get().(*Buffer)
	buf.Reset()
	return buf
}

func ReturnBuffer(b *Buffer) {
	if b.cap < DefaultBufferMaximum {
		pool.Put(b)
	}
}
