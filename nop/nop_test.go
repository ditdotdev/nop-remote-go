/*
 * Copyright Datadatdat.
 */
package nop

import (
	"testing"

	"github.com/datadatdat/remote-sdk-go/remote"
	"github.com/stretchr/testify/assert"
)

func TestRegistered(t *testing.T) {
	r := remote.Get("nop")
	ret, _ := r.Type()
	assert.Equal(t, "nop", ret)
}

func TestFromURL(t *testing.T) {
	r := remote.Get("nop")
	res, err := r.FromURL("nop", map[string]string{})

	if assert.NoError(t, err) {
		assert.Equal(t, 0, len(res))
		assert.Nil(t, err)
	}
}

func TestBadUrl(t *testing.T) {
	r := remote.Get("nop")
	_, err := r.FromURL("not\nurl", map[string]string{})
	assert.Error(t, err)
}

func TestBadAuthority(t *testing.T) {
	r := remote.Get("nop")
	_, err := r.FromURL("nop://foo", map[string]string{})
	assert.Error(t, err)
}

func TestBadProperty(t *testing.T) {
	r := remote.Get("nop")
	_, err := r.FromURL("nop", map[string]string{"a": "b"})
	assert.Error(t, err)
}

func TestToURL(t *testing.T) {
	r := remote.Get("nop")
	u, props, err := r.ToURL(map[string]interface{}{})

	if assert.NoError(t, err) {
		assert.Equal(t, "nop", u)
		assert.Empty(t, props)
	}
}

func TestGetParameters(t *testing.T) {
	r := remote.Get("nop")
	res, err := r.GetParameters(map[string]interface{}{})

	if assert.NoError(t, err) {
		assert.Empty(t, res)
	}
}

func TestValidateRemoteSuccess(t *testing.T) {
	r := remote.Get("nop")
	err := r.ValidateRemote(map[string]interface{}{})
	assert.NoError(t, err)
}

func TestValidateRemoteFailure(t *testing.T) {
	r := remote.Get("nop")
	err := r.ValidateRemote(map[string]interface{}{"a": "b"})
	assert.Error(t, err)
}

func TestValidateParametersSuccess(t *testing.T) {
	r := remote.Get("nop")
	err := r.ValidateParameters(map[string]interface{}{})
	assert.NoError(t, err)
}

func TestValidateParametersFailure(t *testing.T) {
	r := remote.Get("nop")
	err := r.ValidateParameters(map[string]interface{}{"a": "b"})
	assert.Error(t, err)
}

func TestListCommits(t *testing.T) {
	r := remote.Get("nop")
	res, err := r.ListCommits(map[string]interface{}{}, map[string]interface{}{}, []remote.Tag{})

	if assert.NoError(t, err) {
		assert.Len(t, res, 0)
	}
}

func TestGetCommit(t *testing.T) {
	r := remote.Get("nop")
	res, err := r.GetCommit(map[string]interface{}{}, map[string]interface{}{}, "id")

	if assert.NoError(t, err) {
		assert.Equal(t, "id", res.ID)
		assert.Len(t, res.Properties, 0)
	}
}

func TestPush(t *testing.T) {
	n := nopRemote{}
	err := n.Push(map[string]interface{}{}, map[string]interface{}{}, []remote.Commit{}, []remote.Tag{})
	assert.NoError(t, err)
}

func TestListTags(t *testing.T) {
	n := nopRemote{}
	tags, err := n.ListTags(map[string]interface{}{}, map[string]interface{}{})
	if assert.NoError(t, err) {
		assert.Empty(t, tags)
	}
}

func TestGetTag(t *testing.T) {
	n := nopRemote{}
	tag, err := n.GetTag(map[string]interface{}{}, map[string]interface{}{}, "myTag")
	if assert.NoError(t, err) {
		assert.Equal(t, "myTag", tag.Key)
		assert.Nil(t, tag.Value)
	}
}
