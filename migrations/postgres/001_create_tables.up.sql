create table users (
    id uuid primary key not null,
    full_name varchar(30),
    phone varchar(30) unique not null,
    password varchar(30) not null
);