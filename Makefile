all:
	go test -coverprofile=c.out ./...
debug:
	go test -gcflags '-N -l' -c . && dlv exec ./*.test  # better if we get the changed line and set it automatically
	cp $(HOME) ; dlv test --build-flags '-N -l'
