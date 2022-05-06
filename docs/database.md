# BILLMANAGER

## elect_consume

```
MariaDB [BILLMANAGER]> show columns from elect_consume;
+-------------+---------+------+-----+---------------------+----------------+
| Field       | Type    | Null | Key | Default             | Extra          |
+-------------+---------+------+-----+---------------------+----------------+
| id          | int(11) | NO   | PRI | NULL                | auto_increment |
| record_date | date    | NO   |     | NULL                |                |
| created_at  | date    | YES  |     | current_timestamp() |                |
| updated_at  | date    | YES  |     | current_timestamp() |                |
| daytime     | int(11) | YES  |     | NULL                |                |
| nighttime   | int(11) | YES  |     | NULL                |                |
| total       | int(11) | YES  |     | NULL                |                |
+-------------+---------+------+-----+---------------------+----------------+
```
