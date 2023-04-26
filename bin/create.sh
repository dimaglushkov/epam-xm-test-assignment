#!/bin/sh
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.sx0g3qGuN7yCndCZVXpLO3K6rNzO4morT9Da-xwwzts"

curl -i -H "Content-Type: application/json" -d '{
  "name": "some name",
  "description": "some description",
  "employee_cnt": 10,
  "registered": false,
  "type": "NonProfit"
}' -H "Authorization: Bearer $TOKEN" -X POST http://localhost:8080/companies