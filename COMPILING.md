## Steps to compile the test package

1. source .source
2. go install application
3. go run src/application/main.go

## testing the tasks package
1. go test tasks

## Installing and running "rerun"
1. go get github.com/skelterjohn/rerun
2. go install github.com/skelterjohn/rerun
3. go install application
4. rerun --test application
