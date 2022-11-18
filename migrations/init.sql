create extension if not exists "uuid-ossp";

create table balance (
    id uuid primary key,
    user_id uuid unique not null,
    value bigint
        default 0
        check (value >= 0)
        not null
);

create table service (
    id uuid primary key,
    tittle varchar(255) not null
);

create type history_type as enum (
    'replenishment',
    'reserve',
    'confirmation',
    'cancel_reserve'
    );

-- create table type_history (
--     id uuid primary key,
--     description varchar(255) not null
-- );

create table report (
    id uuid primary key,
    service_id uuid references service not null,
    value bigint
        check (value >= 0)
        not null,
    date timestamp not null
);

create table reserve (
    id uuid primary key,
    balance_id uuid references balance not null,
    value bigint check (value >= 0) not null
);

create table history (
    id uuid primary key,
    balance_id uuid references balance not null,
    type_history history_type not null,
    order_id uuid,
    service_id uuid references service,
    value bigint check (value >= 0) not null,
    source_replenishment varchar(16),
    Date timestamp not null
);

insert into service (id, tittle) VALUES ('0ba5b953-9df7-4170-80bf-50d3d8e1111d', 'cleaning');
insert into service (id, tittle) VALUES ('0ba5b953-9df7-4170-80bf-50d3d8e2222d', 'repair');
insert into service (id, tittle) VALUES ('0ba5b953-9df7-4170-80bf-50d3d8e3333d', 'massage');