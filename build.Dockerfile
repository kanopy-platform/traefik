FROM golang:1.16-alpine

RUN apk --update upgrade \
    && apk --no-cache --no-progress add git mercurial bash gcc musl-dev curl tar ca-certificates tzdata \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*

RUN go get golang.org/x/lint/golint \
&& go get github.com/client9/misspell/cmd/misspell

# Which docker version to test on
ARG DOCKER_VERSION=18.09.7

# Download docker
RUN mkdir -p /usr/local/bin \
    && curl -fL https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz \
    | tar -xzC /usr/local/bin --transform 's#^.+/##x'

# Download go-bindata binary to bin folder in $GOPATH
RUN mkdir -p /usr/local/bin \
    && curl -fsSL -o /usr/local/bin/go-bindata https://github.com/containous/go-bindata/releases/download/v1.0.0/go-bindata \
    && chmod +x /usr/local/bin/go-bindata

# Download misspell binary to bin folder in $GOPATH
RUN  curl -sfL https://raw.githubusercontent.com/client9/misspell/master/install-misspell.sh | bash -s -- -b $GOPATH/bin v0.3.4

WORKDIR /go/src/github.com/traefik/traefik

# Download go modules
COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on GOPROXY=https://proxy.golang.org go mod download

COPY . /go/src/github.com/traefik/traefik
