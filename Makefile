test:
	air --build.cmd "go build -o tmp example/main.go"
gravity:
	air --build.cmd "go build -o tmp example/gravity/main.go"
noise:
	air --build.cmd "go build -o tmp example/noise/main.go"
walker:
	air --build.cmd "go build -o tmp example/walker/main.go"
osc:
	air --build.cmd "go build -o tmp example/oscillation/main.go"