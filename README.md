# quake-log-go

## About

This a log parser for a Quake log file. It generates two reports: player ranking and death causes.

## Tests

To run the tests use the following command:

```shell
# If make (Makefile) command available
$ make test

# Windows and Linux
$ go test -v ./... -coverprofile=coverage.out

```

## Player Ranking Report (Item 3.2)

To execute the player ranking report use the following command:

```shell
# Windows
$ go run .\cmd\player-report\player-report.go -filepath C:\Path\To\File\file.extension

# Linux
$ go run ./cmd/player-report/player-report.go -filepath ~/path/to/file/file.extension
```

## Death Cause Report (3.3)

To execute the death cause report use the following command:

```shell
# Windows
$ go run .\cmd\weapon-report\weapon-report.go -filepath C:\Path\To\File\file.extension

# Linux
$ go run ./cmd/weapon-report/weapon-report.go -filepath ~/path/to/file/file.extension
```
