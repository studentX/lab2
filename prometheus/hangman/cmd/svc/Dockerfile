# © 2018 Imhotep Software LLC. All materials licensed under Apache v2.0

# ---
# Build Exec
FROM  golang:1.11-alpine
LABEL maintainer="fernand@k8sland.io"

RUN apk --no-cache add git curl && \
  mkdir -p /hangman/execs

ADD . /hangman
RUN cd /hangman && \
  go build -ldflags "-w -X main.Build=0.0.1" -a -tags netgo -o execs/hangman \
  cmd/svc/main.go

# ----
# Build image
FROM   alpine:3.7
LABEL  maintainer="fernand@k8sland.io"

COPY --from=0 /hangman/execs /app

EXPOSE 5000
CMD    ["/app/hangman", "-d", "dictionary:4000"]