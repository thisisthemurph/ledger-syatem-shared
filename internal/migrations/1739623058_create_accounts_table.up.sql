create table if not exists accounts (
    id uuid primary key default uuid_generate_v4(),
    name text not null,
    balance bigint not null default 0,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create trigger tr_accounts_update_updated_at_timestamp
before update on accounts
for each row
execute function fn_update_updated_at_timestamp();
