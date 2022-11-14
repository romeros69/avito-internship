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
    'cancellation of a reservation'
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
    source_replenishment varchar(16),
    Date timestamp not null
);


-- функция проверки номера карты
-- CREATE OR REPLACE FUNCTION isnumeric(varchar(16))
-- RETURNS BOOLEAN AS $$
-- DECLARE
--     x NUMERIC;
-- BEGIN
--     x = $1::NUMERIC;
--     RETURN TRUE;
-- EXCEPTION WHEN others THEN
--     RETURN FALSE;
-- END;
-- $$ LANGUAGE 'plpgsql';