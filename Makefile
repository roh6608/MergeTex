
build:
	go build -o bin/MergeTex src/main.go

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/MergeTex_linux_amd64 src/main.go
	GOOS=linux GOARCH=arm go build -o bin/MergeTex_linux_arm src/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/MergeTex_windows_amd64.exe src/main.go