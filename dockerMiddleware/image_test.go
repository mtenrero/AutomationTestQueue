package dockerMiddleware

import (
	"testing"
)

func TestImageLocallyAvailable(t *testing.T) {
	image, err := imageLocallyAvalilable("hello-world")

	if err != nil {
		t.Fatal(err)
	}

	if image.RepoTags[0] != "hello-world:latest" {
		t.Fail()
	}
}

func TestImageLocallyFail(t *testing.T) {
	image, err := imageLocallyAvalilable("NOT_FOUND_IMAGE")

	if err == nil {
		t.Fatal("The image shouldn't be found")
	}

	if image != nil {
		t.Fatal("The image shouldn't be found")
	}
}

func TestPullImage(t *testing.T) {
	err := pullImage("alpine")

	if err != nil {
		t.Fatal(err)
	}
}

func TestPullImageFail(t *testing.T) {
	err := pullImage("alpineERROR")

	if err == nil {
		t.Fatal(err)
	}
}
