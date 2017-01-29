## Wstats

* Natural language TCP server.
* A minimal docker container.
* Automatically pushes it to dockerhub if circlec build pass.

## Setup

Installation

```sh
mkdir -p $GOPATH/src/github.com/rafaeljesus
cd $GOPATH/src/github.com/rafaeljesus
git clone https://github.com/rafaeljesus/wstats.git
cd wstats
glide install
sh build && sh build-container
docker run -it -t -p 80:8080 --name wstats rafaeljesus/wstats
```

Running Tests
```sh
go test $(go list ./... | grep -v /vendor/)
```

## API
See [docs](./docs/README.md)

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges

[![Build Status](https://circleci.com/gh/rafaeljesus/wstats.svg?style=svg)](https://circleci.com/gh/rafaeljesus/wstats)
[![Go Report Card](https://goreportcard.com/badge/github.com/rafaeljesus/wstats)](https://goreportcard.com/report/github.com/rafaeljesus/wstats)
[![](https://images.microbadger.com/badges/image/rafaeljesus/wstats.svg)](https://microbadger.com/images/rafaeljesus/wstats "Get your own image badge on microbadger.com")

---

> GitHub [@rafaeljesus](https://github.com/rafaeljesus) &nbsp;&middot;&nbsp;
> Medium [@_jesus_rafael](https://medium.com/@_jesus_rafael) &nbsp;&middot;&nbsp;
> Twitter [@_jesus_rafael](https://twitter.com/_jesus_rafael)
