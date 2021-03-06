//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package main

import (
	"testing"
	"time"

	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/tuqtng/test"
)

type BenchmarkMockQuery struct {
	request     network.QueryRequest
	response    *MockBenchmarkResponse
	stopChannel misc.StopChannel
	startTime   time.Time
}

func (this *BenchmarkMockQuery) Request() network.QueryRequest {
	return this.request
}

func (this *BenchmarkMockQuery) Response() network.QueryResponse {
	return this.response
}

func (this *BenchmarkMockQuery) SetStopChannel(stopChannel misc.StopChannel) {
	this.stopChannel = stopChannel
}

func (this *BenchmarkMockQuery) StartTime() time.Time {
	return this.startTime
}

func BenchmarkMock(b *testing.B) {
	qc := test.Start("mock:items=10000", "p0")
	defer close(qc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runBenchmarkMock(qc, `select * from b0 where id = "1234"`)
	}
}

type MockBenchmarkResponse struct {
	err      query.Error
	results  int
	warnings []query.Error
	done     chan bool
}

func (this *MockBenchmarkResponse) SendError(err query.Error) {
	this.err = err
	if err.IsFatal() {
		close(this.done)
	}
}

func (this *MockBenchmarkResponse) SendResult(val interface{}) {
	this.results++
}

func (this *MockBenchmarkResponse) NoMoreResults() {
	close(this.done)
}

func runBenchmarkMock(qc network.QueryChannel, q string) (int, []query.Error, query.Error) {
	mr := &MockBenchmarkResponse{warnings: []query.Error{}, done: make(chan bool)}
	query := BenchmarkMockQuery{
		request:   network.StringQueryRequest{QueryString: q},
		response:  mr,
		startTime: time.Now(),
	}
	qc <- &query
	<-mr.done
	return mr.results, mr.warnings, mr.err
}
