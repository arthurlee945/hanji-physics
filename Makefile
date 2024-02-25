test:
	air --build.cmd "go build -o tmp example/main.go"
mover:
	air --build.cmd "go build -o tmp example/mover/main.go"
noise:
	air --build.cmd "go build -o tmp example/noise/main.go"
walker:
	air --build.cmd "go build -o tmp example/walker/main.go"