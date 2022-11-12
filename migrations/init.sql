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

create table type_history (
    id uuid primary key,
    description varchar(255) not null
);

create table report (
    id uuid primary key,
    service_id uuid references service not null,
    order_id uuid not null,
    value bigint
        check (value >= 0)
        not null,
    date timestamp not null
);

create table reserve (
    id uuid primary key,
    balance_id uuid references balance not null,
    value bigint check (value >= 0) not null,
    status varchar(9)
        check (status = 'active' or status = 'no active')
        not null
);

create table history_balance (
    id uuid primary key,
    balance_id uuid references balance not null,
    type_history_id uuid references type_history not null,
    reserve_id uuid references reserve,
    report_id uuid references report,
    source_replenishment varchar(16)
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