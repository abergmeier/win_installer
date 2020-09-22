package gcstorage

import "testing"

func TestBucketConfig(t *testing.T) {
	err := Run(make(map[string]interface{}))
	if err == nil {
		t.Fatal("Expected error when not providing bucket")
	}
	if err.Error() != "Missing bucket config" {
		t.Fatalf("Expected other bucket key error: %s", err)
	}
	err = Run(map[string]interface{}{
		"bucket": nil,
	})
	if err == nil {
		t.Fatal("Expected error when not providing bucket value")
	}
	if err.Error() != "Missing value in bucket config" {
		t.Fatalf("Expected other error when not providing bucket value: %s", err)
	}
}
