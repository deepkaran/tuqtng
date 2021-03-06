//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import (
	"reflect"
	"testing"
)

func TestExpressionUnmarshal(t *testing.T) {
	tests := []struct {
		input string
		expr  Expression
	}{
		{
			`{
				"type": "property",
				"path": "beer-sample"
            }`,
			NewProperty("beer-sample"),
		},
		{
			`{
				"type": "dot_member",
				"left": {
					"type": "property",
					"path": "beer-sample"
				},
				"right": {
					"type": "property",
					"path": "name"
				}
			}`,
			NewDotMemberOperator(NewProperty("beer-sample"), NewProperty("name")),
		},
		{
			`{
				"type": "bracket_member",
				"left": {
					"type": "property",
					"path": "beer-sample"
				},
				"right": {
					"type": "literal_number",
					"value": 1
				}
			}`,
			NewBracketMemberOperator(NewProperty("beer-sample"), NewLiteralNumber(1.0)),
		},
		{
			`{
				"type": "bracket_member",
				"left": {
					"type": "dot_member",
					"left": {
						"type": "property",
                        "path": "beer-sample"
                    },
                    "right": {
                        "type": "property",
                        "path": "name"
	                }
                },
                "right": {
                    "type": "literal_number",
                    "value": 1
                }
            }`,
			NewBracketMemberOperator(NewDotMemberOperator(NewProperty("beer-sample"), NewProperty("name")), NewLiteralNumber(1.0)),
		},
	}

	for _, test := range tests {
		output, err := UnmarshalExpression([]byte(test.input))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(output, test.expr) {
			t.Errorf("Expected %v, got %v for %s", test.expr, output, test.input)
		}
	}
}
