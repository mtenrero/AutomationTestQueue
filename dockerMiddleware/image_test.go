package dockerMiddleware

import (
	"testing"
)

func TestImageLocallyAvailable(t *testing.T) {
	image, err := ImageLocallyAvalilable("hello-world")

	if err != nil {
		t.Fatal(err)
	}

	if image.RepoTags[0] != "hello-world:latest" {
		t.Fail()
	}
}

func TestImageLocallyFail(t *testing.T) {
	image, err := ImageLocallyAvalilable("NOT_FOUND_IMAGE")

	if err == nil {
		t.Fatal("The image shouldn't be found")
	}

	if image != nil {
		t.Fatal("The image shouldn't be found")
	}
}

func TestPullImage(t *testing.T) {
	err := PullImage("alpine")

	if err != nil {
		t.Fatal(err)
	}
}

func TestPullImageFail(t *testing.T) {
	err := PullImage("alpineERROR")

	if err == nil {
		t.Fatal(err)
	}
}

func TestRetrieveImages(t *testing.T) {
	_, err := ListLocalImages()

	if err != nil {
		t.Fatal(err)
	}
}
