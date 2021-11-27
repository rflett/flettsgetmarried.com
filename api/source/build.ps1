$Env:GOOS = "linux"
$Env:GOARCH = "amd64"

go build -ldflags="-s -w" -o ../bin/rsvp   rest/rsvp/main.go
go build -ldflags="-s -w" -o ../bin/search rest/search/main.go
go build -ldflags="-s -w" -o ../bin/guests rest/guests/main.go
