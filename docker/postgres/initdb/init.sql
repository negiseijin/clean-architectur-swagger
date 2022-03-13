--ユーザーの作成
CREATE USER db_user;
--DBの作成
CREATE DATABASE mydb;
--ユーザーにDBの権限をまとめて付与
GRANT ALL PRIVILEGES ON DATABASE mydb TO db_user;
--DBを切り替え
\c mydb
--テーブルを作成
-- https://www.postgresql.jp/document/13/html/tutorial-table.html
CREATE TABLE weather (
    city            varchar(80),
    temp_lo         int,           -- 最低気温
    temp_hi         int,           -- 最高気温
    prcp            real,          -- 降水量
    date            date
);
CREATE TABLE cities (
    name            varchar(80),
    location        point
);
--テーブルにデータを挿入
INSERT INTO weather VALUES ('San Francisco', 46, 50, 0.25, '1994-11-27');
INSERT INTO cities VALUES ('San Francisco', '(-194.0, 53.0)');
INSERT INTO weather (city, temp_lo, temp_hi, prcp, date)
    VALUES ('San Francisco', 43, 57, 0.0, '1994-11-29');
INSERT INTO weather (date, city, temp_hi, temp_lo)
    VALUES ('1994-11-29', 'Hayward', 54, 37);
--ビューを作成
CREATE VIEW myview AS
    SELECT city, temp_lo, temp_hi, prcp, date, location
        FROM weather, cities
        WHERE city = name;
