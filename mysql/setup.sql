CREATE TABLE programming_languages (
  id INT unsigned NOT NULL, AUTO_INCREMENT,
  name VARCHAR(20) NOT NULL,
  feature VARCHAR TEXT,
  created_at datetime DEFAULT NULL,
  updated_at datetime DEFAULT NULL,
  PRIMARY KEY (id)
)

inert into programming_languages(name, feature) values ("Go", "静的型付言語。並列処理が容易。読みやすい。");
inert into programming_languages(name, feature) values ("java", "静的型付言語。エンタープライズシステムでの利用例が多い。");
inert into programming_languages(name, feature) values ("Pyhon", "動的型付け言語。機械学習での利用例が多い。");