package couchbase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	//	"github.com/couchbaselabs/clog"
	//	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

type RequestType int

const (
	CREATE RequestType = iota
	DELETE RequestType = iota
	LIST   RequestType = iota
	QUERY  RequestType = iota
)

type IndexRequest struct {
	Type   RequestType
	Index  IndexInfo
	Params QueryParams
}

type IndexInfo struct {
	Name       string
	Id         string
	Using      catalog.IndexType
	CreateStmt string
	Bucket     string
	IsPrimary  bool
}

type QueryParams struct {
	Low       string
	High      string
	Inclusion catalog.RangeInclusion
	Limit     int64
}

type IndexRow struct {
	Key   interface{}
	Value interface{}
}

type IndexError struct {
	From   string
	Reason string
}

func (ie IndexError) Error() string {
	return fmt.Sprintf("Node: %v, reason: %v", ie.From, ie.Reason)
}

type ResponseStatus int

const (
	SUCCESS ResponseStatus = iota
	ERROR   ResponseStatus = iota
)

type IndexResultResponse struct {
	Status ResponseStatus
	Rows   []IndexRow
	Errors []IndexError
}

type IndexMetaResponse struct {
	Status  ResponseStatus
	Errors  []IndexError
	Indexes []IndexInfo
}

var HttpClient = http.DefaultClient

func (li *lsmIndex) DropLsmIndex() query.Error {

	var ir IndexRequest

	ir.Type = DELETE
	ir.Index.Bucket = li.BucketId()
	ir.Index.Name = li.Name()
	ir.Index.On = li.Key().String()
	ir.Index.Using = catalog.LSM
	ir.Index.Id = li.Id()
	ir.Index.IsPrimary = li.IsPrimary()

	j, err := json.Marshal(ir)
	if err != nil {
		return query.NewError(err, fmt.Sprintf("Error Marshalling Index Data for %s", li.Name()))
	}

	u := "http://127.0.0.1:8080/index"
	req, err := http.NewRequest("PUT", u, bytes.NewReader(j))
	if err != nil {
		return query.NewError(err, fmt.Sprintf("Index Server Error Dropping Index %s", li.Name()))
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := HttpClient.Do(req)
	if err != nil {
		return query.NewError(err, fmt.Sprintf("Index Server Error Dropping Index %s", li.Name()))
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		body, _ := ioutil.ReadAll(res.Body)
		return query.NewError(nil, fmt.Sprintf("Error Dropping Index: %v / %s", res.Status, body))
	}
	return nil
}

func loadLsmIndexes(b *bucket) ([]*catalog.Index, error) {

	u := "http://127.0.0.1:8080/index"
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	res, err := HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error starting view req at %v: %v", u, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bod := make([]byte, 512)
		l, _ := res.Body.Read(bod)
		return nil, fmt.Errorf("Error executing view req at %v: %v - %s",
			u, res.Status, bod[:l])
	}

	var m IndexMetaResponse
	d := json.NewDecoder(res.Body)
	if err := d.Decode(&m); err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}

	indexes := make([]*catalog.Index, 0, len(m.Indexes))
	for _, i := range m.Indexes {

		var index catalog.Index

		if i.IsPrimary {
			index = &primaryLsmIndex{
				lsmIndex{
					name:   i.Name,
					id:     i.Id,
					bucket: b,
					using:  catalog.LSM,
				},
			}
			indexes = append(indexes, &index)
		} else {
			expr, err := ast.UnmarshalExpression([]byte(i.On))
			if err != nil {
				return nil, errors.New("Cannot unmarshal expression for index " + i.Name)
			}
			index = &lsmIndex{
				name:   i.Name,
				id:     i.Id,
				bucket: b,
				using:  catalog.LSM,
				on:     expr,
			}
			indexes = append(indexes, &index)
		}

	}
	return indexes, nil

}

func newLsmIndex(name string, on catalog.IndexKey, bkt *bucket) (*lsmIndex, error) {

	var ir IndexRequest

	ir.Type = CREATE
	ir.Index.Bucket = bkt.name
	ir.Index.Name = name
	ir.Index.On = on.String()
	ir.Index.Using = catalog.LSM
	ir.Index.IsPrimary = false

	j, err := json.Marshal(ir)
	if err != nil {
		return nil, err
	}

	u := "http://127.0.0.1:8080/index"
	req, err := http.NewRequest("PUT", u, bytes.NewReader(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("Error installing view: %v / %s",
			res.Status, body)
	}

	inst := lsmIndex{
		name:   name,
		using:  catalog.LSM,
		on:     on,
		bucket: bkt,
	}

	return &inst, nil
}
