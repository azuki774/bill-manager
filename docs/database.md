# BILLMANAGER

## elect_consumes

```
MariaDB [BILLMANAGER]> show columns from elect_consumes;
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
- ただし、record_date として、同じ日付に２つのレコードを登録しない。
