FROM golang AS build
RUN go get -u github.com/hamburghammer/gsave

ENV CGO_ENABLED=0
RUN go build -o /gsave github.com/hamburghammer/gsave

FROM alpine 
COPY --from=build /gsave /usr/bin/

ENTRYPOINT [ "gsave" ]
CMD [ "--verbose" ]
EXPOSE 8080