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

gt:
	ls *.go | grep -v "test" | xargs -I{} bash -c ' grep "package main" {} || sed -i "1ipackage main" {}; gotests -all {} > $$(echo {} | sed "s/.go/_test.go/")'
