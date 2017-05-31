package main

import (
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func runContainer(j Job) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, j.ImageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: j.ImageName,
		Cmd:   j.RunCommand,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
	log.Println("Started with Container ID:", resp.ID)
}
