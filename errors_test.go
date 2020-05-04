package yael_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/snicol/yael"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUnmarshal(t *testing.T) {
	src := yael.New("test_error").WithMeta("foo", "bar").WithReason(yael.New("inner_error"))

	o, err := json.Marshal(src)
	assert.Nil(t, err)
	assert.Equal(t, `{"code":"test_error","reason":{"code":"inner_error"},"meta":{"foo":"bar"}}`, string(o))

	var dst *yael.E

	assert.Nil(t, json.Unmarshal(o, &dst))

	assert.Equal(t, src, dst)
}

func TestWrap(t *testing.T) {
	e2 := yael.New("another_error")
	e1 := yael.New("test_error").WithReason(e2)

	assert.Equal(t, true, errors.Is(e1, e2))
	assert.Equal(t, false, errors.Is(e2, e1))

	e3 := yael.New("parent_error").WithReason(e1)
	assert.Equal(t, true, errors.Is(e3, e2))
}

func TestError(t *testing.T) {
	e := yael.New("some_error")

	var i interface{} = e
	_, ok := i.(error)
	assert.True(t, ok)
}
