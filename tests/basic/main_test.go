package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input = 1
	// 	output = 3
	// )

	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), input %d, actual %d", input, output, actual)
	// }

	assert.Equal(t, AddOne(2), 3, "AddOne(2) should return 3")
	// assert.NotEqual(t, 2, 3)
	// assert.Nil(t, nil, nil)
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not excuting")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("Excuting")
}