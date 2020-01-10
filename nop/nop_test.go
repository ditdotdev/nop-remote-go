package nop

import (
	"github.com/stretchr/testify/assert"
	"github.com/titan-data/remote-sdk-go/remote"
	"testing"
)

func TestRegistered(t *testing.T) {
	r := remote.Get("nop")
	ret, _ := r.Type()
	assert.Equal(t, "nop", ret)
}

func TestFromURL(t *testing.T) {
	r := remote.Get("nop")
	props, err := r.FromURL("nop", map[string]string{})
	assert.Equal(t, 0, len(props))
	assert.Nil(t, err)
}

func TestBadAuthority(t *testing.T) {
	r := remote.Get("nop")
	_, err := r.FromURL("nop://foo", map[string]string{})
	assert.NotNil(t, err)
}

func TestBadProperty(t *testing.T) {
	r := remote.Get("nop")
	_, err := r.FromURL("nop", map[string]string{"a": "b"})
	assert.NotNil(t, err)
}

func TestToURL(t *testing.T) {
	r := remote.Get("nop")
	u, props, _ := r.ToURL(map[string]interface{}{})
	assert.Equal(t, "nop", u)
	assert.Empty(t, props)
}

func TestGetParameters(t *testing.T) {
	r := remote.Get("nop")
	params, _ := r.GetParameters(map[string]interface{}{})
	assert.Empty(t, params)
}
