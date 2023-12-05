package clean

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCleanDirectory(t *testing.T) {
	dirToTest := "Test"
	err := os.Mkdir(dirToTest, 777)
	assert.Nil(t, err)

	// Testing that we can successfully delete the "Test" directory
	err = Dist(dirToTest)
	assert.Nil(t, err)

	// Testing that we exit if the directory isn't provided
	err = Dist("")
	assert.NotNil(t, err)

	// Testing fake directory that does not exists
	err = Dist("aeed/asd//")
	assert.Nil(t, err)
}
