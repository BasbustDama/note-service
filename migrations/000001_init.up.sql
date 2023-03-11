CREATE TABLE "customer" (
    "id" SERIAL NOT NULL UNIQUE,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(360) NOT NULL UNIQUE,
    "password_hash" VARCHAR(255) NOT NULL,

    PRIMARY KEY ("id")
);

CREATE TABLE "note" (
    "id" SERIAL NOT NULL UNIQUE,
    "customer_id" INTEGER NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "description" VARCHAR,

    PRIMARY KEY ("id"),
    FOREIGN KEY ("customer_id") REFERENCES "customer" ("id")
);