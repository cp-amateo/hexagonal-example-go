create schema if not exists example;

create table if not exists customers
(
    id       integer  not null
    primary key,
    name     varchar(50)            null,
    surnames varchar(50)            null,
    dob      date                   null,
    constraint customers_id_uindex
    unique (id)
    );

