-- +migrate Up
CREATE TABLE `elect_consumption` (
  `id` int NOT NULL AUTO_INCREMENT,
  `record_date` date UNIQUE NOT NULL,
  `total_comsumption` int NOT NULL,
  `day_comsumption` int NOT NULL,
  `night_comsumption` int NOT NULL,
  `created_at` timestamp NOT NULL default current_timestamp,
  `updated_at` timestamp default current_timestamp on update current_timestamp,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx1` (`record_date`)
);

-- +migrate Down
DROP TABLE `elect_consumption`;
