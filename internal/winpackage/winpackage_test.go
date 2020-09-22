package winpackage

import "testing"

func TestMissingProduct(t *testing.T) {
	in := isProductIDInstalled(make(map[string]interface{}))
	if in {
		t.Fatal("Expected installed to be false")
	}
}
