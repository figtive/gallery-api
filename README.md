# Gallery API

API backend for Gallery PPL Fasilkom UI.

## Develop

`gallery-api` uses [Go Modules](https://blog.golang.org/using-go-modules) module/dependency manager, hence at least Go 1.11 is required. To ease development, [comstrek/air](https://github.com/cosmtrek/air) is used to live-reload the application. Install the tools as documented.

To begin developing, simply enter the sub-directory and run the development server:

```shell
$ go mod tidy
$ air
```

## Deploy

`gallery-api` is containerized and pushed to [Docker Hub](https://hub.docker.com/r/figtive/galleryppl). It is tagged based on its application version, e.g. `figtive/galleryppl:api` or `figtive/galleryppl:api-v1.1.0`.

To run `gallery-api`, run the following:

```shell
$ docker run --name gallery-api --env-file ./.env -p 8080:8080 -d figtive/galleryppl:api
```

### Dependencies

The following are required for `gallery-api` to function properly:

- PostgreSQL

Their credentials must be provided in the configuration file.

### PostgreSQL UUID Extension

UUID support is also required in PostgreSQL. For modern PostgreSQL versions (9.1 and newer), the contrib module `uuid-ossp` can be enabled as follows:

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```
