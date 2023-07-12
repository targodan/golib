package containers

type ChunkCompleteCallback[T any] func([]T) error

type Chunker[T any] struct {
	Buffer                []T
	ChunkSize             int
	chunkCompleteCallback ChunkCompleteCallback[T]
}

func NewChunker[T any](chunkSize int) *Chunker[T] {
	return &Chunker[T]{
		Buffer:    make([]T, 0, chunkSize),
		ChunkSize: chunkSize,
	}
}

func NewChunkerWithCallback[T any](chunkSize int, callback ChunkCompleteCallback[T]) *Chunker[T] {
	c := NewChunker[T](chunkSize)
	c.chunkCompleteCallback = callback
	return c
}

// Push adds a value to the Chunker. If this fills up a chunk, the complete chunk is
// returned, otherwise nil is returned.
// If the chunk is complete, and the Chunker has a ChunkCompleteCallback, the callback
// is called with the returned chunk. Any returned error is passed through.
func (c *Chunker[T]) Push(v T) ([]T, error) {
	c.Buffer = append(c.Buffer, v)
	if len(c.Buffer) >= c.ChunkSize {
		chunk := c.Buffer
		c.Buffer = make([]T, 0, c.ChunkSize)

		if c.chunkCompleteCallback != nil {
			return chunk, c.chunkCompleteCallback(chunk)
		}

		return chunk, nil
	}
	return nil, nil
}

// Close calls the registered ChunkCompleteCallback one last time, if the current buffer contains
// values. This is a noop, if no ChunkCompleteCallback was registered on creation.
func (c *Chunker[T]) Close() error {
	if c.chunkCompleteCallback != nil && len(c.Buffer) > 0 {
		return c.chunkCompleteCallback(c.Buffer)
	}
	return nil
}
