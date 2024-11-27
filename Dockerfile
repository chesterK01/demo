# Sử dụng hình ảnh Golang chính thức
FROM golang:1.20

# Tạo thư mục làm việc
WORKDIR /app

# Sao chép module và mã nguồn
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Biên dịch ứng dụng
RUN go build -o main .

# Chạy ứng dụng
CMD ["./main"]