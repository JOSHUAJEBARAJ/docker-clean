package dock

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Image struct {
	Id   string
	Name string
}

func Init() (*client.Client, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return cli, nil

}

func GetImages(cli *client.Client) ([]Image, error) {

	var images []Image

	imageslist, err := cli.ImageList(context.Background(), types.ImageListOptions{})

	if err != nil {
		return nil, err
	}

	for _, image := range imageslist {

		var img Image
		img.Id = image.ID

		name := strings.Join(image.RepoTags, " ")
		img.Name = name
		images = append(images, img)

	}

	return images, nil
}

func (img *Image) Delete(id string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = cli.ImageRemove(context.Background(), id, types.ImageRemoveOptions{Force: true})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
