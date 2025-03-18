-- migrate:up
CREATE TABLE ads
(
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  text TEXT NOT NULL,
  url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO ads (name, text, url) VALUES ('hair', 'Hairdryer for sale. 50% off.', '/product/2ZYFJ3GM2N');
INSERT INTO ads (name, text, url) VALUES ('clothing', 'Tank top for sale. 20% off.', '/product/66VCHSJNUP');
INSERT INTO ads (name, text, url) VALUES ('accessories', 'Watch for sale. Buy one, get second kit for free', '/product/1YMWWN1N4O');
INSERT INTO ads (name, text, url) VALUES ('footwear', 'Loafers for sale. Buy one, get second one for free', '/product/L9ECAV7KIM');
INSERT INTO ads (name, text, url) VALUES ('decor', 'Candle holder for sale. 30% off.', '/product/0PUK6V6EV0');
INSERT INTO ads (name, text, url) VALUES ('kitchen', 'Bamboo glass jar for sale. 10% off.', '/product/9SIQT8TOJO');

-- migrate:down

