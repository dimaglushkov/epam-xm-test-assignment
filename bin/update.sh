#!/bin/sh
curl -i -H "Content-Type: application/json" -d '{
  "name": "some name 4",
  "random_field": "random_value",
  "asd": true,
  "employee_cnt": 100
}' -X PATCH http://localhost:8080/companies/5c89ab81-ccf7-4918-9873-1fba08451303