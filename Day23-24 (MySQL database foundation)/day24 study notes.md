

# day15 Classroom materials

## Last class review

mysql database

Log in:

​ dos window: mysql

​ -u username

​ -p password

​ The command line that comes with mysql:

​ Enter the password directly

​	

​ Visualization tools:



Basic knowledge of the database:

​ Relational database: organize data in rows and columns. (Two-dimensional table)

After installing mysql:

​ Create database: (by project)

​ Create a data table: table

​ Row: a piece of data, or a record

​ Column: field, attribute, field



SQL statement:

​ DDL: Data Definition Language

​ create database

​ create table

​ alter talble

​ drop table







1. Create a database:

2. Show all databases: show databases;

3. Switch database: use database name;

4. Create a data table:

​ id: int, primary key (non-empty + unique) --> unique identifier, self-increment: auto_increment

​ Data type:

​ Numerical type: int (length), float, double. . .

​ Character type: char, varchar

​ Date type: date, time, datetime



5. Insert data:

​ insert into table name (field...) values ​​(number...); 

6. Modify the table structure:

​ alter table table name add/modify/change/drop 

​ rename to

7. Delete the data table:

​ drop table table name;

8. Modify the data:

​ update table name set field = value, field = value [where modify conditions]

```
=,!=,<>
>,<,>=,<=
or,and,not
between...and, in()
null is null/is not null
```



9. Delete data:

​ delete from table name [where delete condition];



## One. Constraints

Constraint: Used to limit the storage content of a column in the data table.

​ Non-empty constraint, unique constraint

Primary key constraint: non-empty + unique

​ Used in this table, the field where the primary key is located is the unique identifier of the table.

Foreign key constraints: to ensure the integrity and validity of the data.

​ Two tables:

​ Parent table: main table

​ Primary key

​ Sub-table: From the table

​ Foreign key



The column for setting the foreign key in the child table is the primary key in the parent table. Then the value of the foreign key column in the child table will be constrained by the value of the primary key in the parent table.

Create the parent table:

```Mysql
mysql> create table class(
    -> classno int(4) primary key,
    -> classname varchar(20),
    -> local varchar(30));
```



Create a child table:

```mysql
mysql> create table student(
    -> sid int(4) primary key auto_increment,
    -> sname varchar(30),
    -> age int(3),
    -> sex varchar(3),
    -> classno int(4),
    -> constraint fk_stu foreign key (classno) references class(classno));
```



adding data:

​ Parent table: 1, 2, 3

​ Child table: classno: restricted by the parent table



Cascading operation: When the data of the parent table is deleted, what value is referenced in the child table? on delete xxx

​ Default:

​ cascade: The data of the child table is deleted along with the parent table

​ set null: The parent table is deleted, and the child table is set to null

​ no action:



Delete the original foreign key constraint in the student

```mysql
alter table student drop foreign key fk_stu;
```

New foreign key constraints:

```mysql
mysql> alter table student add constraint fk_stu foreign key(classno) references class(classno) on delete cascade;
```



The data in the parent table is deleted, and the child table is also deleted.





## Second, query

### 2.1 Simple query

1. Query the specified column:

```mysql
select column 1, column 2, column 3. . from table name
```

2. From alias: as can be omitted without writing

```mysql
select column 1 as alias from table name
```

3. Handling null

Use ifnull(a,b), if column a is empty, take the value of b.

```mysql
select ifnull(comm,0) from emp;
```

4. Deduplication: distinct

```mysql
select distinct job from emp;
```



### 2.2 Condition query

When retrieving data in the database, certain conditions must be met before it can be retrieved. Use the where keyword to limit the retrieval conditions.

```mysql
Comparison operators: =,!+=,<>,>,<,>=,<=

Logical operators: and, or, not

Range: between and, in

null: is null, is not null

```



Exercise 1: Query information about employees who joined the company after 1981

Exercise 2: Query the information of employees whose department number is 30 or whose salary is greater than 2000.



**Fuzzy query: like**

%: match 0-many arbitrary characters

_: match 1 arbitrary character

```mysql
//Employee information whose name is a in the third letter
mysql> select * from emp where ename like'__a%';
```



**Sort**: orderby

asc: ascending order, default

desc: descending order

After the select query is completed, the sort should be written at the end of the entire sql statement.



### 2.3 Statistical functions

Also called aggregate function, it is usually used to find the data of a certain column in the entire table: sum, average, maximum, and minimum. It is usually not queried together with the fields in the table.

sum(),

avg(),

max()

min(),

count(*/primary key)



Exercise 1: Find the average salary, total salary, maximum salary, minimum salary, and number of employees in department 20.

mysql> select ename,sum(sal),avg(sal),max(sal),min(sal),count(empno),count(comm)from emp where deptno=20;



### 2.4 Grouping

group by column

Group according to a certain column. If the column has several values, it is divided into several groups.

Exercise 1: Group by department, query the maximum wage, minimum wage, and average wage of each department.

Exercise 2: Find the highest salary, lowest salary, and number of people for each job.

Exercise 3: Query the department with more than 5 people in the department.



### 2.5 Paging

limit start, count.

Exercise: Sort by salary and get the top 5 data.





## Three, multi-meter joint check

### 3.1 Internal connection

The data queried must meet the rules of the link.

```mysql
select e.*,d.* from emp e inner join dept d on e.deptno=d.deptno;
```



### 3.2 Left External Link

Because the query result of the inner connection is not all the data, but the data that meets the rules.

The left outer link and the right outer link are to supplement the query results of the inner link.

The records in the left table will be queried regardless of whether the conditions are met, and the right table can be queried only if the conditions are met. The records in the left table that do not meet the conditions, the right table part is NULL

```mysql
mysql> select * from emp e left outer join dept d on e.deptno=d.deptno;
```

### 3.3 Right Outer Join

The records in the right table will be queried regardless of whether the conditions are met, and the left table can be queried only if the conditions are met. The records in the right table that do not meet the conditions are all NULLs in the left table

```mysql
mysql> select * from emp e right outer join dept d on e.deptno=d.deptno;
```







Exercise 1: Query all departments and corresponding employee information.

Exercise 2: Query the employee information and salary level of each employee. emp, salgrade

Exercise 3: Query each employee’s employee information, department name, department location, salary grade

Exercise 4: Query the employee information, department name, and salary grade in the department in New York.





Exercise 5: Query the number of people in each department, department name, department number.



Self-connection:

Query the name of the employee and the name of the superior:

 select e.empno,e.ename,e.mgr,m.empno,m.ename from emp e,emp m where e.mgr=m.empno;





## Four, subquery

1. Query information about employees whose salary is higher than allen.

select * from emp where sal> (select sal from emp where ename='allen');



Exercise 1: Query information about employees whose salary is not the highest or the lowest.

mysql> select * from emp where sal !=(select max(sal) from emp) and sal !=(select min(sal) from emp);



Exercise 2: Not the employee information of the sales department

dname--->deptno

Idea 1:

select deptno from dept where dname='sales'

mysql> select * from emp where deptno != (select deptno from dept where dname='sales');





Idea 2:

select deptno from dept where dname !='sales';

mysql> select * from emp where deptno in(select deptno from dept where dname !='sales');







Exercise 3: Query employee information, require that the salary is higher than any employee in the department number 10

Idea 1:

select mix(sal) from emp where deptno=10; // 

mysql> select * from emp where sal >(select min(sal) from emp where deptno=10);

Idea 2:

mysql> select * from emp where sal >any (select sal from emp where deptno=10);



Exercise 4: Query employee information, all persons in the department whose salary is greater than 30

Idea 1:

mysql> select * from emp where sal >(select max(sal) from emp where deptno=30);

Idea 2:

mysql> select * from emp where sal> all(select sal from emp where deptno=30);



Exercise 5: Query the detailed information of the company's highest paid employee

select max(sal) from emp;

select * from emp e,dept d

 where sal=(select max(sal) from emp) and e.deptno=d.deptno;

mysql> select e.*,d.* from emp e,dept d where sal=(select min(sal) from emp) and e.deptno=d.deptno;







## Five, summary

Primary and foreign key constraints

​ Primary key

​ Foreign key: two tables: parent table, child table

​ The foreign key in the child table must be the primary key in the parent table.

​ on delete set null/cascade/no action

Inquire

​ Simple query: deduplication, alias, addition, subtraction, multiplication and division operations, ifnull(a, b)

​ Conditional query: where

​ =,!=,<>,>,<,>=,<=

​ and, or, not

​ between and, in 

​ is null ,is not null

​ Sort: order by 

​ asc, desc

​ Fuzzy query: like 

​ _,%



​ Aggregate functions: max(),min(),count(),sum(),avg()

​ Paging: limit start, count

​ Group: group by

​ having 



Multi-meter joint check:

​ Internal connection: select. . . from table 1 alias inner join table 2 alias on join condition

​ Outer connection: left outer, right outer

​ select. . from left table alias left outer join right table alias on join condition

​ select. . from left table alias right outer join right table alias on join condition



Subquery:

​ Select nesting

​ select query results:





