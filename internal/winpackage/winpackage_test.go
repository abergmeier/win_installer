package winpackage

import "testing"

func TestMissingArguments(t *testing.T) {

	args, err := argumentsFromConfig(make(map[string]interface{}))
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	if len(args) != 0 {
		t.Fatalf("Expected empty slice: %#v", args)
	}
	args, err = argumentsFromConfig(map[string]interface{}{
		"arguments": nil,
	})
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if len(args) != 0 {
		t.Fatalf("Expected empty slice: %#v", args)
	}
}

func TestMissingProduct(t *testing.T) {
	in := isProductIDInstalled(make(map[string]interface{}))
	if in {
		t.Fatal("Expected installed to be false")
	}
}
