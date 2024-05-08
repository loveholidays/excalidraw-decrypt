FROM golang:1.22.2 as build

RUN apt-get update && \
    apt-get install -y ca-certificates libssl-dev cpio

WORKDIR /app
ADD . /app

RUN make check build

FROM busybox

COPY --from=build /app/bin/excalidraw-decrypt /go/bin/excalidraw-decrypt
ENV PATH $PATH:/go/bin
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/go/bin/excalidraw-decrypt"]