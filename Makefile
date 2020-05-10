all:
	go test -coverprofile=c.out ./...
debug:
	dlv test --build-flags '-N -l'
d:
	go_add_debug
