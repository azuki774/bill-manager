-- +migrate Up
CREATE TABLE `bill_water` (
  `id` int NOT NULL AUTO_INCREMENT,
  `billing_month` varchar(6) UNIQUE NOT NULL,
  `price` int NOT NULL,
  `consumption` int NOT NULL,
  `detail_water_price` int,
  `detail_sewer_price` int,
  `created_at` timestamp NOT NULL default current_timestamp,
  `updated_at` timestamp default current_timestamp on update current_timestamp,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx1` (`billing_month`)
);

-- +migrate Down
DROP TABLE `bill_water`;
