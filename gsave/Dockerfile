FROM golang AS build
ENV GO111MODULE="on"
ENV CGO_ENABLED=0
RUN go get -u github.com/hamburghammer/gsave@d0802fd3d77160b61cb2303b92642b01f2357d57

FROM alpine 
COPY --from=build /go/bin/gsave /usr/bin/

ENTRYPOINT [ "gsave" ]
CMD [ "--verbose" ]
EXPOSE 8080