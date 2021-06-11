
gdate: cmd/gdate/main.go
	go build -o $@ $^

gdate-server: cmd/gdate-server/main.go
	CGO_ENABLED=0 go build -o $@ $^

.PHONY: clean
clean:
	rm -f gdate gdate-server
