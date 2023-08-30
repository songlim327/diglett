build-linux-64: 
	GOOS=linux GOARCH=amd64 go build -o diglett-amd64-linux main.go

build-linux-32:
	GOOS=darwin GOARCH=386 go build -o diglett-386-linux main.go