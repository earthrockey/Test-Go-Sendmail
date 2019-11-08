FROM golang

WORKDIR /go/src/github.com/earthrockey/Test-Go-Sendmail
COPY . .
RUn go get github.com/line/line-bot-sdk-go/linebot
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]