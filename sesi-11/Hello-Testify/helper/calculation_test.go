package helper

import (
	"fmt"
	"testing"
)

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

// bagian 3 Test Table
func TestTableSum(t *testing.T) {
	tests := []struct {
		request  int
		expected int
		errMsg   string
	}{
		{
			request:  sum(10, 10),
			expected: 20,
			errMsg:   "Result has to be 20",
		},
		{
			request:  sum(20, 20),
			expected: 40,
			errMsg:   "Result has to be 40",
		},
		{
			request:  sum(25, 25),
			expected: 50,
			errMsg:   "Result has to be 50",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errMsg)
		})
	}
}
