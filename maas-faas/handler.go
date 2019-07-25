package function

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"html/template"

	"github.com/docker/docker/client"
)

var (
	allContainersTpl *template.Template
)

func init() {
	allContainersPage := ` 
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Maas Jobs</title>
	</head>
	<body>
		{{range .}}<div>{{ .ID}} {{ .StartedAt}} {{ .FinishedAt}} {{ .RC}}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`
	var err error
	if allContainersTpl, err = template.New("webpage").Parse(allContainersPage); err != nil {
		log.Fatal(err)
	}
}

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

	if gitURL != "" {
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
		return
	}

	if allContainers, err := AllContainers(ctx, cli); err == nil {
		if errTpl := allContainersTpl.Execute(w, allContainers); errTpl != nil {
			handleErr(http.StatusBadRequest, errTpl.Error(), w)
		}
	} else {
		handleErr(http.StatusBadRequest, err.Error(), w)
	}
}

func handleErr(status int, message string, w http.ResponseWriter) {
	io.Copy(os.Stderr, strings.NewReader(fmt.Sprintln(status, message)))
	w.WriteHeader(status)
	w.Write([]byte(message))
}
