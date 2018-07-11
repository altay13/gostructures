package strings

import "testing"

func TestGetAllSubstrings(t *testing.T) {
	dummy := "abcd"
	res := GetAllSubstrings(dummy)
	if len(res) != 10 {
		t.Fatalf("Failed testing GetAllSubstrings. Expected 10, instead %d", len(res))
	}
	if res[0] != "a" {
		t.Fatalf("Failed testing GetAllSubsctrings. Expected a, instead %s", res[0])
	}
	if res[len(res)-1] != "d" {
		t.Fatalf("Failed testing GetAllSubsctrings. Expected d, instead %s", res[len(res)-1])
	}
}
