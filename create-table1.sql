DROP TABLE IF EXISTS user_info;
CREATE TABLE user_info (
  id         INT AUTO_INCREMENT NOT NULL,
  tg_nickname      VARCHAR(255) NOT NULL,
  balance     FLOAT(32) NOT NULL,

  PRIMARY KEY (`id`)
);

INSERT INTO user_info
  (id, tg_nickname, balance)
VALUES
(000,'TG',999999999);