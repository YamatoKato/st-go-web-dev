package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("New() err: %v", err)
	}

	if got.Port != wantPort {
		t.Errorf("New() = %v; want %v", got.Port, wantPort)
	}

	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("New() = %v; want %v", got.Env, wantEnv)
	}
}