-- +goose Up
create table task (
                      id          serial primary key,
                      todo_id     bigint not null,
                      description text   not null,
                      is_complete bool   not null default false
);

-- +goose Down
drop table task
