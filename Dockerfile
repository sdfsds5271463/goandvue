# ====== Build Frontend ======
FROM node:20 AS frontend-builder

WORKDIR /app/frontend

COPY frontend/package*.json ./
RUN npm install

COPY frontend .
RUN npm run build


# ====== Build Backend ======
FROM golang:1.25-alpine AS backend-builder

WORKDIR /app/backend

COPY backend/go.mod ./
RUN go mod download

COPY backend .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app


# ====== Final Image ======
FROM alpine:latest

WORKDIR /app

COPY --from=backend-builder /app/backend/app .
COPY --from=frontend-builder /app/frontend/dist ./static

EXPOSE 8080

CMD ["./app"]