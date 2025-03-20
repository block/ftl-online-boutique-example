-- migrate:up
CREATE TABLE product (
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  picture TEXT,
  price_usd DECIMAL(10, 2) NOT NULL,
  categories TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- migrate:down

