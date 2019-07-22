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

Test using curl command. Pass the git repo URL via the `giturl` query parameter.

```
http://192.168.99.100:31112/function/maas-faas?giturl=https://github.com/ahmedalhulaibi/maas.git
Cloning into 'gitmaas'...
remote: Enumerating objects: 94, done.
remote: Counting objects:   1% (1/94)   
...
remote: Counting objects:  97% (92/94)   
remote: Counting objects:  98% (93/94)   
remote: Counting objects: 100% (94/94)   
remote: Counting objects: 100% (94/94), done.
remote: Compressing objects:   1% (1/66)   
...   
remote: Compressing objects:  96% (64/66)   
remote: Compressing objects:  98% (65/66)   
remote: Compressing objects: 100% (66/66)   
remote: Compressing objects: 100% (66/66), done.
remote: Total 94 (delta 43), reused 75 (delta 26), pack-reused 0
Unpacking objects:   1% (1/94)
...
Unpacking objects:  97% (92/94)
Unpacking objects:  98% (93/94)
Unpacking objects: 100% (94/94)
Unpacking objects: 100% (94/94), done.
[maaslog]: Clean started
[maaslog]: sleep 1
[maaslog]: Alpine
[maaslog]: Clean complete
[maaslog]: Dependencies installation started
[maaslog]: Alpine
[maaslog]: Hello Alpine
[maaslog]: #@apk add --no-cache --update jq zip
[maaslog]: Dependencies installation complete
[maaslog]: Build started
[maaslog]: docker build . -t ahmedalhulaibi/maas:latest
[maaslog]: Sending build context to Docker daemon  219.1kB


[maaslog]: Step 1/6 : FROM docker:latest
[maaslog]:  ---> e1ee9bd2e980
[maaslog]: Step 2/6 : RUN apk add --no-cache --update make git docker bash
[maaslog]:  ---> Using cache
[maaslog]:  ---> 9ba8ddc8cbf3
[maaslog]: Step 3/6 : WORKDIR /
[maaslog]:  ---> Using cache
[maaslog]:  ---> ae24582cef0d
[maaslog]: Step 4/6 : COPY ./maas.sh /usr/local/bin/maas.sh
[maaslog]:  ---> Using cache
[maaslog]:  ---> 66a0660c6231
[maaslog]: Step 5/6 : RUN chmod +x /usr/local/bin/maas.sh
[maaslog]:  ---> Using cache
[maaslog]:  ---> 30bfccccadb3
[maaslog]: Step 6/6 : ENTRYPOINT [ "maas.sh" ]
[maaslog]:  ---> Using cache
[maaslog]:  ---> 13af4d55c078
[maaslog]: Successfully built 13af4d55c078
[maaslog]: Successfully tagged ahmedalhulaibi/maas:latest
[maaslog]: Build complete
```