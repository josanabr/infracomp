FROM golang:alpine AS build
WORKDIR /go/src/proyecto
COPY proyecto.go .
RUN env GOARCH=386 GOOS=linux go build -o /go/bin/proyecto proyecto.go

# 
FROM scratch
COPY --from=build /go/bin/proyecto /go/bin/proyecto
EXPOSE 8000
ENTRYPOINT [ "/go/bin/proyecto" ]
