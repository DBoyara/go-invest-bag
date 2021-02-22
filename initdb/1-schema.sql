CREATE TABLE if not exists position
(
    tiker       TEXT NOT NULL constraint position_pkey PRIMARY KEY,
    created_at  BIGINT NOT NULL,
    updated_at  BIGINT NOT NULL,
    deleted_at  BIGINT NOT NULL,
    name        TEXT NOT NULL,
    type        TEXT NOT NULL,
    "count"     INT NOT NULL,
    price       REAL NOT NULL,
    amount      REAL NOT NULL,
    currency    TEXT NOT NULL
);