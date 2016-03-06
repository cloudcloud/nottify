# The binary requirements building before Docker build is run.
#
# CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o nott .

# simple base image
FROM alpine
MAINTAINER Allan Shone <allan.shone@gmail.com>

# copy the content across
RUN mkdir -p /opt/go/src/github.com/cloudcloud/nottify
COPY . /opt/go/src/github.com/cloudcloud/nottify/

# move the executable into a nicer location
RUN cp /opt/go/src/github.com/cloudcloud/nottify/nott /opt/go/nottify

# set some default environment variables for the container
ENV GOPATH=/opt/go/ \
    GOROOT=/opt/go

# define 80 for opening
ENV PORT 80
EXPOSE 80

# command that will run nottify itself when the container runs
ENTRYPOINT ["/opt/go/nottify", "start"]
