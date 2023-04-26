#!/bin/sh
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.sx0g3qGuN7yCndCZVXpLO3K6rNzO4morT9Da-xwwzts"

curl -i  -H "Authorization: Bearer $TOKEN" -X  DELETE http://localhost:8080/companies/$1