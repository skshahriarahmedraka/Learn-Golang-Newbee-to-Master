

# day23 study notes

## One, the database

### 1.1 Basic knowledge of database

DB:

DBMS:

The structure of databases, data tables, and tables. .

DB: refers to datebase (database)
	A database is a collection of stored data. The database is usually composed of data tables, etc., and the data table is composed of information such as data fields and data values.
DBMS: refers to datebase mangement systerm (database management system)
	It is a system for operating and managing databases. For example, mysql, sqlserver, etc. belong to database management software. People use these systems or tools to manage the data in the database.
DBS: refers to datebase systerm (database system)
	The database system is composed of databases and database management software. The database is a logical concept of storing data, and the corresponding entity is the database stored on the hard disk by the database management software, so the database system includes the database and the database management software.

### 1.2 Mysql installation and uninstallation





### 1.3 Login

Method 1: DOS window: enter the following command:

```shell
C:\Users\ruby>mysql -u root -p
Enter the password after pressing enter
```

![WX20190725-112248](img\WX20190725-112248.png)





Method 2: Login via Mysql Command Line:

```
Just enter the password directly
```

![WX20190725-112402](img\WX20190725-112402.png)

Method 3: Through other visualization tool software:

![WX20190725-112650](img\WX20190725-112650.png)



















### 1.4 Create a database:

1. Create a database:

```sql
//create database [if not exists] database name [default charset utf8 collate utf8_general_ci];
mysql> create database my1905 character set utf8;
Query OK, 1 row affected (0.00 sec)


```

2. Show which databases are available:

```mysql
mysql> show databases;
```



![WX20190725-120339](img\WX20190725-120339.png)



3. Switch to the database: the subsequent operations are for the database, such as building a table. .

```mysql
mysql> use my1905;
```



4. Check what data tables are in the current database:

```mysql
mysql> show tables;
```



![WX20190725-120545](img\WX20190725-120545.png)



5. Delete the database:

```mysql
mysql> drop database if exists my1905;
```



### 1.5 Data Type



char(10)-->fixed-length string

​ "wangergou "

​ "abc "

varchar(10)-->Variable length

​ "wangergou"

​ "abc"







### 1.6 Data Table Operation

1. Create a database:

```mysql
mysql> create database if not exists my1905 default charset utf8 collate utf8_ge
neral_ci;
```



2. Create a data table:

```mysql
mysql> create table users(
    -> id int(4) primary key auto_increment,
    -> username varchar(20),
    -> pwd varchar(30));
```



3. View the table structure: desc-->describe

```mysql
mysql> desc users;
```



![WX20190725-143324](img\WX20190725-143324.png)



4. Display check table statement:

```mysql
mysql> show create table users;
```

![WX20190725-143609](img\WX20190725-143609.png)



be careful:

1. Create the database first

mysql:

​ database1-->oa

​ database2-->bluebird

​. . . .



2. Switch database

​ use database name



3. Create a data table

mysql>create table test1(

​ ->id int(4) auto_increment primary key,

​ ->...);





5. Insert a piece of data:

```mysql
mysql> insert into users(id,username,pwd) values(1,'admin','123456');
Query OK, 1 row affected (0.02 sec)
```



6. Query data:

```mysql
mysql> select * from users;
+----+----------+--------+
| id | username | pwd |
+----+----------+--------+
| 1 | admin | 123456 |
+----+----------+--------+
1 row in set (0.00 sec)
```

### 1.7 Modify the table structure

alter table table name xxx. . .

1. Add a field: add

   ```mysql
   mysql> alter table users add(
       -> age int(4),
       -> birthday date);
   ```

   ![WX20190725-152709](img\WX20190725-152709.png)

   

2. Modify the data type of an existing field: modify

   ```mysql
   mysql> alter table users modify age float(4,1);
   ```

   

![WX20190725-153114](img\WX20190725-153114.png)



Note: The data type of an existing column cannot be changed arbitrarily. Especially if there is already data in the table

​ A: Compatible type: The length can be from small to large, and the existing data cannot exceed the boundary.

​ B: Incompatible type: varchar-->int, change failed.



3. Change the name of the column: change

```mysql
mysql> alter table users change pwd password varchar(30);
```

![WX20190725-153525](img\WX20190725-153525.png)

​	

4. Delete a column: drop

   ```mysql
   mysql> alter table users drop birthday;
   ```

![WX20190725-153854](img\WX20190725-153854.png)



If there is data in this column, the data will also be deleted.

5. Table rename: rename to

```mysql
mysql> alter table users rename to user2;
mysql> rename table user2 to user3;
```

![WX20190725-154455](img\WX20190725-154455.png)



6. Delete the table: drop table

```mysql
mysql> drop table user3;
```





### 1.8 Insert data

1. Insert data:

```mysql
insert into table name (column 1, column 2, column 3....) values ​​(value 1, value 2, value 3....)
```



Full column insertion: If you want to insert data in all columns, you can omit the column name

Default insertion: If there is no value in one or some fields, then the column name and value must be clearly written.

Insert multiple rows at the same time:





### 1.9 Modify data

Grammatical structures:

```mysql
update table name set column 1=value 1, column 2=value 2...[where condition];
```

After where is the modification condition: if it is true, the data will be modified.

```
Operator:
	=, the value is equal
	!=, <>, unequal values
	between ... and, interval
	>
	<
	>=
	<=
	or
	and
	in(value 1, value 2, value 3..)



```

1. Modify the name of the student whose student ID is 1006 to Chen Cong

```mysql
mysql> update student set name='Chen Cong' where no=1006;
Query OK, 1 row affected (0.00 sec)
Rows matched: 1 Changed: 1 Warnings: 0

mysql> select * from student;
+------+--------+------+------+------------+
| no | name | age | sex | birthday |
+------+--------+------+------+------------+
| 1001 | King Two Dogs | 18 | Male | 2007-10-10 |
| 1002 | rose | 19 | Female | 2006-09-09 |
| 1003 | jack | 20 | male | 2005-08-06 |
| 1004 | Zhang San | 18 | Female | 1990-12-12 |
| 1005 | Li Si | 21 | Male | 1991-06-08 |
| 1006 | Chen Cong | 22 | Male | 1992-10-10 |
+------+--------+------+------+------------+
6 rows in set (0.00 sec)
```



2. For students younger than 19 years old, the gender is changed to female

```mysql
mysql> update student set sex='female' where age <19;
Query OK, 1 row affected (0.01 sec)
Rows matched: 2 Changed: 1 Warnings: 0

mysql> select * from student;
+------+--------+------+------+------------+
| no | name | age | sex | birthday |
+------+--------+------+------+------------+
| 1001 | King Two Dogs | 18 | Female | 2007-10-10 |
| 1002 | rose | 19 | Female | 2006-09-09 |
| 1003 | jack | 20 | male | 2005-08-06 |
| 1004 | Zhang San | 18 | Female | 1990-12-12 |
| 1005 | Li Si | 21 | Male | 1991-06-08 |
| 1006 | Chen Cong | 22 | Male | 1992-10-10 |
+------+--------+------+------+------------+
6 rows in set (0.01 sec)
```

3. The name of the classmate who is 18 years or older and 19 years old or less is changed to Ma Dongmei

```mysql
mysql> update student set name='Ma Dongmei' where age >= 18 and age <= 19;
Query OK, 3 rows affected (0.01 sec)
Rows matched: 3 Changed: 3 Warnings: 0

mysql> select *from student;
+------+--------+------+------+------------+
| no | name | age | sex | birthday |
+------+--------+------+------+------------+
| 1001 | Ma Dongmei | 18 | Female | 2007-10-10 |
| 1002 | Ma Dongmei | 19 | Female | 2006-09-09 |
| 1003 | jack | 20 | male | 2005-08-06 |
| 1004 | Ma Dongmei | 18 | Female | 1990-12-12 |
| 1005 | Li Si | 21 | Male | 1991-06-08 |
| 1006 | Chen Cong | 22 | Male | 1992-10-10 |
+------+--------+------+------+------------+
6 rows in set (0.00 sec)
```



4. Modify the name of the classmate between the ages of 19 and 20 to Ma Chunmei:

```mysql
mysql> update student set name='Ma Chunmei' where age between 19 and 20;
Query OK, 2 rows affected (0.01 sec)
Rows matched: 2 Changed: 2 Warnings: 0

mysql> select * from student;
+------+--------+------+------+------------+
| no | name | age | sex | birthday |
+------+--------+------+------+------------+
| 1001 | Ma Dongmei | 18 | Female | 2007-10-10 |
| 1002 | Ma Chunmei | 19 | Female | 2006-09-09 |
| 1003 | Ma Chunmei | 20 | Male | 2005-08-06 |
| 1004 | Ma Dongmei | 18 | Female | 1990-12-12 |
| 1005 | Li Si | 21 | Male | 1991-06-08 |
| 1006 | Chen Cong | 22 | Male | 1992-10-10 |
+------+--------+------+------+------------+
6 rows in set (0.00 sec)
```







## Two, SQL

Structured Query Language (Structured Query Language). Operating the database.

DDL language: Data Definition Language (used to define the table structure of the data) Data Definition Language

​ Create a data table: create table table name

​ Modify the data table: alter table table name 

​ Delete data table: drop table table name

DML language: data manipulation language (used to manipulate the data in the data table) DML-Data Mainpulation Language

​ Add data: insert 

​ Modify data: update

​ Delete data: delete

DQL language: Data Query Language (specially used for data query) DQL-Data Query Language

​ Query data: select

DCL language:



## Three, summary



database:

​ Installation and uninstallation (see documentation)

​ Login to the database:

​ 1.dos window: mysql command ---> configure environment variables

​ -u username

​ -p password

​ 2. mysql command line: directly enter the password

​ 3. Through some visualization tools: such as Navicat







1.show databases;

2.create database if not exists my1905 character set utf8;

​ default charset utf8 collate utf8_general_ci;

3.use my1905;

4. create table student(id int(4) primary key auto_increment, name varchar(30), sex varchar(2));

5.alter table table name

​ add column name data type

​ modify column name data type

​ change The original column name and the new column name data type

​ drop delete column

6.drop table table name;

7. Insert into table name (column 1, column 2, column 3....) values ​​(value 1, value 2, value 3....)

​ Insert all columns:

​ Insert multiple at the same time:

8.update table name set column 1=new value, column 2=new value [where to modify conditions];

​ The expression after where is boolean

​ =,!=,<>,>,<,>=,<=,between and, and, or ,not ....

​ null---> is null ,is not null

9.delete from table name where delete conditions









Constraints: primary key, foreign key

Query: simple query, complex, multiple tables

