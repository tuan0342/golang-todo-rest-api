# Todo App (Golang + Gin + Gorm + PostgreSQL)

Đây là ứng dụng Todo đơn giản được viết bằng Golang, sử dụng:

- **Gin**: Web framework
- **Gorm**: ORM để làm việc với PostgreSQL
- **Docker**: Chạy database PostgreSQL

## Cài đặt

### 1. Clone project

```bash
git clone git@github.com:<username>/<repo>.git
cd <repo>
```

### 2. Chạy postgres docker
```bash
docker run --name pgdb \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=123456 \
  -e POSTGRES_DB=mydb \
  -p 5432:5432 \
  -d postgres:latest
```

### 3. Các lệnh docker compose liên quan
```bash
/// 1. Start backend and postgres
docker compose up

/// 2. Remove images
docker compose down

/// 3. Build lại images khi sửa code
docker compose --build --no-cache
```

### 4. Các câu lệnh terminal của golang thường gặp
#### 4.1. Kiểm tra và tải dependency
```bash
go mod download
```
Lệnh này: tải tất cả dependency từ go.mod về local máy (không phải tải về project)

Sau đó, để đồng bộ lại module (nếu team dùng Go version khác):
```bash
go mod tidy
```
Tác dụng:
- xóa dependency không dùng
- thêm dependency còn thiếu
- update go.sum

#### 4.2. Vendor dependency
Một số project yêu cầu vendor/:
```bash
go mod vendor
```
Dùng khi:
- build trong Docker
- build offline
- dùng CI/CD cần vendor
- muốn đảm bảo dependency cố định
Nếu dự án không dùng vendor → bạn bỏ bước này.

#### 4.2. Chạy dự án
```bash
go run .
```
hoặc
```bash
go run -mod=vendor .
```