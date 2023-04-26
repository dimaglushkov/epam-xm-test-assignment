#!/bin/sh
curl -i -H "Content-Type: application/json" -d '{
  "name": "some name 1",
  "random_field": "random_value",
  "asd": true,
  "employee_cnt": 100
}' -X PATCH http://localhost:8080/companies/$1