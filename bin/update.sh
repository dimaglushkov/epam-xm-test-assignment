#!/bin/sh
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.sx0g3qGuN7yCndCZVXpLO3K6rNzO4morT9Da-xwwzts"

curl -i -H "Content-Type: application/json" -d '{
  "name": "some name 1",
  "random_field": "random_value",
  "asd": true,
  "employee_cnt": 100
}' -H "Authorization: Bearer $TOKEN" -X PATCH http://localhost:8080/companies/$1