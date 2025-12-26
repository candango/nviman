package protocol

import (
	"testing"

	"github.com/candango/httpok/testrunner"
	"github.com/stretchr/testify/assert"
)

func TestGithubTransport(t *testing.T) {
	gt, err := NewGithubTransport()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("NewNonceUrl should return url if DirectoryKey is valid and error otherwise", func(t *testing.T) {
		res, err := gt.GetReleases()
		if err != nil {
			t.Error(err)
		}
		assert.NoError(t, err)
		releases := []map[string]any{}
		testrunner.BodyAsJson(t, res, &releases)
	})
}
