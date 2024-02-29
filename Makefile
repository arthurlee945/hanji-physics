test:
	air --build.cmd "go build -o tmp example/main.go"
gravity:
	air --build.cmd "go build -o tmp example/gravity/main.go"
noise:
	air --build.cmd "go build -o tmp example/noise/main.go"
walker:
	air --build.cmd "go build -o tmp example/walker/main.go"
pendulum:
	air --build.cmd "go build -o tmp example/pendulum/main.go"