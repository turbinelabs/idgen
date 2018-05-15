/*
Copyright 2018 Turbine Labs, Inc.

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

package idgen

import (
	"testing"

	"github.com/turbinelabs/test/assert"
)

func TestCounter(t *testing.T) {
	gen := NewCounter(0)
	want := ID("1")
	got, err := gen()

	assert.Nil(t, err)
	assert.Equal(t, got, want)

	want = ID("2")
	got, err = gen()

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestUUID(t *testing.T) {
	gen := NewUUID()
	first, err := gen()
	assert.Nil(t, err)

	second, err := gen()

	assert.Nil(t, err)
	assert.NotEqual(t, first, second)
}

func TestPrefixing(t *testing.T) {
	gen := NewPrefixing("foo-", NewCounter(0))
	want := ID("foo-1")
	got, err := gen()

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}
