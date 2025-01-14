package images

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestValidRequestForActiveImages(t *testing.T) {
	params := SearchParams{
		Cloud:  "amazon",
		Region: "eu-central-1",
		State:  "active",
	}
	expectedRequest := fmt.Sprintf(
		"/%s/%s/images/%s.json",
		params.Cloud,
		params.Region,
		params.State)

	// We setup a fake http server that mocks a registration server.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != expectedRequest {
			t.Fatalf("Unexpected request. Got %s, expected %s", r.RequestURI, expectedRequest)
		}

		file, err := os.Open("testdata/active.json")
		if err != nil {
			fmt.Fprintln(w, "FAIL!")
			return
		}
		defer file.Close()
		io.Copy(w, file)

	}))
	defer ts.Close()
	params.ApiEndpoint = ts.URL

	images, err := GetImages(params)
	if err != nil {
		t.Fatal("It should've run just fine...")
	}
	if len(images) != 22 {
		t.Fatalf("Unexpected number of images found. Got %d, expected %d", len(images), 22)
	}
}

func TestFilterImages(t *testing.T) {
	params := SearchParams{
		Cloud:     "amazon",
		Region:    "eu-central-1",
		State:     "active",
		NameRegex: "suse-sles-15-sp1-byos.*-hvm-ssd-x86_64",
	}
	expectedRequest := fmt.Sprintf(
		"/%s/%s/images/%s.json",
		params.Cloud,
		params.Region,
		params.State)

	// We setup a fake http server that mocks a registration server.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != expectedRequest {
			t.Fatalf("Unexpected request. Got %s, expected %s", r.RequestURI, expectedRequest)
		}

		file, err := os.Open("testdata/active.json")
		if err != nil {
			fmt.Fprintln(w, "FAIL!")
			return
		}
		defer file.Close()
		io.Copy(w, file)

	}))
	defer ts.Close()
	params.ApiEndpoint = ts.URL

	images, err := GetImages(params)
	if err != nil {
		t.Fatal("It should've run just fine...")
	}
	if len(images) != 1 {
		t.Fatalf("Unexpected number of images found. Got %d, expected %d", len(images), 1)
	}
}

func TestSortAscendingImages(t *testing.T) {
	params := SearchParams{
		Cloud:         "amazon",
		Region:        "eu-central-1",
		State:         "active",
		SortAscending: true,
		NameRegex:     "suse-sles-.*-sapcal.*-hvm-ssd-x86_64",
	}
	expectedRequest := fmt.Sprintf(
		"/%s/%s/images/%s.json",
		params.Cloud,
		params.Region,
		params.State)

	// We setup a fake http server that mocks a registration server.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != expectedRequest {
			t.Fatalf("Unexpected request. Got %s, expected %s", r.RequestURI, expectedRequest)
		}

		file, err := os.Open("testdata/active.json")
		if err != nil {
			fmt.Fprintln(w, "FAIL!")
			return
		}
		defer file.Close()
		io.Copy(w, file)

	}))
	defer ts.Close()
	params.ApiEndpoint = ts.URL

	images, err := GetImages(params)
	if err != nil {
		t.Fatal("It should've run just fine...")
	}
	if len(images) != 3 {
		t.Fatalf("Unexpected number of images found. Got %d, expected %d", len(images), 1)
	}

	expectedIDs := []string{
		"ami-082bfb28e7de47e17",
		"ami-057b6b1654d10ff7b",
		"ami-07dd6bca2aa25c67d",
	}
	for pos, image := range images {
		if image.Id != expectedIDs[pos] {
			t.Fatalf("Sorting error for image at position %d", pos)
		}
	}
}

func TestSortDescendingImages(t *testing.T) {
	// descending order is the default one
	params := SearchParams{
		Cloud:     "amazon",
		Region:    "eu-central-1",
		State:     "active",
		NameRegex: "suse-sles-.*-sapcal.*-hvm-ssd-x86_64",
	}
	expectedRequest := fmt.Sprintf(
		"/%s/%s/images/%s.json",
		params.Cloud,
		params.Region,
		params.State)

	// We setup a fake http server that mocks a registration server.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != expectedRequest {
			t.Fatalf("Unexpected request. Got %s, expected %s", r.RequestURI, expectedRequest)
		}

		file, err := os.Open("testdata/active.json")
		if err != nil {
			fmt.Fprintln(w, "FAIL!")
			return
		}
		defer file.Close()
		io.Copy(w, file)

	}))
	defer ts.Close()
	params.ApiEndpoint = ts.URL

	images, err := GetImages(params)
	if err != nil {
		t.Fatal("It should've run just fine...")
	}
	if len(images) != 3 {
		t.Fatalf("Unexpected number of images found. Got %d, expected %d", len(images), 1)
	}

	expectedIDs := []string{
		"ami-057b6b1654d10ff7b",
		"ami-07dd6bca2aa25c67d",
		"ami-082bfb28e7de47e17",
	}
	for pos, image := range images {
		if image.Id != expectedIDs[pos] {
			t.Fatalf("Sorting error for image at position %d", pos)
		}
	}
}
