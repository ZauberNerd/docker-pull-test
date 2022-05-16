package main

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	ctx := context.Background()
	imageRef := "docker.io/library/node:12-buster-slim"
	pullOptions := types.ImagePullOptions{
		Platform: "linux/amd64",
	}
	reader, err := cli.ImagePull(ctx, imageRef, pullOptions)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		panic(err)
	}
}
