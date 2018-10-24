# © 2018 Imhotep Software LLC. All materials licensed under Apache v2.0

# ---
# Build Exec
FROM golang:1.11-alpine
LABEL maintainer="fernand@k8sland.io"

RUN apk --no-cache add git curl && \
  mkdir -p /hangman_cli/execs

ADD . /hangman_cli
RUN cd /hangman_cli && \
  go build -ldflags "-w -X main.Build=0.0.1" -a -tags netgo -o execs/hangman_cli \
  cmd/cli/main.go

# ----
# Build Image
FROM   alpine:3.7
LABEL  maintainer="fernand@k8sland.io"

COPY --from=0 /hangman_cli/execs /app
CMD    ["/app/hangman_cli"]