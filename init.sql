CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang';
GRANT ALL PRIVILEGES ON testdb.* to 'golang'@'localhost';
CREATE DATABASE testdb;
