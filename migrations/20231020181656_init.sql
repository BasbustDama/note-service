-- +goose Up
-- +goose StatementBegin
CREATE TABLE "customer" (
    "id"        INT GENERATED ALWAYS AS IDENTITY,
    "username"  VARCHAR(255) NOT NULL,
    "password"  VARCHAR(512) NOT NULL,

    PRIMARY KEY ("id")
);

CREATE TABLE "note" (
    "id"            INT GENERATED ALWAYS AS IDENTITY,
    "customer_id"   INT NOT NULL,
    "title"         VARCHAR(255) NOT NULL,
    "description"   VARCHAR(4096),

    PRIMARY KEY ("id"),
    FOREIGN KEY ("customer_id") REFERENCES customer("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "note";

DROP TABLE "customer";
-- +goose StatementEnd
