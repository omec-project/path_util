package version_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/omec-project/path_util/version"
)

func TestVersion(t *testing.T) {
	assert.Equal(t, "2020-03-31-01", version.GetVersion())
}
