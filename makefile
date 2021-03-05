.POSIX:
.SUFFIXES:
.PHONY: clean

sghs: main.go
	go build -ldflags "-w -s" -o sghs -v main.go

install: sghs
	install -c -s -m 0755 sghs $(PREFIX)/bin

clean:
	go clean
	rm -f sghs
