FROM golang AS build
COPY ./go.* /src/
WORKDIR /src
RUN go mod download

COPY . /src

RUN CGO_ENABLED=0 go build -o /gmon

# Build the final image
FROM alpine
COPY --from=build /gmon /usr/bin/

ENTRYPOINT [ "gmon" ]
CMD [ "--config=/data/config.toml", "--rules=/data/rules.toml" ]
