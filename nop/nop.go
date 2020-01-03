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

func (n nopRemote) Type() string {
	return "nop"
}

func (n nopRemote) FromURL(url *url.URL, additionalProperties map[string]string) (map[string]interface{}, error) {
	// nop remotes can only be "nop", which means everything other than "path" must be empty
	if url.Scheme != "" || url.Host != "" || url.User != nil || url.Path != "nop" {
		return nil, errors.New("malformed remote")
	}

	if len(additionalProperties) != 0 {
		return nil, errors.New(fmt.Sprintf("invalid property '%s'", reflect.ValueOf(additionalProperties).MapKeys()[0].String()))
	}

	return map[string]interface{}{}, nil
}

func (n nopRemote) ToURL(properties map[string]interface{}) (string, map[string]string, error) {
	return "nop", map[string]string{}, nil
}

func (n nopRemote) GetParameters(remoteProperties map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func init() {
	remote.Register(nopRemote{})
}
