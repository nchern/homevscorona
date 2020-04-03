create type UserStatus as enum ('HEALTHY', 'POSITIVE', 'RECOVERED');

create table user_list (
    id              uuid            PRIMARY KEY,
    created 		bigint          not null, 
    updated		    bigint          not null,
    status          UserStatus      not null,
    email           TEXT            not null,
    body            JSONB           not null
);

create index user_list_created_idx    on user_list (created);
create index user_list_updated_idx    on user_list (updated);
create index user_list_status_idx     on user_list (status);

create unique index user_list_email_idx  on user_list (email);


create table event (
    id              uuid            PRIMARY KEY,
    user_Id         uuid            not null,
    created 		bigint          not null, 
    body            JSONB           not null
);

create index event_created_idx    on event (created);
create index event_user_id_idx    on event (user_Id);
