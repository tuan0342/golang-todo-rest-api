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

docker run --name pgdb \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=123456 \
  -e POSTGRES_DB=mydb \
  -p 5432:5432 \
  -d postgres:latest
```
