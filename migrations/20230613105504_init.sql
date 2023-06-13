-- +goose Up
-- +goose StatementBegin
create table urls
(
    id bigint primary key,
    short varchar(20) unique not null,
    orig varchar(2048) unique not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists urls;
-- +goose StatementEnd
