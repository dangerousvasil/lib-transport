package interceptor

import (
	"github.com/zeebo/assert"
	"testing"
)

func Test_correspondClaims(t *testing.T) {
	assert.NoError(t, correspondClaims(true, []string{"user"}, []string{"user", "bot"}))
	assert.Error(t, correspondClaims(true, []string{"user"}, []string{"user"}))
	assert.Error(t, correspondClaims(true, []string{"user"}, []string{"admin", "bot"}))
	assert.Error(t, correspondClaims(false, []string{"user"}, []string{"admin", "bot"}))
	assert.NoError(t, correspondClaims(false, []string{"user"}, []string{"user", "bot"}))
	assert.NoError(t, correspondClaims(false, []string{"user"}, []string{"user"}))
}
