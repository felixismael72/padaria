create table product
(
    id              serial primary key,
    name            varchar(50) not null,
    code            varchar(50) not null,
    price           float      not null,
    expiration_date timestamp   not null
);