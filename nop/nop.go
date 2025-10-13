/*
 * Copyright The Titan Project Contributors.
 */
package nop

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"

	"github.com/datadatdat/remote-sdk-go/remote"
)

const remoteType = "nop"

type nopRemote struct{}

// Type returns the remote type identifier
func (n nopRemote) Type() (string, error) {
	return remoteType, nil
}

// FromURL parses a URL and converts it to remote properties
func (n nopRemote) FromURL(rawUrl string, properties map[string]string) (map[string]interface{}, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	// nop remotes can only be "nop", which means everything other than "path" must be empty
	if u.Scheme != "" || u.Host != "" || u.User != nil || u.Path != remoteType {
		return nil, errors.New("malformed remote")
	}

	if len(properties) != 0 {
		return nil, fmt.Errorf("invalid property '%s'", reflect.ValueOf(properties).MapKeys()[0].String())
	}

	return map[string]interface{}{}, nil
}

// ToURL converts remote properties back to URL format
func (n nopRemote) ToURL(_ map[string]interface{}) (string, map[string]string, error) {
	return remoteType, map[string]string{}, nil
}

// GetParameters returns the parameters for remote operations
func (n nopRemote) GetParameters(_ map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

// ValidateRemote validates remote configuration properties
func (n nopRemote) ValidateRemote(properties map[string]interface{}) error {
	for k := range properties {
		return fmt.Errorf("invalid remote property '%s'", k)
	}

	return nil
}

// ValidateParameters validates operation parameters
func (n nopRemote) ValidateParameters(parameters map[string]interface{}) error {
	for k := range parameters {
		if k != "delay" {
			return fmt.Errorf("invalid parameters property '%s'", k)
		}
	}

	return nil
}

// ListCommits returns a list of commits from the remote
func (n nopRemote) ListCommits(_ map[string]interface{}, _ map[string]interface{}, _ []remote.Tag) ([]remote.Commit, error) {
	return []remote.Commit{}, nil
}

// GetCommit retrieves a specific commit from the remote
func (n nopRemote) GetCommit(_ map[string]interface{}, _ map[string]interface{}, commitId string) (*remote.Commit, error) {
	return &remote.Commit{
		ID:         commitId,
		Properties: map[string]interface{}{},
	}, nil
}

// Push sends commits and tags to the remote
func (n nopRemote) Push(_ map[string]interface{}, _ map[string]interface{}, commits []remote.Commit, tags []remote.Tag) error {
	return nil
}

// ListTags returns a list of tags from the remote
func (n nopRemote) ListTags(_ map[string]interface{}, _ map[string]interface{}) ([]remote.Tag, error) {
	return []remote.Tag{}, nil
}

// GetTag retrieves a specific tag from the remote
func (n nopRemote) GetTag(_ map[string]interface{}, _ map[string]interface{}, tagName string) (*remote.Tag, error) {
	return &remote.Tag{Key: tagName, Value: nil}, nil
}

func init() {
	remote.Register(nopRemote{})
}
