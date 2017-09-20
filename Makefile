default: install

build: build-linux build-macos

install: install-lock install-unlock install-manpages

install-manpages:
	cp ./man/* /usr/local/share/man/man1/

install-lock:
	go install ./git-lock-remote

install-unlock:
	go install ./git-unlock-remote

build-linux:
	GOOS=linux GOARCH=amd64 go build -o git-lock-remote.linux-x64 ./git-lock-remote
	GOOS=linux GOARCH=amd64 go build -o git-unlock-remote.linux-x64 ./git-unlock-remote

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o git-lock-remote.darwin-x64 ./git-lock-remote
	GOOS=darwin GOARCH=amd64 go build -o git-unlock-remote.darwin-x64 ./git-unlock-remote
