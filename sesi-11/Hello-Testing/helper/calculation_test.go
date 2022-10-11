package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
// bagian 1
func TestFailSum(t *testing.T) {
	result := sum(10, 10)

	if result != 40 {
		// fail method
		// t.Fail()
		// failNow method
		// t.FailNow()
		// error method
		t.Error("Result has to be 40")
	}

	fmt.Println("TestFailSum Eksekusi Terhenti")
}


*/

// bagian 2 (Testify (assert and require))
func TestFailSum(t *testing.T) {
	result := sum(10, 10)

	require.Equal(t, 40, result, "Result has to be 40")

	fmt.Println("TestFailSum Eksekusi Terhenti")
}

func TestSum(t *testing.T) {
	result := sum(10, 10)

	assert.Equal(t, 20, result, "Result has to be 20")

	fmt.Println("TestSum Eksekusi Terhenti")
}
