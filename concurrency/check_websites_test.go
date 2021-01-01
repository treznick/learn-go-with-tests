package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "lol://notawebsite.lol" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blob.gypsydave5.com",
		"lol://notawebsite.lol",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blob.gypsydave5.com": true,
		"lol://notawebsite.lol":      false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
