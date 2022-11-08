FROM golang:alpine as builder

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify


COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -trimpath -a -installsuffix cgo -o avdoc cmd/avdoc/main.go

FROM archlinux:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /usr/local/bin/

COPY --from=builder /go/src/app/avdoc .

RUN pacman -Suy --noconfirm && pacman -S --needed --noconfirm git base-devel npm nodejs go lazygit ncdu htop neovim ripgrep sudo

CMD [ "./avdoc" ]
