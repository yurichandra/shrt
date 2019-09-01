# shrt
URL shortener service API. To make it short, I called shrt (shirt). Equipped with Redis and Docker (For learning purpose).

## How to use
To use this API, I provide two different ways, with auth and no-auth.

**Auth**

To use with auth, you must be authorized first by hit `/auth/authorize` endpoint and providing two field request `email` and `password`.

- ***Request***

```
POST /auth/authorize
{
  "email": "YOUR_EMAIL_HERE",
  "password": "YOUR_PASSWORD_HERE"
}
```

Hitting endpoint such as above, will returning back response `api_key` that can be useful for shortening url for specific user.

- ***Response***

```
{
  "api_key": "YOUR_API_KEY"
}
```

`api_key` from response above, can be specified as a header request with field `api_key` and `YOUR_API_KEY` as value of attributes while hitting `/shorten` endpoint to shorten an url.

**No-Auth**

To use no-auth, you may directly hit `/shorten` endpoint without specify `api_key` headers. It will generate unique keys for each request.

In development mode, don't forget to specified .env file.

***Dockerize (use docker environment)***

- **Build docker image**

`docker build -f Dockerfile -t <put any tag you want> .`

- **Run container**

`docker-compose up -d`

Please keep in note, you need to start postgres/mysql and redis container first before starts this project container.

***localhost-mode***

- **Download dependencies**

```
go get -v ./... 
dep ensure -update
```

- **Build, migrate, seed and run the app**

```
go build
./shrt migrate
./shrt seed
./shrt serve
```
Any feedback of this project are very welcome as it can be my input to make it better.
