# pets
A go microservice for Swagger Petstore

# Setup
Run the start.sh script.

This will bring up 2 docker containers, one for the postgres database and one for the api.

Enter the postgres container with:
```
docker exec -ti pets_db_1 psql -U postgres
```

Create and enter the pets database:
```
CREATE DATABASE pets;
\c pets
```

Create the schema by copying the contents of db/schema.sql and pasting inside postgres. The output should look like this:
```
postgres=# CREATE DATABASE pets;
CREATE DATABASE
postgres=# \c pets
You are now connected to database "pets" as user "postgres".
pets=# CREATE TABLE IF NOT EXISTS pet_category (
pets(#   pet_category_id   SERIAL PRIMARY KEY NOT NULL,
pets(#   pet_category_name TEXT    NOT NULL,
pets(#   UNIQUE(pet_category_name)
pets(# );
CREATE TABLE
pets=# 
pets=# CREATE TABLE IF NOT EXISTS pet_status (
pets(#   pet_status_id   INTEGER PRIMARY KEY NOT NULL,
pets(#   pet_status_name TEXT    NOT NULL
pets(# );
CREATE TABLE
pets=# 
pets=# CREATE TABLE IF NOT EXISTS pet (
pets(#   pet_id          SERIAL  PRIMARY KEY NOT NULL,
pets(#   pet_name        TEXT    NOT NULL,
pets(#   pet_category_id INTEGER NOT NULL REFERENCES pet_category(pet_category_id),
pets(#   pet_photo_urls  JSONB   NOT NULL,
pets(#   pet_tags        JSONB   NOT NULL,
pets(#   pet_status_id   INTEGER NOT NULL REFERENCES pet_status(pet_status_id)
pets(# );
CREATE TABLE
pets=# 
pets=# CREATE TABLE IF NOT EXISTS "user" (
pets(#   user_id    SERIAL  PRIMARY KEY NOT NULL,
pets(#   username   TEXT    NOT NULL,
pets(#   first_name TEXT    NOT NULL,
pets(#   last_name  TEXT    NOT NULL,
pets(#   email      TEXT    NOT NULL,
pets(#   password   TEXT    NOT NULL,
pets(#   phone      TEXT    NOT NULL,
pets(#   status     INTEGER NOT NULL DEFAULT 0
pets(# );
CREATE TABLE
pets=# 
pets=# CREATE TABLE IF NOT EXISTS order_status (
pets(#   order_status_id   INTEGER PRIMARY KEY NOT NULL,
pets(#   order_status_name TEXT    NOT NULL
pets(# );
CREATE TABLE
pets=# 
pets=# CREATE TABLE IF NOT EXISTS "order" (
pets(#   order_id        SERIAL    PRIMARY KEY NOT NULL,
pets(#   pet_id          INTEGER   NOT NULL REFERENCES pet(pet_id),
pets(#   quantity        INTEGER   NOT NULL,
pets(#   ship_date       TIMESTAMP WITH TIME ZONE NOT NULL,
pets(#   order_status_id INTEGER   NOT NULL REFERENCES order_status(order_status_id)
pets(# );
CREATE TABLE
pets=# 
```

Create initial values by copying the contents of db/initial_values.sql. The output should look like this:
```
pets=# INSERT INTO pet_status (pet_status_id, pet_status_name)
pets-# VALUES
pets-#   (10, 'available'),
pets-#   (20, 'pending'),
pets-#   (30, 'sold')
pets-# ON CONFLICT (pet_status_id)
pets-# DO NOTHING;
INSERT 0 3
pets=# 
pets=# INSERT INTO order_status (order_status_id, order_status_name)
pets-# VALUES
pets-# (10, 'placed'),
pets-# (20, 'approved'),
pets-# (30, 'delivered')
pets-# ON CONFLICT (order_status_id)
pets-# DO NOTHING;
INSERT 0 3
pets=# 
```

# Notes
This was created in 4 hours, most endpoints have not been implemented.
