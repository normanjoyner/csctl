FROM golang:1.12.1-alpine3.9 as builder

# Add build tools
RUN apk update && \
    apk add --no-cache git gcc musl-dev curl

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && \
    chmod +x /usr/local/bin/dep

ENV SRC_DIR=/go/src/github.com/containership/csctl/

WORKDIR /app

# Install deps before adding rest of source so we can cache the resulting vendor dir
COPY Gopkg.toml Gopkg.lock $SRC_DIR
RUN cd $SRC_DIR && \
    dep ensure -vendor-only

# Add the source code
COPY . $SRC_DIR

# Build it:
ARG GIT_DESCRIBE
ARG GIT_COMMIT
RUN cd $SRC_DIR && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
        go build -ldflags \
        "-X github.com/containership/csctl/pkg/buildinfo.gitDescribe=${GIT_DESCRIBE} \
        -X github.com/containership/csctl/pkg/buildinfo.gitCommit=${GIT_COMMIT} \
        -X github.com/containership/csctl/pkg/buildinfo.unixTime=`date '+%s'` \
        -w" \
        -a -tags netgo \
        -o csctl main.go && \
    cp csctl /app/


# Create Docker image of just the binary
FROM scratch as runner
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/csctl .

ENTRYPOINT ["./csctl"]
