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

Test using curl command:

```
http://192.168.99.100:31112/function/maas-faas?giturl=https://github.com/ahmedalhulaibi/maas.git
Cloning into 'gitmaas'...
remote: Enumerating objects: 94, done.
remote: Counting objects:   1% (1/94)   
remote: Counting objects:   2% (2/94)   
remote: Counting objects:   3% (3/94)   
remote: Counting objects:   4% (4/94)   
remote: Counting objects:   5% (5/94)   
remote: Counting objects:   6% (6/94)   
remote: Counting objects:   7% (7/94)   
remote: Counting objects:   8% (8/94)   
remote: Counting objects:   9% (9/94)   
remote: Counting objects:  10% (10/94)   
remote: Counting objects:  11% (11/94)   
remote: Counting objects:  12% (12/94)   
remote: Counting objects:  13% (13/94)   
remote: Counting objects:  14% (14/94)   
remote: Counting objects:  15% (15/94)   
remote: Counting objects:  17% (16/94)   
remote: Counting objects:  18% (17/94)   
remote: Counting objects:  19% (18/94)   
remote: Counting objects:  20% (19/94)   
remote: Counting objects:  21% (20/94)   
remote: Counting objects:  22% (21/94)   
remote: Counting objects:  23% (22/94)   
remote: Counting objects:  24% (23/94)   
remote: Counting objects:  25% (24/94)   
remote: Counting objects:  26% (25/94)   
remote: Counting objects:  27% (26/94)   
remote: Counting objects:  28% (27/94)   
remote: Counting objects:  29% (28/94)   
remote: Counting objects:  30% (29/94)   
remote: Counting objects:  31% (30/94)   
remote: Counting objects:  32% (31/94)   
remote: Counting objects:  34% (32/94)   
remote: Counting objects:  35% (33/94)   
remote: Counting objects:  36% (34/94)   
remote: Counting objects:  37% (35/94)   
remote: Counting objects:  38% (36/94)   
remote: Counting objects:  39% (37/94)   
remote: Counting objects:  40% (38/94)   
remote: Counting objects:  41% (39/94)   
remote: Counting objects:  42% (40/94)   
remote: Counting objects:  43% (41/94)   
remote: Counting objects:  44% (42/94)   
remote: Counting objects:  45% (43/94)   
remote: Counting objects:  46% (44/94)   
remote: Counting objects:  47% (45/94)   
remote: Counting objects:  48% (46/94)   
remote: Counting objects:  50% (47/94)   
remote: Counting objects:  51% (48/94)   
remote: Counting objects:  52% (49/94)   
remote: Counting objects:  53% (50/94)   
remote: Counting objects:  54% (51/94)   
remote: Counting objects:  55% (52/94)   
remote: Counting objects:  56% (53/94)   
remote: Counting objects:  57% (54/94)   
remote: Counting objects:  58% (55/94)   
remote: Counting objects:  59% (56/94)   
remote: Counting objects:  60% (57/94)   
remote: Counting objects:  61% (58/94)   
remote: Counting objects:  62% (59/94)   
remote: Counting objects:  63% (60/94)   
remote: Counting objects:  64% (61/94)   
remote: Counting objects:  65% (62/94)   
remote: Counting objects:  67% (63/94)   
remote: Counting objects:  68% (64/94)   
remote: Counting objects:  69% (65/94)   
remote: Counting objects:  70% (66/94)   
remote: Counting objects:  71% (67/94)   
remote: Counting objects:  72% (68/94)   
remote: Counting objects:  73% (69/94)   
remote: Counting objects:  74% (70/94)   
remote: Counting objects:  75% (71/94)   
remote: Counting objects:  76% (72/94)   
remote: Counting objects:  77% (73/94)   
remote: Counting objects:  78% (74/94)   
remote: Counting objects:  79% (75/94)   
remote: Counting objects:  80% (76/94)   
remote: Counting objects:  81% (77/94)   
remote: Counting objects:  82% (78/94)   
remote: Counting objects:  84% (79/94)   
remote: Counting objects:  85% (80/94)   
remote: Counting objects:  86% (81/94)   
remote: Counting objects:  87% (82/94)   
remote: Counting objects:  88% (83/94)   
remote: Counting objects:  89% (84/94)   
remote: Counting objects:  90% (85/94)   
remote: Counting objects:  91% (86/94)   
remote: Counting objects:  92% (87/94)   
remote: Counting objects:  93% (88/94)   
remote: Counting objects:  94% (89/94)   
remote: Counting objects:  95% (90/94)   
remote: Counting objects:  96% (91/94)   
remote: Counting objects:  97% (92/94)   
remote: Counting objects:  98% (93/94)   
remote: Counting objects: 100% (94/94)   
remote: Counting objects: 100% (94/94), done.
remote: Compressing objects:   1% (1/66)   
remote: Compressing objects:   3% (2/66)   
remote: Compressing objects:   4% (3/66)   
remote: Compressing objects:   6% (4/66)   
remote: Compressing objects:   7% (5/66)   
remote: Compressing objects:   9% (6/66)   
remote: Compressing objects:  10% (7/66)   
remote: Compressing objects:  12% (8/66)   
remote: Compressing objects:  13% (9/66)   
remote: Compressing objects:  15% (10/66)   
remote: Compressing objects:  16% (11/66)   
remote: Compressing objects:  18% (12/66)   
remote: Compressing objects:  19% (13/66)   
remote: Compressing objects:  21% (14/66)   
remote: Compressing objects:  22% (15/66)   
remote: Compressing objects:  24% (16/66)   
remote: Compressing objects:  25% (17/66)   
remote: Compressing objects:  27% (18/66)   
remote: Compressing objects:  28% (19/66)   
remote: Compressing objects:  30% (20/66)   
remote: Compressing objects:  31% (21/66)   
remote: Compressing objects:  33% (22/66)   
remote: Compressing objects:  34% (23/66)   
remote: Compressing objects:  36% (24/66)   
remote: Compressing objects:  37% (25/66)   
remote: Compressing objects:  39% (26/66)   
remote: Compressing objects:  40% (27/66)   
remote: Compressing objects:  42% (28/66)   
remote: Compressing objects:  43% (29/66)   
remote: Compressing objects:  45% (30/66)   
remote: Compressing objects:  46% (31/66)   
remote: Compressing objects:  48% (32/66)   
remote: Compressing objects:  50% (33/66)   
remote: Compressing objects:  51% (34/66)   
remote: Compressing objects:  53% (35/66)   
remote: Compressing objects:  54% (36/66)   
remote: Compressing objects:  56% (37/66)   
remote: Compressing objects:  57% (38/66)   
remote: Compressing objects:  59% (39/66)   
remote: Compressing objects:  60% (40/66)   
remote: Compressing objects:  62% (41/66)   
remote: Compressing objects:  63% (42/66)   
remote: Compressing objects:  65% (43/66)   
remote: Compressing objects:  66% (44/66)   
remote: Compressing objects:  68% (45/66)   
remote: Compressing objects:  69% (46/66)   
remote: Compressing objects:  71% (47/66)   
remote: Compressing objects:  72% (48/66)   
remote: Compressing objects:  74% (49/66)   
remote: Compressing objects:  75% (50/66)   
remote: Compressing objects:  77% (51/66)   
remote: Compressing objects:  78% (52/66)   
remote: Compressing objects:  80% (53/66)   
remote: Compressing objects:  81% (54/66)   
remote: Compressing objects:  83% (55/66)   
remote: Compressing objects:  84% (56/66)   
remote: Compressing objects:  86% (57/66)   
remote: Compressing objects:  87% (58/66)   
remote: Compressing objects:  89% (59/66)   
remote: Compressing objects:  90% (60/66)   
remote: Compressing objects:  92% (61/66)   
remote: Compressing objects:  93% (62/66)   
remote: Compressing objects:  95% (63/66)   
remote: Compressing objects:  96% (64/66)   
remote: Compressing objects:  98% (65/66)   
remote: Compressing objects: 100% (66/66)   
remote: Compressing objects: 100% (66/66), done.
remote: Total 94 (delta 43), reused 75 (delta 26), pack-reused 0
Unpacking objects:   1% (1/94)
Unpacking objects:   2% (2/94)
Unpacking objects:   3% (3/94)
Unpacking objects:   4% (4/94)
Unpacking objects:   5% (5/94)
Unpacking objects:   6% (6/94)
Unpacking objects:   7% (7/94)
Unpacking objects:   8% (8/94)
Unpacking objects:   9% (9/94)
Unpacking objects:  10% (10/94)
Unpacking objects:  11% (11/94)
Unpacking objects:  12% (12/94)
Unpacking objects:  13% (13/94)
Unpacking objects:  14% (14/94)
Unpacking objects:  15% (15/94)
Unpacking objects:  17% (16/94)
Unpacking objects:  18% (17/94)
Unpacking objects:  19% (18/94)
Unpacking objects:  20% (19/94)
Unpacking objects:  21% (20/94)
Unpacking objects:  22% (21/94)
Unpacking objects:  23% (22/94)
Unpacking objects:  24% (23/94)
Unpacking objects:  25% (24/94)
Unpacking objects:  26% (25/94)
Unpacking objects:  27% (26/94)
Unpacking objects:  28% (27/94)
Unpacking objects:  29% (28/94)
Unpacking objects:  30% (29/94)
Unpacking objects:  31% (30/94)
Unpacking objects:  32% (31/94)
Unpacking objects:  34% (32/94)
Unpacking objects:  35% (33/94)
Unpacking objects:  36% (34/94)
Unpacking objects:  37% (35/94)
Unpacking objects:  38% (36/94)
Unpacking objects:  39% (37/94)
Unpacking objects:  40% (38/94)
Unpacking objects:  41% (39/94)
Unpacking objects:  42% (40/94)
Unpacking objects:  43% (41/94)
Unpacking objects:  44% (42/94)
Unpacking objects:  45% (43/94)
Unpacking objects:  46% (44/94)
Unpacking objects:  47% (45/94)
Unpacking objects:  48% (46/94)
Unpacking objects:  50% (47/94)
Unpacking objects:  51% (48/94)
Unpacking objects:  52% (49/94)
Unpacking objects:  53% (50/94)
Unpacking objects:  54% (51/94)
Unpacking objects:  55% (52/94)
Unpacking objects:  56% (53/94)
Unpacking objects:  57% (54/94)
Unpacking objects:  58% (55/94)
Unpacking objects:  59% (56/94)
Unpacking objects:  60% (57/94)
Unpacking objects:  61% (58/94)
Unpacking objects:  62% (59/94)
Unpacking objects:  63% (60/94)
Unpacking objects:  64% (61/94)
Unpacking objects:  65% (62/94)
Unpacking objects:  67% (63/94)
Unpacking objects:  68% (64/94)
Unpacking objects:  69% (65/94)
Unpacking objects:  70% (66/94)
Unpacking objects:  71% (67/94)
Unpacking objects:  72% (68/94)
Unpacking objects:  73% (69/94)
Unpacking objects:  74% (70/94)
Unpacking objects:  75% (71/94)
Unpacking objects:  76% (72/94)
Unpacking objects:  77% (73/94)
Unpacking objects:  78% (74/94)
Unpacking objects:  79% (75/94)
Unpacking objects:  80% (76/94)
Unpacking objects:  81% (77/94)
Unpacking objects:  82% (78/94)
Unpacking objects:  84% (79/94)
Unpacking objects:  85% (80/94)
Unpacking objects:  86% (81/94)
Unpacking objects:  87% (82/94)
Unpacking objects:  88% (83/94)
Unpacking objects:  89% (84/94)
Unpacking objects:  90% (85/94)
Unpacking objects:  91% (86/94)
Unpacking objects:  92% (87/94)
Unpacking objects:  93% (88/94)
Unpacking objects:  94% (89/94)
Unpacking objects:  95% (90/94)
Unpacking objects:  96% (91/94)
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