-- Pets Schema

CREATE TABLE IF NOT EXISTS pet_category (
  pet_category_id   INTEGER PRIMARY KEY NOT NULL,
  pet_category_name TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS pet_status (
  pet_status_id   INTEGER PRIMARY KEY NOT NULL,
  pet_status_name TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS pet (
  pet_id          SERIAL  PRIMARY KEY NOT NULL,
  pet_name        TEXT    NOT NULL,
  pet_category_id INTEGER NOT NULL REFERENCES pet_category(pet_category_id),
  pet_tags        JSONB   NOT NULL
);

CREATE TABLE IF NOT EXISTS "user" (
  user_id    SERIAL  PRIMARY KEY NOT NULL,
  username   TEXT    NOT NULL,
  first_name TEXT    NOT NULL,
  last_name  TEXT    NOT NULL,
  email      TEXT    NOT NULL,
  password   TEXT    NOT NULL,
  phone      TEXT    NOT NULL,
  status     INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS order_status (
  order_status_id   INTEGER PRIMARY KEY NOT NULL,
  order_status_name TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS "order" (
  order_id        SERIAL    PRIMARY KEY NOT NULL,
  pet_id          INTEGER   NOT NULL REFERENCES pet(pet_id),
  quantity        INTEGER   NOT NULL,
  ship_date       TIMESTAMP WITH TIME ZONE NOT NULL,
  order_status_id INTEGER   NOT NULL REFERENCES order_status(order_status_id)
);
