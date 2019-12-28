package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

func main() {

	endpoint, err := createLocalClient("unix:///var/run/docker.sock", "1.39")
	if err != nil {
		panic("Cannot connect to unix:///var/run/docker.sock. Is mobi deamon running? Do you have at least moby with api version 1.39?")
	}

	var psCmd = &cobra.Command{
		Use:   "ps",
		Short: "list all running containers",
		Run: func(cmd *cobra.Command, args []string) {
			LsContainers(endpoint)
		},
	}

	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "run container",
		Run: func(cmd *cobra.Command, args []string) {
			containerName, err := cmd.Flags().GetString("name")
			if err != nil {
				panic("missing container name to run")
			}
			imageName, _ := cmd.Flags().GetString("image")
			if err != nil {
				panic("missing image name to run")
			}
			ContainerStart(endpoint, imageName, containerName)
		},
	}

	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "stop container with specific name",
		Run: func(cmd *cobra.Command, args []string) {
			containerName, err := cmd.Flags().GetString("name")
			if err != nil {
				panic("cannot recognize container name to stop")
			}
			ContainerStop(endpoint, containerName)
		},
	}

	var rootCmd = &cobra.Command{Use: "mobicli"}
	startCmd.PersistentFlags().String("image", "", "name image to run continainer from it")
	startCmd.PersistentFlags().String("name", "", "name container to run")
	stopCmd.PersistentFlags().String("name", "", "name container to stop")
	rootCmd.AddCommand(psCmd, startCmd, stopCmd)
	err = rootCmd.Execute()
	if err != nil {
		panic("fatal error")
	}
}

func createLocalClient(url, apiVer string) (*client.Client, error) {
	clientMap := make(map[string]string)
	return client.NewClient(
		url,
		apiVer,
		&http.Client{},
		clientMap,
	)
}

func LsContainers(cli *client.Client) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf(" %s   %s  %s  %v  %s\n", container.ID[:10], container.Image, container.State, container.Ports, container.Names)
	}
}

func ContainerStart(docker *client.Client, imageName string, containerName string) {
	resp, err := docker.ContainerCreate(context.Background(), &container.Config{
		Image: imageName,
		Tty:   true,
	}, nil, nil, containerName)
	if err != nil {
		panic(err)
	}

	if err := docker.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}

func ContainerStop(docker *client.Client, containerName string) {
	containers, err := docker.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(containers)
	for _, container := range containers {
		if container.Names[0][1:] == containerName {
			fmt.Print("Stopping container ", container.ID[:10], "... ")
			if err := docker.ContainerStop(context.Background(), container.ID, nil); err != nil {
				panic(err)
			}
			fmt.Println("Success")
		}

	}
}
