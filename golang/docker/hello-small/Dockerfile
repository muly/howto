FROM golang:onbuild as BuildStage
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
ADD app /


FROM scratch
WORKDIR /app
COPY --from=BuildStage app /app/
CMD ["./app"]