CREATE TABLE programming_langs (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name VARCHAR(20) NOT NULL,
  feature TEXT,
  created_at datetime DEFAULT NULL,
  updated_at datetime DEFAULT NULL,
  PRIMARY KEY (id)
);

ALTER DATABASE sample CHARACTER SET utf8mb4;
ALTER TABLE programming_langs CONVERT TO CHARACTER SET utf8mb4;

