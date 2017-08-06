# golodns

`golodns` is a simple DNS server replying to type A and NXDOMAIN queries and maps those to a specific IP address.

It is meant to help with local development so that there is no need to fiddle around with `/etc/hosts` anymore.

## Installation

As it is a simple binary you can just download the respective file.

>It is recommended to place the binary to a `${PATH}` accessible location (f.e. `/usr/local/bin` on macOS).

## Configuration

In order to setup your local environment `golodns` provides means to add entries to the `resolver` system.

There are three commands that can help you with managing the domain mappings:

* `list`: List all registered resolvers (should be empty per default)
* `install`: Register a new resolver for a specific domain
* `uninstall`: Uninstall an existing resolver (if created by `golodns`)

### List

In order to list the current resolvers run:

```
golodns list
```

>This outputs something like "No resolvers found." initially as normally there are none. If there are registered ones 
it will indicate which of them are managed by `golodns`.

### Install

If you want to setup a new resolver run

```
sudo golodns install -domain dev -addr 127.0.0.1 -port 5300
```

to resolve the `dev` TLD via a nameserver listening on `127.0.0.1:5300`.

>You need to run it as `sudo` user as it effectively writes to the `/etc` path which is not writable for normal users.

### Uninstall

If you want to remove a resolver run

```
sudo golodns uninstall -domain dev
```

>`golodns` only removes resolvers that where created by it.

### Autocompletion

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

>You need to add special flags like `-port` etc. ofc if desired / configured. Per default `golodns` resolves via 
`127.0.0.1:5300` to `127.0.0.1` for domains.

Then test the setup by doing f.e. a `ping` call:

```
ping -c 1 some.dev
```

## Development

This project uses `glide` to handle dependencies in a more reliable way.

In order to build the binary just clone the repository and build it `go build`. For a smaller binary please 
use `go build -ldflags "-w"` instead.

## License

MIT

## Credits

`golodns` started as a clone of [https://github.com/robbiev/devdns](devdns).
