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

## Openfaas

Deploy the `maas-faas` function using `maas-faas.yml`:

```
faas-cli up -f ./maas-faas.yml
```

To start a build use call the maas-faas endpoint and pass the git repo URL via the `giturl` query parameter. This will return a container ID where your build job will run. You can specify which make targets to build using the `makecmd` query parameter.

```
$ curl http://192.168.99.100:31112/function/maas-faas?giturl=https://github.com/ahmedalhulaibi/maas.git&makecmd=install-tools&makecmd=build

db8e5686f368a58e08e4376d261c03bd758618c30b327e15d1c2daf1f8991928
```
To query the status of your build job call the same endpoint with the `container` query parameter
```
$ curl http://192.168.99.100:31112/function/maas-faas?container=db8e5686f368a58e08e4376d261c03bd758618c30b327e15d1c2daf1f8991928

Container started at: 2019-07-23T20:11:38.156508058Z
git URL: https://github.com/ahmedalhulaibi/maas.git
Make args: install-tools build
Cloning into 'gitmaas'...
...
Container finished at: 2019-07-23T20:13:55.553940387Z
Container exit code:  0
```