FROM golang:1.21

WORKDIR /go/src

RUN go get -u github.com/spf13/cobra@latest && \
    go install github.com/golang/mock/mockgen@latest && \
    go install github.com/spf13/cobra-cli@latest


RUN apt-get update && apt-get install sqlite3 -y

RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache

COPY go.mod go.sum ./
RUN go mod download

COPY . .

USER www-data

CMD ["go", "run", "./cmd/main.go"]
