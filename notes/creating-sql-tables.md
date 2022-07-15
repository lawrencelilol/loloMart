 # Creating SQL tables

 ```sql
 CREATE TABLE table_name (
    field_name TYPE CONSTRAINTS,
    field_name TYPE(args) CONSTRAINTS
 );
 ```

 ```sql
CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL
);
 ```