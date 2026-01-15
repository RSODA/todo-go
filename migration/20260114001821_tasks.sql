-- +goose Up
-- +goose StatementBegin
create table task {
    id serial primary key,
    todo_id bigint not null,
    description text not null,
    is_complete bool not null default false
}
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table task
-- +goose StatementEnd
