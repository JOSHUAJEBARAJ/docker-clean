package dock

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Image struct {
	Id   string
	Name string
}

func GetImages() []Image {

	var images []Image

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	imageslist, err := cli.ImageList(context.Background(), types.ImageListOptions{})

	if err != nil {
		panic(err)
	}

	for _, image := range imageslist {

		var img Image
		img.Id = image.ID

		name := strings.Join(image.RepoTags, " ")
		img.Name = name
		images = append(images, img)

	}

	return images
}

func (img *Image) Delete(id string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	_, err2 := cli.ImageRemove(context.Background(), id, types.ImageRemoveOptions{Force: true})
	if err2 != nil {
		fmt.Println("Error")
	}
}
