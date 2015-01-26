package jq

import "testing"

func TestNewJqValid(t *testing.T) {
	jq, err := NewJq(".")

	if err != nil {
		t.Errorf("expected error to be nil, got: %s", err)
	}

	if jq == nil {
		t.Error("expected jq to not be nil")
	}
	if jq.state == nil {
		t.Error("expected jq.state to not be nil")
	}

}

func TestNewJqInvalid(t *testing.T) {

	jq, err := NewJq("INVALID")
	if err != errInvalidProgram {
		t.Errorf("expected error to be %s, got: %s", errInvalidProgram, err)
	}

	if jq != nil {
		t.Error("expected jq to be nil")
	}
}
