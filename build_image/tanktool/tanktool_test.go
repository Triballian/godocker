package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestIndexHanlder(t *testing.T) {
	assert := assert.New(t)
	// 	type HTTP.Params struct{}
	// var Param p
	res := httprouter.Params{
		httprouter.Param{"param1", "value1"},
		httprouter.Param{"param2", "value2"},
		httprouter.Param{"param3", "value3"},
	}
	for i := range res {
		if val := res.ByName(res[i].Key); val != res[i].Value {
			t.Errorf("Wrong value for %s: Got %s; Want %s", res[i].Key, val, res[i].Value)
		}
	}
	if val := res.ByName("noKey"); val != "" {
		t.Errorf("Expected empty string for not found key; got: %s", val)
	}
	rw := httptest.NewRecorder()

	r, _ := http.NewRequest("GET", "/", nil)
	Index(rw, r, res)
	// if rw.Code != 200 {
	// 	t.Fatal("Expected status to == 200, but got %d", rw.Code)
	assert.Equal(rw.Code, 200)
	assert.Contains(rw.Body.String(), "Tank")
}
