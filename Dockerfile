FROM golang:1.22.1-bookworm as build

COPY . /snapsync
WORKDIR /snapsync

RUN go build -o snapsync main.go

FROM ubuntu:22.04

RUN apt update && apt-get upgrade -y && apt install rsync -y

COPY --from=build /snapsync/snapsync /snapsync/snapsync
COPY --from=build /snapsync/entrypoint.sh /snapsync/entrypoint.sh
COPY --from=build /snapsync/config.yml /snapsync/config.yml

WORKDIR /snapsync

ENTRYPOINT [ "./entrypoint.sh" ] 