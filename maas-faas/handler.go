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
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
		<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
		<script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
		<script>
		document.addEventListener('DOMContentLoaded', function() {
			var elems = document.querySelectorAll('.collapsible');
			var instances = M.Collapsible.init(elems);
		  });
		</script>

	</head>
	<body>
		<ul class="collapsible">
		{{range .}}
			<li>
				<div class="collapsible-header"><i class="material-icons">build</i>Container ID: {{ .ID}}</div>
				<div class="collapsible-body">
					<div>
						<div><strong>Started At:</strong> {{ .StartedAt}}</div>
						<div><strong>Finished At:</strong> {{ .FinishedAt}}</div>
						<div><strong>Exit Code:</strong> {{ .RC}}</div>
					</div>
				</div>
			  </li>
			  {{end}}
		</ul>

		{{range .}}
		<div class="row">
    		<div class="col s12 m6">
      			<div class="card blue-grey darken-1">
        			<div class="card-content white-text">
          				<span class="card-title"><strong>Container ID:</strong>{{ .ID}}</span>
						  <div><strong>Started At:</strong> {{ .StartedAt}}</div>
						  <div><strong>Finished At:</strong> {{ .FinishedAt}}</div>
						  <div><strong>Exit Code:</strong> {{ .RC}}</div>
        			</div>
        			<div class="card-action">
          				<a href="#">This is a link</a>
          				<a href="#">This is a link</a>
        			</div>
      			</div>
    		</div>
  		</div>{{end}}

    
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
