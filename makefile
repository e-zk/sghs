.POSIX:
.SUFFIXES:
.PHONY: clean uninstall

# install location
PREFIX = /usr/local
BINDIR = bin

sghs: main.go
	go build -ldflags "-w -s" -o sghs -v main.go

install: sghs
	install -c -m 0755 sghs $(PREFIX)/$(BINDIR)

uninstall:
	rm -f $(PREFIX)/$(BINDIR)/sghs

clean:
	go clean
	rm -f sghs
