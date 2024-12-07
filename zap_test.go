package gclog

import (
	"testing"
)

func TestNewProduction(t *testing.T) {
	lg, err := zapNewProduction("foo")
	if err != nil {
		t.Error(err)
	} else if lg == nil {
		t.Error("nil")
	}
}

func TestNewDevelopment(t *testing.T) {
	lg, err := zapNewDevelopment("foo")
	if err != nil {
		t.Error(err)
	} else if lg == nil {
		t.Error("nil")
	}
}

func TestNewLocal(t *testing.T) {
	lg, err := zapNewLocal("foo")
	if err != nil {
		t.Error(err)
	} else if lg == nil {
		t.Error("nil")
	}
}
