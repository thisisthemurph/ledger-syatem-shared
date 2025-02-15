create type transaction_status as enum ('pending', 'posted', 'failed');

create table if not exists ledger (
    id uuid primary key default uuid_generate_v4(),
    transaction_id uuid not null,
    account_id uuid not null references accounts(id),
    amount bigint not null,
    status transaction_status not null default 'pending',
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create trigger tr_ledger_update_updated_at_timestamp
before update on ledger
for each row
execute function fn_update_updated_at_timestamp();
