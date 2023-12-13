CREATE TABLE IF NOT EXISTS links
(
    id      serial        not null unique,
    base_url varchar(1024) not null,
    short_url   varchar(10)   not null unique
);