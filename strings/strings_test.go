package strings

import (
	"testing"
)

func TestContainsString(t *testing.T) {
	list := []string{"apple", "banana", "orange"}
	target := "banana"
	if !ContainsString(list, target) {
		t.Errorf("Expected %s to be in the list", target)
	}

	target = "grape"
	if ContainsString(list, target) {
		t.Errorf("Expected %s not to be in the list", target)
	}
}

func TestHasPrefix(t *testing.T) {
	list := []string{"apple", "banana", "orange"}
	prefix := "app"
	if !HasPrefix(list, prefix) {
		t.Errorf("Expected at least one string with prefix %s in the list", prefix)
	}

	prefix = "gr"
	if HasPrefix(list, prefix) {
		t.Errorf("Expected no strings with prefix %s in the list", prefix)
	}
}

func TestStrArrayPtrRemoveEmpty(t *testing.T) {
	stringList := &[]string{"", "apple", "", "banana", "orange", ""}
	expected := &[]string{"apple", "banana", "orange"}

	result := StrArrayPtrRemoveEmpty(stringList)
	if !isEqual(result, expected) {
		t.Errorf("Expected %v, got %v", *expected, *result)
	}
}

func isEqual(a, b *[]string) bool {
	if len(*a) != len(*b) {
		return false
	}
	for i, v := range *a {
		if v != (*b)[i] {
			return false
		}
	}
	return true
}

func TestBoolString(t *testing.T) {
	if BoolString(true) != "true" {
		t.Errorf("Expected true, got %s", BoolString(true))
	}

	if BoolString(false) != "false" {
		t.Errorf("Expected false, got %s", BoolString(false))
	}
}

func TestBoolStatus(t *testing.T) {
	if BoolStatus(true) != "ok" {
		t.Errorf("Expected ok, got %s", BoolStatus(true))
	}

	if BoolStatus(false) != "fail" {
		t.Errorf("Expected fail, got %s", BoolStatus(false))
	}
}

func TestBoolYesNo(t *testing.T) {
	if BoolYesNo(true) != "yes" {
		t.Errorf("Expected yes, got %s", BoolYesNo(true))
	}

	if BoolYesNo(false) != "no" {
		t.Errorf("Expected no, got %s", BoolYesNo(false))
	}
}
