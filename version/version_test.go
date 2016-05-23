package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionSetter(t *testing.T) {
	assert := assert.New(t)

	SetVersion("first")
	assert.Equal("first", Version)

	SetVersion("second")
	assert.Equal("first", Version)
}

func TestCommitSetter(t *testing.T) {
	assert := assert.New(t)

	SetCommit("first")
	assert.Equal("first", Commit)

	SetCommit("second")
	assert.Equal("first", Commit)
}

func TestBranchSetter(t *testing.T) {
	assert := assert.New(t)

	SetBranch("first")
	assert.Equal("first", Branch)

	SetBranch("second")
	assert.Equal("first", Branch)
}
