package function

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	gitURL := r.URL.Query().Get("giturl")

	if gitURL == "" {
		handleErr(http.StatusBadRequest, "Missing 'giturl' query parameter", w)
		return
	}

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		handleErr(http.StatusInternalServerError, err.Error(), w)
		return
	}

	reader, err := cli.ImagePull(ctx, "ahmedalhulaibi/maas:latest", types.ImagePullOptions{})
	if err != nil {
		handleErr(http.StatusInternalServerError, err.Error(), w)
		return
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:      "ahmedalhulaibi/maas:latest",
		Entrypoint: []string{"maas.sh", gitURL},
		Tty:        true,
		Volumes:    map[string]struct{}{"/var/run/docker.sock:/var/run/docker.sock": {}},
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
		handleErr(http.StatusInternalServerError, err.Error(), w)
		return
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		handleErr(http.StatusInternalServerError, err.Error(), w)
		return
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			handleErr(http.StatusInternalServerError, err.Error(), w)
			return
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		handleErr(http.StatusInternalServerError, err.Error(), w)
		return
	}

	output, err := ioutil.ReadAll(out)
	if err != nil {
		handleErr(http.StatusInternalServerError, err.Error(), w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func handleErr(status int, message string, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}
