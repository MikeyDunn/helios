FROM golang:1.10 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/mikeydunn/helios
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./

# Set build environment
ENV CGO_ENABLED=0
ENV GOOS=linux

# Build go binary
RUN go build -a \
	-ldflags '-extldflags "-static"' \
	-installsuffix cgo \
	-o helios \
	github.com/mikeydunn/helios/cmd/helios

# Copy binary to scratch image
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/mikeydunn/helios/helios .
ENTRYPOINT ["/helios"]
