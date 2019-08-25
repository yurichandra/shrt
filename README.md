# shrt
URL shortener service. To make it short, I called shrt (shirt). Equipped with Redis and Docker (For learning purpose).

## How to use
In development mode, don't forget to specified .env file.

***Dockerize (use docker environment)***

**Build docker image**

`docker build -f Dockerfile -t <put any tag you want> .`

**Run container**

`docker-compose up -d`

Please keep in note, you need to start postgres/mysql and redis container first before starts this project container.

***localhost-mode***

**Download dependencies**

```
go get -v ./... 
dep ensure -update
```

**Build, migrate and run the app**

```
go build
./shrt migrate
./shrt serve
```
