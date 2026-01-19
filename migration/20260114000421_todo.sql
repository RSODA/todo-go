-- +goose Up
create table todo (
    id serial primary key,
    title text not null,
    created_at timestamp not null default now()
);

-- +goose Down
drop table todo
