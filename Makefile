.POSIX:
.SUFFIXES:
.PHONY: clean uninstall

# install location
PREFIX = /usr/local

sghs: main.go
	go build -ldflags "-w -s" -o sghs -v main.go

install: sghs
	install -c -m 0755 sghs $(PREFIX)/bin

uninstall:
	rm -f $(PREFIX)/bin/sghs

clean:
	go clean
	rm -f sghs
