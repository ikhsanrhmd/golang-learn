FROM golang:latest as builder
WORKDIR /app
#copy file go.mod & go.sum ke directory /app
COPY go.mod go.sum ./ 
RUN go mod download 
COPY . .
RUN CGO_ENABLED=0 go build -o main .
# EXPOSE 80
# ENTRYPOINT ["./main"]

FROM gcr.io/distroless/static-debian11
WORKDIR /app
COPY --from=builder /app/main /app
COPY . . 
EXPOSE 80
CMD ["/app/main"]