/*
Copyright 2023 The KubeService-Stack Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"crypto"
	"testing"

	sigs "github.com/sigstore/cosign/pkg/signature"
	"github.com/stretchr/testify/assert"
)

func Test_verifyPemDecode(t *testing.T) {
	assert := assert.New(t)
	aa := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDWRqOf5XV7h2Ae7eQwHl7OsNMRwK
08eOOwEkJXIQAgFJETF28GurSzLEs3GlTZuoo89yfuW73PjlY/1xbWZ3og==
-----END PUBLIC KEY-----`
	k, err := sigs.LoadPublicKeyRaw([]byte(aa), crypto.SHA256)
	assert.Nil(err)
	assert.NotEqual(k, "")
}

func Test_verifyPemDecodeString(t *testing.T) {
	assert := assert.New(t)
	aa := "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDWRqOf5XV7h2Ae7eQwHl7OsNMRwK\n08eOOwEkJXIQAgFJETF28GurSzLEs3GlTZuoo89yfuW73PjlY/1xbWZ3og==\n-----END PUBLIC KEY-----\n"
	k, err := sigs.LoadPublicKeyRaw([]byte(aa), crypto.SHA256)
	assert.Nil(err)
	assert.NotEqual(k, "")
}

func Test_verifyPemDecodeError(t *testing.T) {
	assert := assert.New(t)
	aa := "dongjiang"
	k, err := sigs.LoadPublicKeyRaw([]byte(aa), crypto.SHA256)
	assert.NotNil(err)
	assert.Empty(k)
}

func Test_VerifyPublicKey(t *testing.T) {
	assert := assert.New(t)
	aa := "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDWRqOf5XV7h2Ae7eQwHl7OsNMRwK\n08eOOwEkJXIQAgFJETF28GurSzLEs3GlTZuoo89yfuW73PjlY/1xbWZ3og==\n-----END PUBLIC KEY-----\n"
	k, err := VerifyPublicKey("dongjiang1989/node-metrics@sha256:c1aa0f2861d2eb744efb8f82a1d7d5f1b74919d1cc6501e799daeac1991fc282", aa)
	assert.Nil(err)
	assert.True(k)
}

func Test_VerifyPublicKey2(t *testing.T) {
	assert := assert.New(t)
	aa := "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDWRqOf5XV7h2Ae7eQwHl7OsNMRwK\n08eOOwEkJXIQAgFJETF28GurSzLEs3GlTZuoo89yfuW73PjlY/1xbWZ3og==\n-----END PUBLIC KEY-----\n"
	k, err := VerifyPublicKey("dongjiang1989/node-metrics:latest", aa)
	assert.Nil(err)
	assert.True(k)
}

func Test_VerifyPublicKeyError(t *testing.T) {
	assert := assert.New(t)
	aa := "aaaa\n"
	k, err := VerifyPublicKey("dongjiang1989/node-metrics:latest", aa)
	assert.NotNil(err)
	assert.False(k)
}

func Test_VerifyPublicKey3(t *testing.T) {
	assert := assert.New(t)
	aa := "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDWRqOf5XV7h2Ae7eQwHl7OsNMRwK\n08eOOwEkJXIQAgFJETF28GurSzLEs3GlTZuoo89yfuW73PjlY/1xbWZ3og==\n-----END PUBLIC KEY-----\n"
	k, err := VerifyPublicKey("dongjiang1989/pingmesh-agent:latest", aa)
	assert.NotNil(err)
	assert.False(k)
}
