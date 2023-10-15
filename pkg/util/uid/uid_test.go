package uid

import (
	"fmt"
	"testing"
)

func TestULID(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(ULID())
		}()
		go func() {
			fmt.Println(ULID())
		}()
		go func() {
			fmt.Println(ULID())
		}()
		go func() {
			fmt.Println(ULID())
		}()
		go func() {
			fmt.Println(ULID())
		}()
		go func() {
			fmt.Println(ULID())
		}()
		go func() {
			fmt.Println(ULID())
		}()
	}
}
