# Conditions and comparison operators

In many programming and scripting languages, conditions and comparison operators are fundamental concepts used to make decisions and control the flow of execution. Here’s an overview of common conditions and comparison operators found in various languages, including Bash, Python, and SQL.

### 1. **Bash (Shell Scripting)**

**Conditions:**
- Conditions in Bash are typically used with `if`, `while`, and `until` statements.
- The `test` command or `[` and `]` are used to evaluate conditions.

**Comparison Operators:**
- **String Comparison:**
  - `=`: Check if two strings are equal.
  - `!=`: Check if two strings are not equal.
  - `<`: Check if one string is less than another (in alphabetical order).
  - `>`: Check if one string is greater than another (in alphabetical order).
- **Numeric Comparison:**
  - `-eq`: Equal to.
  - `-ne`: Not equal to.
  - `-lt`: Less than.
  - `-le`: Less than or equal to.
  - `-gt`: Greater than.
  - `-ge`: Greater than or equal to.

**Example:**
```bash
#!/bin/bash

num1=10
num2=20
if [ $num1 -lt $num2 ]; then
    echo "$num1 is less than $num2"
fi

str1="hello"
str2="world"
if [ "$str1" = "$str2" ]; then
    echo "Strings are equal"
else
    echo "Strings are not equal"
fi
```

### 2. **Python**

**Conditions:**
- Conditions in Python are used with `if`, `elif`, and `else` statements.
- Python uses boolean expressions and relies on indentation for blocks of code.

**Comparison Operators:**
- `==`: Equal to.
- `!=`: Not equal to.
- `<`: Less than.
- `<=`: Less than or equal to.
- `>`: Greater than.
- `>=`: Greater than or equal to.

**Example:**
```python
num1 = 10
num2 = 20
if num1 < num2:
    print(f"{num1} is less than {num2}")

str1 = "hello"
str2 = "world"
if str1 == str2:
    print("Strings are equal")
else:
    print("Strings are not equal")
```

### 3. **SQL**

**Conditions:**
- SQL uses conditions in `WHERE`, `HAVING`, and `CASE` statements to filter records and perform conditional logic.

**Comparison Operators:**
- `=`: Equal to.
- `<>` or `!=`: Not equal to.
- `<`: Less than.
- `<=`: Less than or equal to.
- `>`: Greater than.
- `>=`: Greater than or equal to.
- `BETWEEN`: Range check.
- `LIKE`: Pattern matching.
- `IN`: Check if a value is in a list of values.
- `IS NULL`: Check for NULL values.

**Example:**
```sql
SELECT *
FROM employees
WHERE salary > 50000;

SELECT name
FROM employees
WHERE department IN ('HR', 'IT');

SELECT name,
       CASE
           WHEN salary >= 60000 THEN 'High'
           WHEN salary >= 40000 THEN 'Medium'
           ELSE 'Low'
       END AS salary_range
FROM employees;
```

### Summary

- **Bash**: Uses `test` command or `[ ]` with `-eq`, `-ne`, `-lt`, `-le`, `-gt`, `-ge` for numeric comparisons and `=`, `!=`, `<`, `>` for string comparisons.
- **Python**: Uses `==`, `!=`, `<`, `<=`, `>`, `>=` for both numeric and string comparisons.
- **SQL**: Uses `=`, `<>`, `!=`, `<`, `<=`, `>`, `>=`, `BETWEEN`, `LIKE`, `IN`, and `IS NULL` for comparisons and conditions.

Understanding these operators and how to use them in different contexts is crucial for writing effective conditional statements and performing comparisons in your scripts and queries.
