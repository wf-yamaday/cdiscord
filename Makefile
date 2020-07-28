.PHONY: release
release:
		GOOS=darwin GOARCH=amd64 go build -o cdiscord main.go
		tar czf cdiscord-darwin-amd64.tar.gz cdiscord
		rm -f cdiscord
		GOOS=linux GOARCH=amd64 go build -o cdiscord main.go
		tar czf cdiscord-linux-amd64.tar.gz cdiscord
		rm -f cdiscord
