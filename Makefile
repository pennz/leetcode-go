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

gpm: # add package main to each file
	ls todo/*.go | grep -v "_test.go" | xargs -I{} bash -c 'sed -i "s/package main//" {}; head -n 1 {} | grep "package main" {} >/dev/null 2>&1 || sed -i "1ipackage main" {}'

gt:
	ls todo/*.go | grep -v "test" | head -n 1 | xargs -I{} bash -xc 'gotests -all {} > $$(echo {} | sed "s/.go$$/_test.go/")'

get_job_trace:
	glc get project-job-trace 11180959 $$(glc list project-jobs 11180959 -f json | sed '1d' | jq '.[0].id') -i
