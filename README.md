# maas - Make As A Service

Run make tasks on request in a docker container.

## How it works

Execute the container as a program and a script will `git clone` the given repo and run the `make` command.

## Quickstart

Run the container `ahmedalhulaibi/maas:latest`. 

The example below will clone this repo and run the `make` command. It also shows how to expose the host machine's Docker socket to run Docker commands (starting/stopping containers, build, push, pull) which is required for this project.

```
$ docker run -v /var/run/docker.sock:/var/run/docker.sock -it ahmedalhulaibi/maas:latest https://github.com/ahmedalhulaibi/maas.git
```

## Slower start

### Create a Make

```
$ git clone https://github.com/ahmedalhulaibi/maas.git
$ cd maas
$ ls
Dockerfile maas.sh Makefile osname.sh README.md
```

Notice the Makefile is in the root of the repository `./maas/Makefile`. **This is mandatory**.


Modify the `Makefile` replacing my dockerhub username `ahmedalhulaibi` with your dockerhub repo.

```
$ sed "s/ahmedalhulaibi/YOUR_USERNAME_HERE/g" Makefile
```

If you run the above `sed` command you may also want to check `test` target in the `Makefile` as there is a git URL supplied there. You'll probably want to change this.

Run `make build test` which will build the Docker image for make-as-a-service

## Roadmap

- [ ] Openfaas template