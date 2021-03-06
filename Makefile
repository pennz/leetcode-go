export PATH := $(shell dirname $$(which npm)):$(PWD)/node_modules/.bin:$(HOME)/.local/bin:$(PATH)

PROJECT=leetcode-go
PROJECT_ID := $(shell glc list projects --owned -s $(PROJECT) -f json | sed '1d' | jq '.[0].id')

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
	ls *.go | grep -v "_test.go" |  xargs -I{} bash -xc 'gotests -all {} | sed "1d" > $$(echo {} | sed "s/.go$$/_test.go/")'

get_jobs:
	glc list project-jobs $(PROJECT_ID) -f json | sed '1d' | jq '.[].id' | sort
get_job_trace:
	glc get project-job-trace $(PROJECT_ID) $$(glc list project-jobs $(PROJECT_ID) -f json | sed '1d' | jq '.[0].id') -i

lc:
	leetcode show $(N)

check:
	which leetcode
	echo $(PATH)
