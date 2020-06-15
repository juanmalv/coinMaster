package webService

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/juanmalv/coinMaster/api/src/api/utils"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Ping(c)

	result, _:= utils.ReadBody(w.Result())

	var got gin.H
	err := json.Unmarshal(result, &got)
	if err != nil {
		t.Fatal(err)
	}

	expected := gin.H{"message":"pong"}

	assert.Equal(t, expected, got)
}

