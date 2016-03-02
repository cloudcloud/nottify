# empty base image
FROM alpine
MAINTAINER Allan Shone <allan.shone@gmail.com>

# copy the content across
RUN mkdir -p /opt/go/src/github.com/cloudcloud/nottify
COPY . /opt/go/src/github.com/cloudcloud/nottify/

RUN cp /opt/go/src/github.com/cloudcloud/nottify/nott /opt/go/nottify

ENV GOPATH=/opt/go/ \
    GOROOT=/opt/go

# define 80 for opening
ENV PORT 80
EXPOSE 80

ENTRYPOINT ["/opt/go/nottify", "start"]

