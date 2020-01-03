package nop

import (
	"github.com/stretchr/testify/assert"
	"github.com/titan-data/remote-sdk-go/remote"
	"net/url"
	"testing"
)

func TestRegistered(t *testing.T) {
	r := remote.Get("nop")
	assert.Equal(t, "nop", r.Type())
}

func TestFromURL(t *testing.T) {
	r := remote.Get("nop")
	u, _ := url.Parse("nop")
	props, err := r.FromURL(u, map[string]string{})
	assert.Equal(t, 0, len(props))
	assert.Nil(t, err)
}

func TestBadAuthority(t *testing.T) {
	r := remote.Get("nop")
	u, _ := url.Parse("nop://foo")
	_, err := r.FromURL(u, map[string]string{})
	assert.NotNil(t, err)
}

func TestBadProperty(t *testing.T) {
	r := remote.Get("nop")
	u, _ := url.Parse("nop")
	_, err := r.FromURL(u, map[string]string{"a": "b"})
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
