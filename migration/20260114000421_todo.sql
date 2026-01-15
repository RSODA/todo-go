-- +goose Up
-- +goose StatementBegin
create table todo {
    id serial primary key,
    title text not null,
    created_at timestamp not null default now()
}

-- +goose Down
-- +goose StatementBegin
drop table todo
-- +goose StatementEnd
