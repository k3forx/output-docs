CREATE TABLE `users` (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL
);

INSERT INTO `users` (name, email) VALUES ("taro", "taro@example.com");
