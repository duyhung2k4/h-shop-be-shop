FROM golang:1.22

WORKDIR /usr/local/bin

COPY . .

RUN go mod download
RUN go build -o main .

CMD [ "./main" ]