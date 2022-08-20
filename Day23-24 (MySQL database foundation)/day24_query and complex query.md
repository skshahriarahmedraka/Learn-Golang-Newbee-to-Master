

# Data table query statement

The most commonly used operations for users such as data tables or views are queries. Also called retrieval. Through the select statement to achieve
Grammar rules:

```mysql
select {columns}
  from {table|view|other select}
  [where query conditions]
  [group by group condition]
  [having group and then limit]
  [order by sort]
```

Precautions:

1. SQL statements are not case sensitive. The value is also not divided. (The value in Oracle is case sensitive)
2. Keyword cannot be branched

## One, simple query

Among the select retrieval statements, there are only select and from clauses. It is also required for a select statement.

### 1.1 Specify column name

Query the specified column:

```mysql
 select column name, column name from table name
```

>If you query all the fields in the table, use *.

### 1.2 Specify alias

 ```mysql
select column name as alias, column name alias from table name
 ```

>as keyword: In fact, it can also be omitted.

### 1.3 Column Operation

Arithmetic expressions can be used

  In the select statement, if the retrieved field type is number, date can use arithmetic expressions
  Add: +, subtract: -, multiply: *, divide: /, parentheses

> If a column has a null value, then the calculated result is also null. General function: ifnull(a,b), if a is null then take b, if a is not null then take the value of a.

### 1.4 String type can do continuous operations

concat("My name is", name,'My year...');

  Splice the searched columns and columns for display.

### 1.5 De-duplication: distinct field name

  distinct column name, which means removing duplicates
  Distinct column name 1, column name 2, the data in a row are the same, will be considered as duplicate data, and removed.

## Second, condition query

In the actual query process, we need to filter out the data we want according to our needs. This kind of query is called a conditional query. When retrieving data information, use limited conditions. Indicates that the condition will be retrieved only when the conditions are met.

Use the where clause syntax format:

```mysql
select search column from table name where filter conditions
```

> The conditional query uses the `where` clause to filter the data in the table, and the rows with the result of `true` will appear in the result set.
>
> Conditional expressions support multiple operations:

### 2.1 Comparison operators

-Equal to `=`
-Greater than `>`
-Less than `<`
-Greater than or equal to `>=`
-Less than or equal to `<=`
-Not equal to `!=` or `<>`

---

Find people older than 30:
```mysql
select * from stus where age>30;
```



The query name is not `Wang Ergou`:

```mysql
select * from stus where name!='Wang Ergou';
or
select * from stus where name<>'Wang Ergou';
```
### 2.2 Logical operators

-and, && (and)
-or, || (or)
-not,! (non)

---

Query if the age is greater than 30 and the gender is female:
```mysql
select * from stus where age>30 and sex='女';
or
select * from stus where age>30 && sex='女';

```
Check if the age is less than 30 or the gender is male:

```mysql
select * from stus where age<30 or sex='male';
or
select * from stus where age<30 || sex='男';
```

---

### 2.3 Fuzzy query

For example, if you want to query the requirement that the name contains'Zhang', you need to use fuzzy query.

Fuzzy query uses the keyword `like`

```mysql
select * from stus name like'%张%';
Li Li Zhang Zhang
Wang Zhang
open
Zhang San
Zhang Sansan
```



grammar:

1. `%` means any number of characters (0-multiple)

2. `_` (underscore) means any character



The query name is two characters:

```mysql
select * from stus where name like'张__';
```

---

### 2.4 Interval query

#### 2.4.1 in (...)

-`in` means to search in a discrete and non-continuous range

Find the value of id equal to 1 or 3 or 4:

```mysql
select * from stus where id in (1, 3,4);
```


#### 2.4.2 between...and...

-`between...and...` 
-Means to search in a continuous interval.

Find the id value between 3 and 7:
```mysql
select * from stus where id between 3 and 7;
```

### 2.5 NULL judgment

`is NULL` is used to determine whether a column is empty. Note: It is also possible to use lowercase `NULL`.



## Three, sort

By default, the order we found is in accordance with the order in which the data is inserted in the database.
But in actual situations, we need to sort according to different conditions, such as sorting by update date, or sorting by price from low to high, or sorting by popularity from high to bottom, etc.
**mysql supports reordering the results of the query.** 
Use the order by column clause to complete the sorting function. The following columns indicate the sorting rules.

> order by clause. Syntactically located at the end of a SQL statement.

### 3.1 Ascending order (asc)

The default is ascending sort
Sort by age from lowest to highest.

```mysql
select * from stus order by age;
```



> Description:
>
> 1. By default, it is sorted in ascending order. asc
>
> You can also add asc (ascend ascend) after the column to indicate ascending order.
>
> 2. Use desc to indicate descending order
>
> desc is the abbreviation of the word descend (descend) to indicate descending order.

### 3.2 Descending (desc) 

```mysql
select * from stus order by age asc;
```



### 3.3 Multiple sorting rules 

You can sort according to certain rules first, when encountering the same situation, sort according to the second rule, ...
Sort in ascending order by age first, and then sort in descending order by id if the ages are equal.

```mysql
select * from stus order by age asc,id desc;
```



## Fourth, aggregate function

In the query, statistics, summation, etc. are very commonly used, and it can be very easy to complete through aggregation functions.

If there is no aggregate function, various sub-queries may be required, and various SQL nesting can be completed.

But with **aggregate functions**, many problems can be solved.

MySQL provides 5 commonly used aggregate functions: `sum(), max(), min(), avg(), count()`

Of course, grouping `group by and having clauses` are also used with aggregate functions, which will be discussed in detail in the next section.

### 4.1 sum (column)

This aggregate function is used to sum the values ​​of those columns that are not `null`.

-Sum all ages
  ​

-Summation of ages in multiple locations less than 100

  

### 4.2 max (column)

This aggregate function is used to calculate the maximum value of those columns that are not `null`.

-Find the oldest age

  

### 4.3 min (column)

This aggregate function is used to calculate the minimum value of those columns that are not `null`.

-Find the youngest age


### 4.4 avg (column)

This aggregate function is used to calculate the average value of those columns that are not `null`.

-Calculate the average value of `age`

  

### 4.5 count (column)

Statistical functions.

count(*)--->Statistics of all rows of data, 16

count (primary key column)-->16

count (non-primary key column 1), whether there is a null value.

count(comm)-->6



-Count the total number of rows of data

  

-Count the number of data rows whose gender is not `null`

  


## Five, group query

Group by
The group by column name is grouped according to the specified column, and those with the same value will be grouped together.
grammar:

```mysql
select column name from table name group by column name;
```


instruction:

1. The column name after select must be consistent with the column name after group by.
2. When group by is used alone, only the first record of each group is displayed. Therefore, the meaning of group by alone is not significant. Most of them must be combined with aggregate functions.
3. After group by, you can also group multiple columns, which means that these columns are in a group when they are the same. 

### 5.1 Group by column

```mysql
select sex from stus group by sex;
```

### 5.2 Group by multiple columns

```mysql
select name, sex from stus group by name,sex; 
```

### 5.3 Use aggregate functions after grouping

```mysql
select sex,count(*) from stus group by sex;
```



> Note:
>
> 1. If there is no grouping, then the group function (max, min, avg, count...) is applied to the entire table.
> 2. If there is grouping, the group function is applied to the grouped data.
> 3. In the write select clause, if it is not in the group function, then it must be after the group by.

```mysql
select a, b, count (c), sum (d) from table group by a, b;
```



### 5.4 Limit query after grouping: having

Secondary screening: Secondary screening refers to filtering the data after grouping.
Need to use the having clause to complete.

```mysql
select column name from table name group by column name having conditions
```

In other words, limit after grouping, use the having clause

The having clause filters the results after group by grouping again. The final output will satisfy the having condition.

> The usage of the having clause and the where clause is similar, both of which are used to qualify conditions.
>
> Contrast:
> 1. where and having are followed by conditions
> 2. where is the original filtering of the data in the table
> 3. Having is a secondary screening of the results of group by
> 4. Having must be used in conjunction with group by, and is generally used with aggregate functions
> 5. You can have where first, followed by group by and having
>    
> Differences and conclusions:
> 1. Syntactically: group functions (max, min, avg, count...) are used in having, and group functions cannot be used after where.
> 2. Implementation: where is to filter and then group. Having is to group and then filter.



For example: filter out sex that is not null.

For example: filter out the male or female more than 6




## Six, pagination query

Limit clause (dialect)

> The dialect means limit is unique to mysql.

Limit is used to limit the starting row of the query results and the total number of rows.

```mysql
select * from table name limit start,count;
```

For example: the first line of the query is line 5, a total of 3 lines

select * from stus limit 4, 3;

​ Where 4 means starting from the 5th row, of which 2 tables are query 3 rows. That is, 5, 6, 7 rows of records.

## Seven, built-in functions [extended]

### 7.1 String functions

-View the ascii code value of the character ascii(str), return 0 when str is an empty string

```mysql
select ascii('a');
```

-View the character char (number) corresponding to the ascii code value

```mysql
select char(97);

```

-Concatenation string concat(str1,str2...)

```mysql
select concat(12,34,'ab');

```

-Contains the number of characters length(str)

```mysql
select length('abc');

```

-Intercept string
  -left(str,len) returns len characters from the left end of the string str
  -right(str,len) returns len characters from the right end of the string str
  -substring(str,pos,len) returns len characters from the position pos of the string str, starting from 1

```mysql
select substring('abc123',2,3);

```

-Remove spaces
  -ltrim(str) returns the string str with the left space removed
  -rtrim(str) returns the string str with the right space removed
  -trim([direction remstr from str) returns the string str after removing remstr from a certain side. The direction words include both, leading, trailing, indicating both sides, left and right

```mysql
select trim(' bar');
select trim(leading'x' FROM'xxxbarxxx');
select trim(both'x' FROM'xxxbarxxx');
select trim(trailing'x' FROM'xxxbarxxx');
```

-Return a string space(n) consisting of n space characters

```mysql
select space(10);

```

-Replace string replace(str, from_str, to_str)

```mysql
select replace('abc123','123','def');
```

-Case conversion, the function is as follows
  -lower(str)
  -upper(str)

```mysql
select lower('aBcD');

```

### 7.2 Mathematical functions

-Find the absolute value abs(n)

```mysql
select abs(-32);
```

-Find the remainder mod(m,n) of m divided by n, same as operator%

```mysql
select mod(10,3);
select 10%3;
```

-Floor(n), which represents the largest integer not greater than n

```mysql
select floor(2.3);
```

-Ceiling(n), which represents the largest integer not less than n

```mysql
select ceiling(2.3);

```

-Find the rounding value round(n, d), n represents the original number, d represents the decimal position, and the default is 0

```mysql
select round(1.6);

```

-Find x to the power of y pow(x,y)

```mysql
select pow(2,3);

```

-Get PI()

```mysql
select PI();
```

-Random number rand(), a floating-point number with a value of 0-1.0

```mysql
select rand();
```

-There are many other trigonometric functions, you can query the document when you use it

### 7.3 Date and time functions

-Get the sub-value, the syntax is as follows
  -year(date) returns the year of date (range 1000 to 9999)
  -month(date) returns the month value in date
  -day(date) returns the date value in date
  -hour(time) returns the number of hours of time (range is 0 to 23)
  -minute(time) returns the number of minutes of time (range is 0 to 59)
  -second(time) returns the number of seconds of time (range is 0 to 59)

```mysql
select year('2016-12-21');
```

-Date calculation, use the +- operator, the keywords after the number are year, month, day, hour, minute, second

```mysql
select '2016-12-21'+interval 1 day;
select '2017-12-12' + interval 3 month;
select date_add('2017-12-12', interval 1 day);
```

-Date format date_format(date,format), the available values ​​of format parameter are as follows

  -Get the year %Y and return a 4-digit integer

    * Get year %y, return a 2-digit integer

    * Get month %m, the value is an integer of 1-12

  -Get day %d, return integer

    * When obtaining %H, the value is an integer from 0-23

    * When obtaining %h, the value is an integer of 1-12

    * Get points %i, the value is an integer from 0-59

    * Get the second %s, the value is an integer of 0-59

```mysql
select date_format('2016-12-21','%Y %m %d');
select str_to_date('12/12/2017','%d/%m/%Y');
```

-Current date current_date()

```mysql
select current_date();

```

-Current time current_time()

```mysql
select current_time();

```

-Current date and time now()

```mysql
select now();
```



## Seven, multi-table query

Join query: When you need to query multiple tables with relationships, you need to use join

**Merge result set**

1. In the table to be merged, the type and number of columns are the same
2. UNION, remove duplicate rows
3. UNION ALL, do not remove duplicate rows



**Classification** 

	1. Internal connection
	2. External connection
		 A: Left outer connection
		 B: Right outer connection
		 C: Full external connection (Mysql does not support)
	3. Self-connection (term a simplified way)

### 7.1 Cartesian product

If the two tables are in the connection query, if there is no connection condition, then a Cartesian product will be generated. (Redundant data).

### 7.1 Internal connection

All the records in the inner join query meet the conditions

```mysql
Dialect: select * from table 1 alias 1, table 2 alias 2 where alias 1.xx=alias 2.xx
	Equivalent link
Standard: select * from table 1 alias 1 inner join table 2 alias 2 on alias 1.xx=alias 2.xx
Natural: select * from table 1 alias 1 natural join table 2 alias 2

```



### 7.2 External connection

The results retrieved by the inner join are all satisfying the join conditions. The outer link is the result set retrieved by the extended inner link. The results returned by the external link include records that meet the link conditions as well as those that do not meet the link conditions.

#### 7.2.1 Left outer join

The records in the left table will be queried regardless of whether the conditions are met, and the right table can be queried only if the conditions are met. The records in the left table that do not meet the conditions, the right table part is NULL

Syntax format:

```mysql
Left outer:
	select * from table 1 alias 1 left outer join table 2 alias 2 on alias 1.xx=alias 2.xx
	
Left outside nature:
	select * from table 1 alias 1 natural left outer join table 2 alias 2 
```



#### 7.2.2 Right Outer Join

The records in the right table will be queried regardless of whether the conditions are met, and the left table can be queried only if the conditions are met. The records in the right table that do not meet the conditions are all NULLs in the left table

Syntax format:

```mysql

Right outer:
	select * from table 1 alias 1 right outer join table 2 alias 2 on alias 1.xx=alias 2.xx
	

Right outside nature:
	select * from table 1 alias 1 natural right outer join table 2 alias 2 

```

#### 7.2.3 Full external connection

Mysql does not support full outer join, but union can be used to merge the results of left outer join and right outer join to achieve the effect of full outer join.

```mysql
Left outer join
SELECT e.ename,e.sal, d.dname,e.deptno,d.deptno
FROM emp e LEFT OUTER JOIN dept d
ON e.deptno=d.deptno
UNION
Right outer join
SELECT e.ename,e.sal, d.dname,e.deptno,d.deptno
FROM emp e RIGHT OUTER JOIN dept d
ON e.deptno=d.deptno
```



### 7.3 Self-connection

Self-connection (a simplified way of terminology). In fact, it is a table of connection bytes.

Query the employee’s superior name

```mysql
SELECT e.empno,e.ename,e.mgr,m.empno,m.ename
FROM emp e LEFT OUTER JOIN emp m
ON e.mgr=m.empno
```

## Eight, subquery

Subquery refers to the SQL statement contains another select statement. Inner query or inner select statement

> A SQL statement contains multiple select keywords	
>
> External query, internal query

Where subqueries can appear:

​ After from, as a table

​ After where, as a condition

Precautions:
1. The subquery must be in ()
2. The order by clause cannot be used in the subquery.
3. Sub-queries can be nested sub-queries, up to 255 levels.



Sub-queries are subdivided into: single-line sub-queries, multi-line sub-queries, and related sub-queries.

1. Single row subquery
  The result of the subquery is a single row of data.
    After the where condition, you need to cooperate with the single-line operator: >,<,>=,<=,!=,=
2. Multi-line subquery
  The result of the subquery is multiple rows of data.
    Need to cooperate after the where condition: in, any, all operators
    in:
    any: Match any one in the result set of the subquery.
    all: Match all the results of the subquery.

3. Associated subqueries
  For single-row subquery and multi-row subquery, outer query and inner query are executed separately.
  If the outer query uses the result of the inner query, use the associated subquery.
  Associated subquery refers to the need to resort to outer query in inner query, and outer query cannot be separated from the execution of inner query. It's called a correlated subquery.

## Nine, summary

The complete select statement is written:

```mysql
select distinct *
from table name
where ....
group by ... having ...
order by ...
limit start,count
```

Order of execution:

```mysql
from table name
where ...
group by ...
having ...
select distinct*..
order by ...
limit start,count
```

In actual use, only a combination of some parts of the sentence is used, but not all of them.

