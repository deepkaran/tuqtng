//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package xpipeline

import (
	"testing"

	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/misc"
)

var testData = dparval.ValueCollection{}

func init() {
	doc := dparval.NewValue(map[string]interface{}{"name": "mike", "age": 100.0})
	doc.SetAttachment("meta", map[string]interface{}{"id": "1"})
	testData = append(testData, doc)

	doc = dparval.NewValue(map[string]interface{}{"name": "dustin"})
	doc.SetAttachment("meta", map[string]interface{}{"id": "1"})
	testData = append(testData, doc)

	doc = dparval.NewValue(map[string]interface{}{"name": "bob", "age": nil})
	doc.SetAttachment("meta", map[string]interface{}{"id": "1"})
	testData = append(testData, doc)

	doc = dparval.NewValue(map[string]interface{}{"name": "marty", "age": 99.0})
	doc.SetAttachment("meta", map[string]interface{}{"id": "1"})
	testData = append(testData, doc)

	doc = dparval.NewValue(map[string]interface{}{"name": "steve", "age": 200.0})
	doc.SetAttachment("meta", map[string]interface{}{"id": "2"})
	testData = append(testData, doc)

	doc = dparval.NewValue(map[string]interface{}{"name": "gerald", "age": 175.0})
	doc.SetAttachment("meta", map[string]interface{}{"id": "3"})
	testData = append(testData, doc)

	doc = dparval.NewValue(map[string]interface{}{"name": "siri", "age": 74.0})
	doc.SetAttachment("meta", map[string]interface{}{"id": "4"})
	testData = append(testData, doc)

	doc = dparval.NewValue(map[string]interface{}{"name": "ali", "age": 100.0})
	doc.SetAttachment("meta", map[string]interface{}{"id": "1"})
	testData = append(testData, doc)
}

func TestFilterTrue(t *testing.T) {

	stubSource := NewStubSource(testData)

	filter := NewFilter(ast.NewLiteralBool(true))
	filter.SetSource(stubSource)

	filterItemChannel, _ := filter.GetChannels()

	stopChannel := make(misc.StopChannel)
	go filter.Run(stopChannel)

	count := 0
	for _ = range filterItemChannel {
		count++
	}

	if count != len(testData) {
		t.Errorf("Expected %d items, got %d", len(testData), count)
	}

}

func TestFilterFalse(t *testing.T) {

	stubSource := NewStubSource(testData)

	filter := NewFilter(ast.NewLiteralBool(false))
	filter.SetSource(stubSource)

	filterItemChannel, _ := filter.GetChannels()

	stopChannel := make(misc.StopChannel)
	go filter.Run(stopChannel)

	count := 0
	for _ = range filterItemChannel {
		count++
	}

	if count != 0 {
		t.Errorf("Expected %d items, got %d", 0, count)
	}

}

func TestFilterSome(t *testing.T) {

	stubSource := NewStubSource(testData)

	filter := NewFilter(ast.NewGreaterThanOperator(ast.NewProperty("name"), ast.NewLiteralString("n")))
	filter.SetSource(stubSource)

	filterItemChannel, _ := filter.GetChannels()

	stopChannel := make(misc.StopChannel)
	go filter.Run(stopChannel)

	count := 0
	for _ = range filterItemChannel {
		count++
	}

	if count != 2 {
		t.Errorf("Expected %d items, got %d", 2, count)
	}

}
