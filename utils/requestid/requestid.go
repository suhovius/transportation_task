package requestid

import (
	"fmt"

	"github.com/google/uuid"
)

type key int

const (
	// RequestIDKey is used to obtain request id from the context
	RequestIDKey key = 0
)

// Next returns next request ID
func Next() string {
	return fmt.Sprintf("%s", uuid.New())
}
