// Copyright 2021 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package inventory

import (
	"testing"

	"github.com/mendersoftware/mender/client/api"
	"github.com/stretchr/testify/assert"
)

func TestInventoryDataDecoder(t *testing.T) {

	idec := NewInventoryDataDecoder()
	assert.NotNil(t, idec)

	idec.AppendFromRaw(map[string][]string{
		"foo": []string{"bar"},
	})

	assert.Contains(t, idec.GetInventoryData(), api.InventoryAttribute{
		Name:  "foo",
		Value: "bar",
	})

	idec.AppendFromRaw(map[string][]string{
		"foo": []string{"baz"},
	})
	assert.Contains(t, idec.data, "foo")
	assert.Contains(t, idec.GetInventoryData(),
		api.InventoryAttribute{
			Name:  "foo",
			Value: []string{"bar", "baz"}})

	idec.AppendFromRaw(map[string][]string{
		"bar": []string{"zen"},
	})
	assert.Contains(t, idec.GetInventoryData(),
		api.InventoryAttribute{
			Name:  "foo",
			Value: []string{"bar", "baz"}})
	assert.Contains(t, idec.GetInventoryData(), api.InventoryAttribute{
		Name:  "bar",
		Value: "zen",
	})

	idata := idec.GetInventoryData()
	assert.Len(t, idata, 2)
	assert.Contains(t, idata, api.InventoryAttribute{
		Name:  "foo",
		Value: []string{"bar", "baz"}})
	assert.Contains(t, idata, api.InventoryAttribute{
		Name:  "bar",
		Value: "zen"})
}
