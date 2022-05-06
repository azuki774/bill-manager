CREATE DATABASE IF NOT EXISTS BILLMANAGER DEFAULT CHARACTER SET utf8mb4;
grant all privileges on *.* to root@"%";
USE BILLMANAGER
CREATE TABLE IF NOT EXISTS elect_consume (
    id INT AUTO_INCREMENT,
    record_date DATE NOT NULL,
    created_at DATE DEFAULT CURRENT_TIMESTAMP,
    updated_at DATE DEFAULT CURRENT_TIMESTAMP,
    daytime INT,
    nighttime INT,
    total INT,
    PRIMARY KEY (`id`)
);
