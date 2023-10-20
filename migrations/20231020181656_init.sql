-- +goose Up
-- +goose StatementBegin
CREATE TABLE "note" (
    "id"            INT GENERATED ALWAYS AS IDENTITY,
    "title"         VARCHAR(255) NOT NULL,
    "description"   VARCHAR(4096),

    PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "note";
-- +goose StatementEnd
