dependencies:
	cd src/simplecmd; dep ensure

run:
	cd src/simplecmd; go run app.go

install:
	cd src/simplecmd; go install -v ./...