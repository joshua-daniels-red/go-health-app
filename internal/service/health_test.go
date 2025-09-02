package service

import (
	"reflect"
	"testing"
)

func TestHealthService_Check(t *testing.T) {
	hs := NewHealthService()

	got := hs.Check()
	want := map[string]string{"status": "ok"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("HealthService.Check() = %v, want %v", got, want)
	}
}
