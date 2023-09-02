FROM node:lts-alpine as node
ARG TARGETOS
ARG TARGETARCH

RUN apk add g++ make py3-pip

WORKDIR /frontend

COPY package.json package-lock.json .postcssrc tailwind.config.js ./
COPY web web

RUN npm install
RUN npm run build

FROM golang:1.18-alpine as golang

WORKDIR /app
COPY . .
COPY --from=node /frontend .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /server cmd/likeit-server/main.go

FROM gcr.io/distroless/static-debian11

COPY --from=golang /server .

EXPOSE 8080

CMD ["/server"]
