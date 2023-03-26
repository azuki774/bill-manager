# BILLMANAGER

## elect_consumption

```
MariaDB [billmanager]> show columns from ElectConsumption;
+-------------------+-----------+------+-----+---------------------+-------------------------------+
| Field             | Type      | Null | Key | Default             | Extra                         |
+-------------------+-----------+------+-----+---------------------+-------------------------------+
| id                | int(11)   | NO   | PRI | NULL                | auto_increment                |
| record_date       | date      | NO   | UNI | NULL                |                               |
| total_comsumption | int(11)   | NO   |     | NULL                |                               |
| day_comsumption   | int(11)   | NO   |     | NULL                |                               |
| night_comsumption | int(11)   | NO   |     | NULL                |                               |
| created_at        | timestamp | NO   |     | current_timestamp() |                               |
| updated_at        | timestamp | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------------+-----------+------+-----+---------------------+-------------------------------+
```
- ただし、record_date として、同じ日付に２つのレコードを登録しない。


## bill_elect
```
+-------------------+------------+------+-----+---------------------+-------------------------------+
| Field             | Type       | Null | Key | Default             | Extra                         |
+-------------------+------------+------+-----+---------------------+-------------------------------+
| id                | int(11)    | NO   | PRI | NULL                | auto_increment                |
| billing_month     | varchar(6) | NO   | UNI | NULL                |                               |
| price             | int(11)    | NO   |     | NULL                |                               |
| total_consumption | int(11)    | NO   |     | NULL                |                               |
| created_at        | timestamp  | NO   |     | current_timestamp() |                               |
| updated_at        | timestamp  | NO   |     | current_timestamp() | on update current_timestamp() |
+-------------------+------------+------+-----+---------------------+-------------------------------+
```

## bill_gas
```
+---------------+------------+------+-----+---------------------+-------------------------------+
| Field         | Type       | Null | Key | Default             | Extra                         |
+---------------+------------+------+-----+---------------------+-------------------------------+
| id            | int(11)    | NO   | PRI | NULL                | auto_increment                |
| billing_month | varchar(6) | NO   | UNI | NULL                |                               |
| price         | int(11)    | NO   |     | NULL                |                               |
| consumption   | int(11)    | NO   |     | NULL                |                               |
| created_at    | timestamp  | NO   |     | current_timestamp() |                               |
| updated_at    | timestamp  | NO   |     | current_timestamp() | on update current_timestamp() |
+---------------+------------+------+-----+---------------------+-------------------------------+
```

## bill_water
```
+--------------------+------------+------+-----+---------------------+-------------------------------+
| Field              | Type       | Null | Key | Default             | Extra                         |
+--------------------+------------+------+-----+---------------------+-------------------------------+
| id                 | int(11)    | NO   | PRI | NULL                | auto_increment                |
| billing_month      | varchar(6) | NO   | UNI | NULL                |                               |
| price              | int(11)    | NO   |     | NULL                |                               |
| consumption        | int(11)    | NO   |     | NULL                |                               |
| detail_water_price | int(11)    | YES  |     | NULL                |                               |
| detail_sewer_price | int(11)    | YES  |     | NULL                |                               |
| created_at         | timestamp  | NO   |     | current_timestamp() |                               |
| updated_at         | timestamp  | NO   |     | current_timestamp() | on update current_timestamp() |
+--------------------+------------+------+-----+---------------------+-------------------------------+
```

