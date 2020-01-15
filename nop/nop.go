/*
 * Copyright The Titan Project Contributors.
 */
package nop

import (
	"errors"
	"fmt"
	"github.com/titan-data/remote-sdk-go/remote"
	"net/url"
	"reflect"
)

type nopRemote struct {
}

func (n nopRemote) Type() (string, error) {
	return "nop", nil
}

func (n nopRemote) FromURL(rawUrl string, properties map[string]string) (map[string]interface{}, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	// nop remotes can only be "nop", which means everything other than "path" must be empty
	if u.Scheme != "" || u.Host != "" || u.User != nil || u.Path != "nop" {
		return nil, errors.New("malformed remote")
	}

	if len(properties) != 0 {
		return nil, errors.New(fmt.Sprintf("invalid property '%s'", reflect.ValueOf(properties).MapKeys()[0].String()))
	}

	return map[string]interface{}{}, nil
}

func (n nopRemote) ToURL(_ map[string]interface{}) (string, map[string]string, error) {
	return "nop", map[string]string{}, nil
}

func (n nopRemote) GetParameters(_ map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func (n nopRemote) ValidateRemote(properties map[string]interface{}) error {
	for k := range properties {
		return fmt.Errorf("invalid remote property '%s'", k)
	}
	return nil
}

func (n nopRemote) ValidateParameters(parameters map[string]interface{}) error {
	for k := range parameters {
		if k != "delay" {
			return fmt.Errorf("invalid parameters property '%s'", k)
		}
	}
	return nil
}

func (n nopRemote) ListCommits(_ map[string]interface{}, _ map[string]interface{}, _ []remote.Tag) ([]remote.Commit, error) {
	return []remote.Commit{}, nil
}

func (n nopRemote) GetCommit(_ map[string]interface{}, _ map[string]interface{}, commitId string) (*remote.Commit, error) {
	return &remote.Commit{
		Id:         commitId,
		Properties: map[string]interface{}{},
	}, nil
}

func init() {
	remote.Register(nopRemote{})
}
