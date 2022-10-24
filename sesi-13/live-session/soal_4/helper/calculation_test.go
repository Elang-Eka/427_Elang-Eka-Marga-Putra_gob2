package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHitung(t *testing.T) {
	// test for empty result
	result := Hello(10, 10, 50, 50)
	require.Equal(t, 1010, result, "Result has to be 1010, a=10 b=10 c=50 e=50")

	fmt.Println("Test Eksekusi Terhenti")
}
