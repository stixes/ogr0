FROM alpine as keybringer
RUN apk add openssh-client; ssh-keygen -f /host_key

FROM golang:alpine AS appbuilder

WORKDIR /app

COPY ogr0.go /app/
COPY --from=keybringer /host_key /app/

RUN go mod init ogr0
RUN go get -d 
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o ogr0
RUN chmod 555 host_key; chmod 777 ogr0 

FROM scratch

COPY --from=appbuilder /app/host_key /app/ogr0 /
USER 65534
ENV USER=65534

ENTRYPOINT ["/ogr0"]
