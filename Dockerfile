FROM golang:alpine as build
MAINTAINER Sajay Antony <sajaya@microsoft.com>
WORKDIR /go/src/github.com/sajayantony/simple-auth-plugin 
COPY . .
RUN go build

FROM golang:alpine
COPY --from=build /go/src/github.com/sajayantony/simple-auth-plugin/simple-auth-plugin /bin/simple-auth-plugin

ENTRYPOINT ["/bin/simple-auth-plugin"]