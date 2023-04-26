#!/bin/sh
curl -i -H "Content-Type: application/json" -d '{
  "name": "some name",
  "description": "some description",
  "employee_cnt": 10,
  "registered": false,
  "type": "NonProfit"
}' -X POST http://localhost:8080/companies