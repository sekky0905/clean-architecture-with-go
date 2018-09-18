CREATE DATABASE IF NOT EXISTS sample;
USE sample;

CREATE TABLE programming_languages (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name VARCHAR(20) NOT NULL,
  feature TEXT,
  created_at datetime DEFAULT NULL,
  updated_at datetime DEFAULT NULL,
  PRIMARY KEY (id)
);

insert into programming_languages(name, feature) values("Go", "静的型付言語。並行処理が書きやすい。");
insert into programming_languages(name, feature) values("java", "静的型付言語。エンタープライズシステムでの利用例が多い。");
insert into programming_languages(name, feature) values("Pyhon", "動的型付け言語。機械学習での利用例が多い。");