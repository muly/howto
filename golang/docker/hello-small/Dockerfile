FROM golang:latest as BuildStage
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM scratch
WORKDIR /api
COPY --from=BuildStage /build/app /api/
CMD ["./app"]