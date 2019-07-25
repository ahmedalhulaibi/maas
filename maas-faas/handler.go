package function

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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
	//if container query param present, get container status + logs
	if statusReq != "" {
		output, err := ContainerStatus(ctx, statusReq, cli)
		if err != nil {
			handleErr(http.StatusInternalServerError, err.Error(), w)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(output)
		//end of status request, do not continue
		return
	}

	//start a build job
	gitURL := r.URL.Query().Get("giturl")

	if gitURL == "" {
		handleErr(http.StatusBadRequest, "Missing 'giturl' query parameter", w)
		return
	}

	makeCmds := []string{"maas.sh", gitURL}
	makeCmds = append(makeCmds, r.URL.Query()["makecmd"]...)

	//start container and write container ID (SHA) to response body
	if containerID, err := ScheduleContainer(ctx, cli, gitURL, makeCmds); err == nil {
		w.WriteHeader(http.StatusOK)
		//
		w.Write([]byte(containerID))
	} else {
		handleErr(http.StatusBadRequest, err.Error(), w)
	}
}

func handleErr(status int, message string, w http.ResponseWriter) {
	io.Copy(os.Stderr, strings.NewReader(fmt.Sprintln(status, message)))
	w.WriteHeader(status)
	w.Write([]byte(message))
}
