-- Assuming that company types list is not a target of frequent updates,
-- I'd prefer to use enum, but the same logic could be achieved using an external table.
CREATE TYPE company_type AS ENUM('Corporations', 'NonProfit', 'Cooperative', 'Sole Proprietorship');

CREATE TABLE IF NOT EXISTS company (
    id UUID PRIMARY KEY,
    name VARCHAR(15) NOT NULL UNIQUE,
    description VARCHAR(3000),
    employee_cnt INT NOT NULL,
    registered boolean NOT NULL,
    type company_type NOT NULL
)