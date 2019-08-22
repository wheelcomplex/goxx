# Gox - Simple Go Cross Compilation

Gox is a simple, no-frills tool for Go cross compilation that behaves a
lot like standard `go build`. Gox will parallelize builds for multiple
platforms. Gox will also build the cross-compilation toolchain for you.

## Installation

To install Gox, please use `go get`. We tag versions so feel free to
checkout that tag and compile.

```bash
$ go get github.com/wheelcomplex/goxx
...
$ goxx -h
...
```

## Usage

If you know how to use `go build`, then you know how to use Gox. For
example, to build the current package, specify no parameters and just
call `goxx`. Gox will parallelize based on the number of CPUs you have
by default and build for every platform by default:

```bash
$ goxx
Number of parallel builds: 4

-->      darwin/386: github.com/wheelcomplex/goxx
-->    darwin/amd64: github.com/wheelcomplex/goxx
-->       linux/386: github.com/wheelcomplex/goxx
-->     linux/amd64: github.com/wheelcomplex/goxx
-->       linux/arm: github.com/wheelcomplex/goxx
-->     freebsd/386: github.com/wheelcomplex/goxx
-->   freebsd/amd64: github.com/wheelcomplex/goxx
-->     openbsd/386: github.com/wheelcomplex/goxx
-->   openbsd/amd64: github.com/wheelcomplex/goxx
-->     windows/386: github.com/wheelcomplex/goxx
-->   windows/amd64: github.com/wheelcomplex/goxx
-->     freebsd/arm: github.com/wheelcomplex/goxx
-->      netbsd/386: github.com/wheelcomplex/goxx
-->    netbsd/amd64: github.com/wheelcomplex/goxx
-->      netbsd/arm: github.com/wheelcomplex/goxx
-->       plan9/386: github.com/wheelcomplex/goxx
```

Or, if you want to build a package and sub-packages:

```bash
$ goxx ./...
...
```

Or, if you want to build multiple distinct packages:

```bash
$ goxx github.com/wheelcomplex/goxx github.com/hashicorp/serf
...
```

Or if you want to just build for linux:

```bash
$ goxx -os="linux"
...
```

Or maybe you just want to build for 64-bit linux:

```bash
$ goxx -osarch="linux/amd64"
...
```

And more! Just run `goxx -h` for help and additional information.

## Versus Other Cross-Compile Tools

A big thanks to these other options for existing. They each paved the
way in many aspects to make Go cross-compilation approachable.

* [Dave Cheney's golang-crosscompile](https://github.com/davecheney/golang-crosscompile) -
  Gox compiles for multiple platforms and can therefore easily run on
  any platform Go supports, whereas Dave's scripts require a shell. Gox
  will also parallelize builds. Dave's scripts build sequentially. Gox has
  much easier to use OS/Arch filtering built in.

* [goxc](https://github.com/laher/goxc) -
  A very richly featured tool that can even do things such as build system
  packages, upload binaries, generate download webpages, etc. Gox is a
  super slim alternative that only cross-compiles binaries. Gox builds packages in parallel, whereas
  goxc doesn't. Gox doesn't enforce a specific output structure for built
  binaries.
