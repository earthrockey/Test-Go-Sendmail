FROM golang

WORKDIR /go/src/github.com/earthrockey/Test-Go-Sendmail
COPY . .
RUN go get github.com/emersion/go-smtp
RUN go get github.com/emersion/go-sasl
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]