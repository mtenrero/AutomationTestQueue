package dockerMiddleware

import (
	"context"
	"errors"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// getClient returns the Docker Client
func getClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.35"))
	if err != nil {
		panic(err)
	}

	return cli
}

// listLocalImages returns an array containing all images locally available
func listLocalImages() ([]types.ImageSummary, error) {
	client := getClient()

	images, err := client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		logger.WithField("query", "ImageList").Warn(err)
		return nil, err
	}

	logger.WithField("query", "ImageList").Info("Retrieved local images")
	return images, nil
}

// imageLocallyAvalilable returns the specified image Summary if it's available in the system, otherwise, it will return an error
func imageLocallyAvalilable(imageName string) (*types.ImageSummary, error) {
	client := getClient()

	images, err := client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		for _, tag := range image.RepoTags {
			if strings.Contains(tag, imageName) {
				logger.WithField("image", imageName).Info("Image Found Locally")
				return &image, nil
			}
		}
	}

	return nil, errors.New("The image " + imageName + " is not available in the system!")
}

// pullImage pulls the specified image from Docker Registry or updates it if it's already downloaded
func pullImage(imageName string) error {
	client := getClient()

	out, err := client.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		logger.WithField("image", imageName).Warning("Image Pull failed")
		return err
	}

	logger.WithField("image", imageName).Info("Image succesfully pulled")
	defer out.Close()

	return nil
}
