# Dating App REST API
### By Kris Sukoco
### As a requirement for technical test completion on Deall. Written in Go/Golang.
<br/>

## See the app in action
- [API Documentation](https://api.krissukoco.dev/dating-app-deall/api/v1/docs/index.html)
- [API Endpoint](https://api.krissukoco.dev/dating-app-deall/api/v1)

## How the code is structured
1. Entrypoint of the server is on `cmd` directory: `cmd/server/main.go`
2. `config` directory contains configuration files for the server. On `config.go` file, you can find the configuration for the server, database, and the JWT secret key, using `viper` to Unmarshal yaml files.
3. `docs` directory contains the API documentation, using `swagger` to generate the documentation.
4. `internal` directory contains all the services and internal packages for the server, including new database connection using `GORM` for `PostgreSQL` and database entity `models`.
5. `pkg` directory contains packages that can be reused, currently containing a simple `utils` package.
6. `tests` directory contains utilities required to run unit tests, including mocks and API testing.
7. Each functional components on `internal` contains its own `repository`, which handles queries and database operations, `service` which handles business logic, and `api` which handles the API endpoints.
<br/>
<br/>

## Services
There are 4 main services in this app:
1. `user` service handle user data, including creation, getting, and queries required for other services.
2. `auth` service handle authentication, including login and registration. Also includes auth middleware for other services.
3. `subscription` service handle user subscription, which can be used to check if a user is subscribed or not and whether the subscription is active. If user subscribed, user will get <b>Unlimited Swipes</b>.
4. `match` service handle logic for matching users, including get a new match and like a match. It will also check limitation of number of swipes per day.
<br/>

## API
- Base Path: `/api/v1`
- Swagger Documentation is available at `/api/v1/docs/index.html`

## How to Run
1. Clone this repository
2. Install `docker` and `docker-compose` if not yet on your machine.
3. Run `docker compose -f docker-compose-local.yaml up` on the root directory of this repository.
4. Server will be available on `localhost:8080`
5. Feel free to explore docs and test the API.

## Perform tests
```go
go test ./... --cover
```