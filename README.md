# shitty go http server

bare-bones and basically useless http server for serving stuff quickly.

## install

to install the latest version to $GOPATH/bin:

	$ go install go.zakaria.org/sghs@latest

## usage

serve the current directory on port 8080:

	$ sghs
	Serving ./ on :8080

serve `/var/www` on `0.0.0.0:7070`:

	$ sghs -p /var/www -a "0.0.0.0:7070"
	Serving /var/www on 0.0.0.0:7070

