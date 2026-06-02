/*
 * Copyright Dit.
 */
package nop

import (
	"testing"

	"github.com/ditdotdev/remote-sdk-go/remote"
	"github.com/stretchr/testify/assert"
)

// getNopRemote retrieves the registered nop remote, failing the test if it is
// not registered. This ensures a broken init() registration is not silently
// masked by tests that discard the bool from remote.Get.
func getNopRemote(t *testing.T) remote.Remote {
	t.Helper()
	r, ok := remote.Get("nop")
	if !ok {
		t.Fatal("nop remote not registered")
	}
	return r
}

func TestRegistered(t *testing.T) {
	r := getNopRemote(t)
	ret, _ := r.Type()
	assert.Equal(t, "nop", ret)
}

func TestFromURL(t *testing.T) {
	r := getNopRemote(t)
	res, err := r.FromURL("nop", map[string]string{})

	if assert.NoError(t, err) {
		assert.Equal(t, 0, len(res))
		assert.Nil(t, err)
	}
}

func TestFromURLNilProperties(t *testing.T) {
	r := getNopRemote(t)
	res, err := r.FromURL("nop", nil)

	if assert.NoError(t, err) {
		assert.Equal(t, 0, len(res))
	}
}

func TestBadUrl(t *testing.T) {
	r := getNopRemote(t)
	_, err := r.FromURL("not\nurl", map[string]string{})
	assert.Error(t, err)
}

func TestBadAuthority(t *testing.T) {
	r := getNopRemote(t)
	_, err := r.FromURL("nop://foo", map[string]string{})
	assert.Error(t, err)
}

func TestBadProperty(t *testing.T) {
	r := getNopRemote(t)
	_, err := r.FromURL("nop", map[string]string{"a": "b"})
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "'a'")
	}
}

func TestToURL(t *testing.T) {
	r := getNopRemote(t)
	u, props, err := r.ToURL(map[string]interface{}{})

	if assert.NoError(t, err) {
		assert.Equal(t, "nop", u)
		assert.Empty(t, props)
	}
}

func TestToURLNilProperties(t *testing.T) {
	r := getNopRemote(t)
	u, props, err := r.ToURL(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "nop", u)
		assert.Empty(t, props)
	}
}

func TestGetParameters(t *testing.T) {
	r := getNopRemote(t)
	res, err := r.GetParameters(map[string]interface{}{})

	if assert.NoError(t, err) {
		assert.Empty(t, res)
	}
}

func TestGetParametersNilProperties(t *testing.T) {
	r := getNopRemote(t)
	res, err := r.GetParameters(nil)

	if assert.NoError(t, err) {
		assert.Empty(t, res)
	}
}

func TestValidateRemoteSuccess(t *testing.T) {
	r := getNopRemote(t)
	err := r.ValidateRemote(map[string]interface{}{})
	assert.NoError(t, err)
}

func TestValidateRemoteNilProperties(t *testing.T) {
	r := getNopRemote(t)
	err := r.ValidateRemote(nil)
	assert.NoError(t, err)
}

func TestValidateRemoteFailure(t *testing.T) {
	r := getNopRemote(t)
	err := r.ValidateRemote(map[string]interface{}{"a": "b"})
	assert.Error(t, err)
}

func TestValidateParametersSuccess(t *testing.T) {
	r := getNopRemote(t)
	err := r.ValidateParameters(map[string]interface{}{})
	assert.NoError(t, err)
}

func TestValidateParametersNilParameters(t *testing.T) {
	r := getNopRemote(t)
	err := r.ValidateParameters(nil)
	assert.NoError(t, err)
}

func TestValidateParametersDelayAllowed(t *testing.T) {
	r := getNopRemote(t)
	err := r.ValidateParameters(map[string]interface{}{"delay": "5"})
	assert.NoError(t, err)
}

func TestValidateParametersFailure(t *testing.T) {
	r := getNopRemote(t)
	err := r.ValidateParameters(map[string]interface{}{"a": "b"})
	assert.Error(t, err)
}

func TestListCommits(t *testing.T) {
	r := getNopRemote(t)
	res, err := r.ListCommits(map[string]interface{}{}, map[string]interface{}{}, []remote.Tag{})

	if assert.NoError(t, err) {
		assert.Len(t, res, 0)
	}
}

func TestGetCommit(t *testing.T) {
	r := getNopRemote(t)
	res, err := r.GetCommit(map[string]interface{}{}, map[string]interface{}{}, "id")

	if assert.NoError(t, err) {
		assert.Equal(t, "id", res.ID)
		assert.Len(t, res.Properties, 0)
	}
}
