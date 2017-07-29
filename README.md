# golodns

`golodns` is a simple DNS server replying to type A and NXDOMAIN queries and maps those to a specific IP address.

It is meant to help with local development so that there is no need to fiddle around with `/etc/hosts` anymore.

## Installation

As it is a simple binary you can just download the respective file.

Alternatively you can also use the Docker image.

## Configuration

Per default `golodns` listens on `127.0.0.1:5300`. This can be changed by runnnig `golodns -ip="127.0.0.1:6300"`.

## Development

In order to build the binary just clone the repository and build it `go build`. For a smaller binary please 
use `go build -ldflags "-w"` instead.

## License

MIT

## Credits

`golodns` started as a clone of [https://github.com/robbiev/devdns](devdns).

