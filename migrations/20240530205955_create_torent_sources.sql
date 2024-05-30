-- +goose Up
-- +goose StatementBegin
create table torent_sources (
    id serial primary key,
    name varchar(255) not null,
    url varchar(255) not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists torent_sources;
-- +goose StatementEnd
