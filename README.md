# ayo

Spin up an io interface in a pinch.

ayo is a minimal input/output proxy for common interfaces. It manages connections for inputs and relays messages to outputs in a "turnout" concurrency pattern.

## Install

```bash
go get github.com/jamiekieranmartin/tryp
```

Otherwise you can download the binary from [Releases](https://github.com/jamiekieranmartin/tryp/releases)

## Usage

### CLI

```bash
ayo -config "./config.toml"
```

### Golang SDK

```go
// make new ayo instance
io, err := ayo.New(*config)
if err != nil {
	panic(err)
}

// listen and serve io interfaces
err = io.ListenAndServe()
if err != nil {
	panic(err)
}
```

## CLI flags

### `-config`

Path a custom configuration file. Defaults to `./config.toml`.

```
ayo -config "./path/to/my/file.toml"
```

## TOML Configuration

```toml
# Inputs

[[input.tcp]]
port="5001"

[[input.udp]]
port="5002"

[[input.http]]
port="80"

# Outputs

[[output.tcp]]
hostname="1.1.1.1"
port="6001"

[[output.udp]]
hostname="1.1.1.1"
port="6002"

[[output.http]]
uri="https://example.com"
```
