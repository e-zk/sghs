# shitty go http server

bare-bones and basically useless http server for serving stuff quickly.

## features

- serves stuff
- logs all requests (`-q` to disable)
- optional tls support (`-t`)
- optional chroot support (`-c`)

## installation

to install the latest version to `$GOPATH/bin`:

	$ go install go.zakaria.org/sghs@latest

## usage

	usage: sghs [-l string] [-p path] [-q] [-t string]
	where:
	  -c                            chroot to the path that is being served
	  -l string			listen string "[address]:port" (default ":8080")
	  -p path			path to serve (default ".")
	  -q				quiet mode (don't print logs)
	  -t 'cert_path:key_path'	enable TLS using the given cert and cert key

### examples

serve the current directory on port 8080:

	$ sghs
	Serving ./ on :8080

serve `/var/www` on `0.0.0.0:7070`:

	$ sghs -p /var/www -l "0.0.0.0:7070"
	Serving /var/www on 0.0.0.0:7070

serve the current directory on port 7000 using certs in `/etc/ssl`:

	$ sghs -l ":7000" -t "/etc/ssl/mycert.pem:/etc/ssl/private/mycert-key.pem"
	Serving . on :7070 (tls)
	Using:
	cert: /etc/ssl/mycert.pem
	key: /etc/ssl/private/mycert-key.pem
