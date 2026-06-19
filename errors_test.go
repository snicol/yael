package yael_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/snicol/yael"
)

func TestMarshalUnmarshal(t *testing.T) {
	src := yael.New("test_error").WithMeta("foo", "bar").WithReasons(yael.New("inner_error"))

	o, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	want := `{"code":"test_error","reasons":[{"code":"inner_error"}],"meta":{"foo":"bar"}}`
	if string(o) != want {
		t.Errorf("got %s, want %s", o, want)
	}

	var dst *yael.E
	if err := json.Unmarshal(o, &dst); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if dst.Code != src.Code {
		t.Errorf("code: got %s, want %s", dst.Code, src.Code)
	}
}

func TestIs(t *testing.T) {
	sentinel := yael.New("not_found")
	other := yael.New("not_found")

	if !errors.Is(other, sentinel) {
		t.Error("expected errors with the same code to match")
	}

	wrapped := yael.New("wrapper").WithReasons(other)
	if !errors.Is(wrapped, sentinel) {
		t.Error("expected sentinel to match through wrapping")
	}
}

func TestWrap(t *testing.T) {
	e2 := yael.New("another_error")
	e1 := yael.New("test_error").WithReasons(e2)

	if !errors.Is(e1, e2) {
		t.Error("expected e1 to wrap e2")
	}

	if errors.Is(e2, e1) {
		t.Error("expected e2 not to wrap e1")
	}

	e3 := yael.New("parent_error").WithReasons(e1)
	if !errors.Is(e3, e2) {
		t.Error("expected e3 to transitively wrap e2")
	}
}

func TestWrapMultiple(t *testing.T) {
	e1 := yael.New("error_one")
	e2 := yael.New("error_two")
	e3 := yael.New("error_three")

	combined := yael.New("combined").WithReasons(e1, e2, e3)

	if !errors.Is(combined, e1) {
		t.Error("expected combined to contain e1")
	}

	if !errors.Is(combined, e2) {
		t.Error("expected combined to contain e2")
	}

	if !errors.Is(combined, e3) {
		t.Error("expected combined to contain e3")
	}
}

func TestError(t *testing.T) {
	e := yael.New("some_error")

	var i any = e
	if _, ok := i.(error); !ok {
		t.Error("expected E to implement error interface")
	}
}
