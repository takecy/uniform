prepare:
	go get -u github.com/tools/godep
	godep restore

update:
	go get -u ./...
	rm -rf Godep
	godep save ./...

test:
	go test ./...
