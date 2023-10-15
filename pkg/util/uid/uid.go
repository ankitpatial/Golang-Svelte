package uid

import (
	"bytes"
	"crypto/rand"
	"github.com/oklog/ulid/v2"
	"sync"
	"time"
)

var randPool = sync.Pool{
	New: func() any {
		b := make([]byte, 32)
		_, _ = rand.Read(b)
		return bytes.NewReader(b)
	},
}

// ULID is Universally Unique Lexicographically Sortable Identifier
// - Is compatible with UUID/GUID's
// - 1.21e+24 unique ULIDs per millisecond (1,208,925,819,614,629,174,706,176 to be exact)
// - Lexicographically sortable
// - Canonically encoded as a 26 character string, as opposed to the 36 character UUID
// - Uses Crockford's base32 for better efficiency and readability (5 bits per character)
// - Case insensitive
// - No special characters (URL safe)
// - Monotonic sort order (correctly detects and handles the same millisecond)
func ULID() string {
	entropy := randPool.Get().(*bytes.Reader)
	ms := ulid.Timestamp(time.Now())
	if id, err := ulid.New(ms, entropy); err != nil {
		panic(err)
	} else {
		return id.String()
	}
}
