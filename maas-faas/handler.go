package function

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/docker/docker/pkg/stdcopy"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		handleErr(http.StatusInternalServerError, err.Error(), w)
		return
	}

	statusReq := r.URL.Query().Get("container")

	if statusReq != "" {
		output, err := GetContainerStatus(ctx, statusReq, cli)
		if err != nil {
			handleErr(http.StatusInternalServerError, err.Error(), w)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
		//end of status request, do not continue
		return
	}

	gitURL := r.URL.Query().Get("giturl")

	if gitURL == "" {
		handleErr(http.StatusBadRequest, "Missing 'giturl' query parameter", w)
		return
	}

	makeCmds := []string{"maas.sh", gitURL}
	makeCmds = append(makeCmds, r.URL.Query()["makecmd"]...)

	if output, err := ScheduleContainer(ctx, cli, gitURL, makeCmds); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(output))
	} else {
		handleErr(http.StatusBadRequest, err.Error(), w)
	}
}

func GetContainerStatus(ctx context.Context, containerID string, cli *client.Client) ([]byte, error) {

	outBuff := new(bytes.Buffer)
	containerJSON, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return outBuff.Bytes(), err
	}

	outBuff.WriteString(fmt.Sprintf("Container started at: %s", containerJSON.State.StartedAt))

	out, err := cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err == nil {
		_, err = stdcopy.StdCopy(outBuff, outBuff, out)
	}
	if !containerJSON.State.Running && !containerJSON.State.Paused {
		outBuff.WriteString(fmt.Sprintf("Container finished at: %s", containerJSON.State.FinishedAt))
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

func handleErr(status int, message string, w http.ResponseWriter) {
	io.Copy(os.Stderr, strings.NewReader(fmt.Sprintln(status, message)))
	w.WriteHeader(status)
	w.Write([]byte(message))
}
