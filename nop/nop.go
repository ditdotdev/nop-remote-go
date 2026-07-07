// Copyright Dit 2026
// SPDX-License-Identifier: BUSL-1.1

package nop

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/ditdotdev/remote-sdk-go/remote"
)

const remoteType = "nop"

type nopRemote struct{}

// Type returns the remote type identifier
func (nopRemote) Type() (string, error) {
	return remoteType, nil
}

// FromURL parses a URL and converts it to remote properties
func (nopRemote) FromURL(rawUrl string, properties map[string]string) (map[string]interface{}, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	// nop remotes can only be "nop", which means everything other than "path" must be empty
	if u.Scheme != "" || u.Host != "" || u.User != nil || u.Path != remoteType {
		return nil, errors.New("malformed remote")
	}

	for k := range properties {
		return nil, fmt.Errorf("invalid property '%s'", k)
	}

	return map[string]interface{}{}, nil
}

// ToURL converts remote properties back to URL format
func (nopRemote) ToURL(_ map[string]interface{}) (string, map[string]string, error) {
	return remoteType, map[string]string{}, nil
}

// GetParameters returns the parameters for remote operations
func (nopRemote) GetParameters(_ map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

// ValidateRemote validates remote configuration properties
func (nopRemote) ValidateRemote(properties map[string]interface{}) error {
	for k := range properties {
		return fmt.Errorf("invalid remote property '%s'", k)
	}

	return nil
}

// ValidateParameters validates operation parameters
func (nopRemote) ValidateParameters(parameters map[string]interface{}) error {
	for k := range parameters {
		if k != "delay" {
			return fmt.Errorf("invalid parameters property '%s'", k)
		}
	}

	return nil
}

// ListCommits returns a list of commits from the remote
func (nopRemote) ListCommits(_ map[string]interface{}, _ map[string]interface{}, _ []remote.Tag) ([]remote.Commit, error) {
	return []remote.Commit{}, nil
}

// GetCommit retrieves a specific commit from the remote
func (nopRemote) GetCommit(_ map[string]interface{}, _ map[string]interface{}, commitId string) (*remote.Commit, error) {
	return &remote.Commit{
		ID:         commitId,
		Properties: map[string]interface{}{},
	}, nil
}

func init() {
	remote.Register(nopRemote{})
}
