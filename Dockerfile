# simple base image
FROM alpine
MAINTAINER Allan Shone <allan.shone@gmail.com>

# volumes for information
VOLUME ["/opt/music/", "/opt/nottify/"]

# set some default environment variables for the container
ENV GOPATH=/opt/go/ \
    GOROOT=/opt/go \
    NOTTIFY_PORT=80

# define 80 for opening
EXPOSE 80

# command that will run nottify itself when the container runs
ENTRYPOINT ["/opt/go/bin/nottify", "start"]

# copy the content across
RUN mkdir -p /opt/go/src/github.com/cloudcloud/nottify
COPY . /opt/go/src/github.com/cloudcloud/nottify/

# move the executable into a nicer location
COPY nott /opt/go/bin/nottify
