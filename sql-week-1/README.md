# SQL – Week 1 Challenges
For these challenges, you'll find starter files marked with the name of the challenge itself. 

## 1. CREATE (Hiring)
### Task
For this challenge, you need to create a table and add data to it.

### To-do
* Create a table called `employees` with:
  * `employee_id` as an `INT` (make sure this can never been `null` and that it auto increments)
  * `last_name` as a string of characters that can be up to 255 characters long
  * `first_name` as a string of characters that can be up to 255 characters long
* Insert into the table one record of an employee. The data can be anything, just make sure that it is the correct data type.
* Create a `SELECT` statement that retrieves everything from the `employees` table

### Table Schemas
#### employees 
##### Employees at your new company
| NAME          | TYPE                              |
| ------------- | --------------------------------- |
| `employee_id` | `INT` (not null & auto increment) |
| `last_name`   | `STRING`                          |
| `first_name`  | `STRING`                          |

## 2. ORDER BY (Happy Birthday)
### Task
For this challenge, you are given a table of student birthdays with their name, the month number of their birthday and the day number of their birthday.

### To-do
* Create a `SELECT` statement that uses `ORDER BY` statement to organize the birthdays in sequential order
* The resulting query should have two columns: student and their birthday
  * For birthday, you need to concatenate the month and day with a dash(-) between them.
  * Ex. If you have 1 for the month and 4 for the day, your result should be "1-4"

### Table Schemas
#### students
##### A list of students and their birthdays
| NAME      | TYPE           |
| --------- | -------------- |
| `name`    | `VARCHAR(255)` |
| `month`   | `INT`          |
| `day`     | `INT`          |

### Query Schemas
#### Resulting query
##### Results of students ordered by their birthday
| NAME         | TYPE           |
| ------------ | -------------- |
| `name`       | `VARCHAR(255)` |
| `birthday`   | month-day      |

## 3. LIKE (Like Wildcards)
### Task
For this challenge, you are given a table of customers with an id, name, and address.

### To-do
* Create a `SELECT` statement that uses wildcards and the `LIKE` statement to get the customers who live in apartments (Apt.)
  * **HINT:** If you want to see the data on the table, just use `SELECT * FROM customers`
* Resulting query should only have the customer's name and address

### Table Schemas
#### customers
##### A list of customers with their name and address
| NAME      | TYPE           |
| --------- | -------------- |
| `id`      | `INT`          |
| `name`    | `VARCHAR(255)` |
| `address` | `VARCHAR(255)` |

### Query Schemas
#### Resulting query 
##### Results of customers who live in apartments
| NAME      | TYPE           |
| --------- | -------------- |
| `name`    | `VARCHAR(255)` |
| `address` | `VARCHAR(255)` |

## 4. DISTINCT (Unique Customers)
### Task
For this challenge, you are given a table with a week's worth of transactions. Each transaction consists of a customer id, price of the item, and a card number used in the transaction. 

### To-do
* Create a `SELECT` statement that counts the unique number of customers that came in that week
  *  Ex. If the same customer came in multiple times, you would only count them once
* Resulting query should only have one column named `unique_customers` and the number of unique customers based on the data

### Table Schemas
#### transactions
##### A list of transactions
| NAME           | TYPE           |
| -------------- | -------------- |
| `customer_id`  | `INT`          |
| `price`        | `VARCHAR(255)` |
| `card_number`  | `VARCHAR(255)` |

### Query Schemas
#### Resulting query
##### Results of the number of unique customers
| NAME               | TYPE  |
| ------------------ | ----- |
| `unique_customers` | `INT` |

## 5. GROUP/ORDER BY (Counting Cats)
### Task
For this challenge, you are given a table of cats with an id, name, and age of the cats.

### To-do
* Create a `GROUP BY` statement to group all the cats by their age and count the cats who have the same age
* Format the data by using the `ORDER BY` statement to order the result in ascending order based on their age group

### Table Schemas
#### cats
##### A list of cats with their name and age
| NAME   | TYPE           |
| ------ | -------------- |
| `id`   | `INT`          |
| `name` | `VARCHAR(255)` |
| `age`  | `INT`          |

### Query Schemas
#### Resulting query
##### Results of cats grouped by their age and a count of those who represent the age
| NAME        | TYPE  |
| ----------- | ----- |
| `id`        | `INT` |
| `cat_count` | `INT` |

## 6. MAX (Super Score)
### Task
For this challenge, you are given a table of ACT scores. Each row consists of the id of the student who took the test and the score they got on that test. Students are allowed to take the test multiple times and use the best score they got.

### To-do
* Create a `SELECT` statement that gets the highest score for each of the students
  * Ex. If the student took the test twice and got a 25 and 29, the resulting query should return 29 for that student only
* Resulting query should only have two columns named `student_id` and `highest_score` which will be the highest score obtained by the student

### Table Schemas
#### scores
##### List of ACT scores taken by students
| NAME         | TYPE  |
| ------------ | ----- |
| `student_id` | `INT` |
| `score`      | `INT` |

### Query Schemas
#### Resulting query
##### List of students' id and their highest score obtained
| NAME            | TYPE  |
| --------------- | ----- |
| `student_id`    | `INT` |
| `highest_score` | `INT` |

## 7. SUM (Top 10 Customers)
### Task
For this challenge, you are given a table of transactions for a week. Each transaction consists of a customer id, price of the item, and a card number used in the transaction.

### To-do
* Create a `SELECT` statement that adds up the total spent by every customer
  * Ex. If a customer came in two times during the week and spent 20 dollars and 50 dollars, their total would show up as 70 dollars
* Resulting query should only have two columns named `customer` and `total` which will be the total the customer bought from the store that week
* Limit the list to find the 10 customers with the highest total spent

### Table Schemas
#### transactions
##### A list of transactions
| NAME          | TYPE           |
| ------------- | -------------- |
| `customer_id` | `INT`          |
| `price`       | `INT`          |
| `card_number` | `VARCHAR(255)` |

### Query Schemas
#### Resulting query
##### List only the 10 customers with the highest totals
| NAME       | TYPE  |
| ---------- | ----- |
| `customer` | `INT` |
| `total`    | `INT` |