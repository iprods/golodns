# golodns

`golodns` is a simple DNS server replying to type A and NXDOMAIN queries and maps those to a specific IP address.

It is meant to help with local development so that there is no need to fiddle around with `/etc/hosts` anymore.

## Installation

As it is a simple binary you can just download the respective file.

## Configuration

### Local

Per default `golodns` listens on `127.0.0.1:5300`. This can be changed by running `golodns -ip="127.0.0.1:6300"`.

With this set you need additionally to prepare the `resolv` system. Therefore add the following file `/etc/resolver/dev`:

```
sudo mkdir -p /etc/resolver
sudo tee /etc/resolver/dev >/dev/null <<EOF
nameserver 127.0.0.1
port 5300
EOF
```

Then test the setup by doing f.e. a `ping` call:

```
ping -c 1 some.dev
```

#### Autocompletion

Autocompletion for bash and zsh goes as simple as:

```
golodns -autocomplete-install
```

Please re-login to the respective shell after that.

The removal is equally easy (even with autcompletion):

```
golodns -autocomplete-uninstall
```

## Usage

If you have set up your local environment you just need to run:

```
golodns serve
```

>You need to add special flags like `-port` etc. ofc if desired.

## Development

This project uses `glide` to handle dependencies in a more reliable way.

In order to build the binary just clone the repository and build it `go build`. For a smaller binary please 
use `go build -ldflags "-w"` instead.

## License

MIT

## Credits

`golodns` started as a clone of [https://github.com/robbiev/devdns](devdns).
