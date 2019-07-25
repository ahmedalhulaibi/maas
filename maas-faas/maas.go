package function

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func ContainerStatus(ctx context.Context, containerID string, cli *client.Client) ([]byte, error) {

	outBuff := new(bytes.Buffer)
	containerJSON, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return outBuff.Bytes(), err
	}

	outBuff.WriteString(fmt.Sprintf("Container started at: %s\n", containerJSON.State.StartedAt))

	out, err := cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err == nil {
		_, err = stdcopy.StdCopy(outBuff, outBuff, out)
	}
	if !containerJSON.State.Running && !containerJSON.State.Paused {
		outBuff.WriteString(fmt.Sprintf("Container finished at: %s\n", containerJSON.State.FinishedAt))
		outBuff.WriteString(fmt.Sprintln("Container exit code: ", containerJSON.State.ExitCode))
	}
	return outBuff.Bytes(), err
}

func ScheduleContainer(ctx context.Context, cli *client.Client, gitURL string, makeCmds []string) (string, error) {

	reader, err := cli.ImagePull(ctx, "ahmedalhulaibi/maas:latest", types.ImagePullOptions{})
	if err != nil {
		//handleErr(http.StatusInternalServerError, err.Error(), w)
		return "", err
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:      "ahmedalhulaibi/maas:latest",
		Entrypoint: makeCmds,
		Tty:        false,
		Labels: map[string]string{
			"maas": "",
		},
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: "/var/run/docker.sock",
				Target: "/var/run/docker.sock",
			},
		},
	}, nil, "")

	if err != nil {
		//handleErr(http.StatusInternalServerError, err.Error(), w)
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		//handleErr(http.StatusInternalServerError, err.Error(), w)
		return "", err
	}

	return resp.ID, err
}
