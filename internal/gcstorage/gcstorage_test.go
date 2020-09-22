package gcstorage

import "testing"

func TestBucketConfig(t *testing.T) {
	err := Run(make(map[string]interface{}))
	if err == nil {
		t.Fatal("Expected error when not providing bucket")
	}
}
