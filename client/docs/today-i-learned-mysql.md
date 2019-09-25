---
title: Today I Learned MySQL
lang: en-US
tags:
  - mysql
---

# {{ $page.title }}

## Display Output in a Vertical Format

Output for tables with lots of columns can be hard to read and sometimes overflow the terminal window. Consider the output from [Show Indexes For A Table](show-indexes-for-a-table.md):

```sql
> show indexes in users;
+-------+------------+--------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| Table | Non_unique | Key_name     | Seq_in_index | Column_name | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment |
+-------+------------+--------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| users |          0 | PRIMARY      |            1 | id          | A         |           0 |     NULL | NULL   |      | BTREE      |         |               |
| users |          0 | unique_email |            1 | email       | A         |           0 |     NULL | NULL   |      | BTREE      |         |               |
+-------+------------+--------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
```

We can vertically orient the output of a statement by terminating it with `\G` instead of `;` (or `\g`).

```sql
> show indexes in users\G
*************************** 1. row ***************************
        Table: users
   Non_unique: 0
     Key_name: PRIMARY
 Seq_in_index: 1
  Column_name: id
    Collation: A
  Cardinality: 0
     Sub_part: NULL
       Packed: NULL
         Null:
   Index_type: BTREE
      Comment:
Index_comment:
*************************** 2. row ***************************
        Table: users
   Non_unique: 0
     Key_name: unique_email
 Seq_in_index: 1
  Column_name: email
    Collation: A
  Cardinality: 0
     Sub_part: NULL
       Packed: NULL
         Null:
   Index_type: BTREE
      Comment:
Index_comment:
```

## Dump a Database to a File

The `mysqldump` client is a handy tool for creating a backup or snapshot of a MySQL database. The standard use of this command produces an alphabetical series of statements that comprise the structure and data of the specified database. It directs all of this to `stdout`. You'll likely want to redirect it to a file.

```bash
$ mysqldump my_database > my_database_backup.sql
```

The output will include special comments with MySQL directives that disable things like constraint checking. This is what allows the output to be in alphabetical order without necessarily violating any foreign key constraints.

If you need to dump multiple databases, include the `--databases` flag with a space-separated list of database names. Or dump all of them with `--all-databases`.

See `man mysqldump` for more details.

## List Databases and Tables

If you've started a [mysql](https://dev.mysql.com/) session, but haven't connected to a particular database yet, you can list the available databases like so:

```sql
> show databases;
+-----------------------------+
| Database                    |
+-----------------------------+
| information_schema          |
| my_app_dev                  |
+-----------------------------+
```

If you are curious about the tables in a particular database, you can list them by specifying the database's name:

```sql
> show tables in my_app_dev;
+------------------------------+
| Tables_in_my_app_dev         |
+------------------------------+
| pokemons                     |
| trainers                     |
+------------------------------+
```

Alternatively, you can connect to the database of interest and then there is no need to specify the name of the database going forward.

```sql
> use my_app_dev;
> show tables;
+------------------------------+
| Tables_in_my_app_dev         |
+------------------------------+
| pokemons                     |
| trainers                     |
+------------------------------+
```

## Show Indexes for a Table

When describing a table, such as the table `users`:

```sql
> describe users;
+------------+-----------------------+------+-----+---------+----------------+
| Field      | Type                  | Null | Key | Default | Extra          |
+------------+-----------------------+------+-----+---------+----------------+
| id         | mediumint(8) unsigned | NO   | PRI | NULL    | auto_increment |
| first_name | varchar(80)           | NO   |     | NULL    |                |
| last_name  | varchar(80)           | NO   |     | NULL    |                |
| email      | varchar(80)           | NO   | UNI | NULL    |                |
+------------+-----------------------+------+-----+---------+----------------+
```

I can see in the `Key` column that there is a primary key and a unique key for this table on `id` and `email`, respectively.

These keys are indexes. To get more details about each of the indexes on this table, we can use the [`show indexes`](https://dev.mysql.com/doc/refman/5.7/en/show-index.html)
command.

```sql
> show indexes in users;
+-------+------------+--------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| Table | Non_unique | Key_name     | Seq_in_index | Column_name | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment |
+-------+------------+--------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| users |          0 | PRIMARY      |            1 | id          | A         |           0 |     NULL | NULL   |      | BTREE      |         |               |
| users |          0 | unique_email |            1 | email       | A         |           0 |     NULL | NULL   |      | BTREE      |         |               |
+-------+------------+--------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
```

## Show Tables that Match a Pattern

An unfamiliar database with tons of tables can be a difficult thing to navigate. You may have an idea of the kind of table you are looking for based on a domain concept you've seen elsewhere.

You can pare down the results returned by `show tables` by including a `like` clause with a pattern. For example, this statement will show me only tables that have the word `user` in them:

```sql
> show tables like '%user%';
+-------------------------------+
| Tables_in_jbranchaud (%user%) |
+-------------------------------+
| admin_users                   |
| users                         |
+-------------------------------+
```
