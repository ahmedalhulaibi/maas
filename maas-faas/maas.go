package function

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/docker/docker/api/types/filters"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

/*ContainerStatus returns the stdout & stderr in a byte slice for the given container ID*/
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

/*ScheduleContainer starts a maas container on the docker host the given git URL and make targets
and returns the corresponding container ID*/
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
			"maas":          "",
			"maas.gitURL":   gitURL,
			"maas.makecmds": strings.Join(makeCmds[2:], ","),
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

/*ContainerStatusRecord wraps high-level container status info*/
type ContainerStatusRecord struct {
	ID         string
	StartedAt  string
	FinishedAt string
	GitURL     string
	Targets    []string
	RC         int
}

/*AllContainers returns a list of ContainerStatusRecord*/
func AllContainers(ctx context.Context, cli *client.Client) ([]ContainerStatusRecord, error) {
	filterOpts := filters.NewArgs()
	filterOpts.Add("label", "maas")

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{Filters: filterOpts, All: true})
	containerStatusRecs := []ContainerStatusRecord{}
	for _, container := range containers {

		containerJSON, err := cli.ContainerInspect(ctx, container.ID)
		if err != nil {
			return nil, err
		}
		newRec := ContainerStatusRecord{
			ID:         container.ID,
			StartedAt:  containerJSON.State.StartedAt,
			FinishedAt: containerJSON.State.FinishedAt,
			RC:         containerJSON.State.ExitCode,
			GitURL:     containerJSON.Config.Labels["maas.gitURL"],
			Targets:    strings.Split(containerJSON.Config.Labels["maas.makecmds"], ","),
		}
		containerStatusRecs = append(containerStatusRecs, newRec)
	}

	sort.Slice(containerStatusRecs, func(i, j int) bool {
		return containerStatusRecs[i].StartedAt > containerStatusRecs[j].StartedAt
	})
	return containerStatusRecs, err
}
