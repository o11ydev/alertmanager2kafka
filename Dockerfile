FROM golang:1.15 as build

WORKDIR /go/src/github.com/webdevops/alertmanager2kafka

# Get deps (cached)
COPY ./go.mod /go/src/github.com/webdevops/alertmanager2kafka
COPY ./go.sum /go/src/github.com/webdevops/alertmanager2kafka
COPY ./Makefile /go/src/github.com/webdevops/alertmanager2kafka
RUN make dependencies

# Compile
COPY ./ /go/src/github.com/webdevops/alertmanager2kafka
RUN make test
# RUN make lint
RUN make build
RUN ./alertmanager2kafka --help

#############################################
# FINAL IMAGE
#############################################
FROM gcr.io/distroless/static
ENV LOG_JSON=1
COPY --from=build /go/src/github.com/webdevops/alertmanager2kafka/alertmanager2kafka /
USER 1000
ENTRYPOINT ["/alertmanager2kafka"]
