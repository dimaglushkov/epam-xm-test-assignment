curl -i -H "Content-Type: application/json" -d '{
  "name": "some name",
  "description": "some description",
  "employee_cnt": 10,
  "registered": true,
  "type": "Corporations"
}' -X POST http://localhost:8080/company