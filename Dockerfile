# GitHub:       https://github.com/gohugoio
# Twitter:      https://twitter.com/gohugoio
# Website:      https://gohugo.io/

#FROM golang:1.19-alpine AS build

# Optionally set HUGO_BUILD_TAGS to "extended" or "nodeploy" when building like so:
#   docker build --build-arg HUGO_BUILD_TAGS=extended .
#ARG HUGO_BUILD_TAGS

#ARG CGO=1
#ENV CGO_ENABLED=${CGO}
#ENV GOOS=linux
#ENV GO111MODULE=on

#WORKDIR /go/src/github.com/gohugoio/hugo

#COPY . /go/src/github.com/gohugoio/hugo/

# gcc/g++ are required to build SASS libraries for extended version
#RUN apk update && \
#    apk add --no-cache gcc g++ musl-dev git && \
#    go install github.com/magefile/mage

#RUN mage hugo && mage install

# ---

#FROM alpine:3.16

#COPY ./monegoo /usr/bin/monegoo

# libc6-compat & libstdc++ are required for extended SASS libraries
# ca-certificates are required to fetch outside resources (like Twitter oEmbeds)
#RUN apk update && \
#    apk add --no-cache ca-certificates libc6-compat libstdc++ git

#VOLUME /testdata
#WORKDIR /

# Expose port for live server
#EXPOSE 8080

#ENTRYPOINT ["monegoo"]
#CMD ["--help"]



FROM golang:1.19-alpine AS build
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN go build monegoo


FROM alpine:3.16
WORKDIR /

COPY --from=build /app/monegoo /usr/bin/monegoo
COPY --from=build /app/testdata /testdata

EXPOSE 8080
ENTRYPOINT monegoo
