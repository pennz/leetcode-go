all:
	go test -coverprofile=c.out ./...
tv:
	go test . -v
t:
	go test .
debug:
	dlv test --build-flags '-N -l'
d:
	go_add_debug
