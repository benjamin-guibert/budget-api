# Budget API

Flexible budget management application with monthly adjustment and month-length
costs splitting (API side).

[![license-shield][]](LICENSE)
![last-commit-shield][]
![go-version-shield][]

## Table Of Contents

- [Setup](#setup)
- [Usage](#usage)
- [Deployment](#deployment)
- [Scripts](#scripts)
- [Built With](#built-with)
- [Release History](#release-history)
- [Versionning](#versionning)
- [Authors](#authors)
- [License](#license)

## Setup

- Install the dependencies:

```bash
go get .
```

- Run the database (see [usage](#Usage))
- Create the `.env` file from [`.env.sample`](.env.sample) and set the values:

```bash
cp .env.sample .env
```

- seed the database & create the test database:

```bash
make db-seed
make db-test-create
```

See [`docker-compose`](docker-compose.yml) configuration and the [`Makefile`](Makefile).

## Usage

- Run the database:

```bash
make db-start
```

- Run the API locally:

```bash
go run .
```

- The API will be available from `http://localhost:8080` by default.

- Get more information about commands:

```bash
go run . --help
```

See the [`Makefile`](Makefile) for available commands.

## Deployment

## Built With

[Go](https://golang.org) |
[PostgreSQL](https://www.postgresql.org) |
[Docker](https://www.docker.com) |
[GORM](https://gorm.io/index.html)
[mux](https://github.com/gorilla/mux) |
[godotenv](https://github.com/joho/godotenv)
[validator](https://github.com/go-playground/validator)

## Release History

Check the [`CHANGELOG.md`](CHANGELOG.md) file for the release history.

## Versionning

We use [SemVer](http://semver.org/) for versioning. For the versions available,
see the [tags on this repository][tags-link].

## Authors

- **[Benjamin Guibert](https://github.com/benjamin-guibert)**: Creator & main
  contributor

See also the list of [contributors][contributors-link] who participated in this
project.

## License

[![license-shield][]](LICENSE)

This project is licensed under the MIT License. See the [`LICENSE`](LICENSE)
file for details.

[test-workflow-shield]: https://github.com/benjamin-guibert/budget-api/workflows/Test/badge.svg?branch=main
[contributors-link]: https://github.com/benjamin-guibert/budget-api/contributors
[license-shield]: https://img.shields.io/github/license/benjamin-guibert/budget-api.svg
[go-version-shield]: https://img.shields.io/github/go-mod/go-version/benjamin-guibert/budget-api
[last-commit-shield]:
https://img.shields.io/github/last-commit/benjamin-guibert/budget-api
