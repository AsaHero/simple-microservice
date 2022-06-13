create extension if not exists "uuid-ossp";

create type Panel as enum ('IPS', 'OLED');

create table CPU (
    id serial primary key,
    brand varchar not null,
    name varchar not null,
    number_cores int not null,
    number_threads int not null,
    min_ghz float not null,
    max_ghz float not null
);

create table Screen (
    id serial primary key,
    width int not null,
    height int not null,
    size_inch float not null,
    panel Panel,
    multitouch boolean not null
);

create table Laptop (
    id serial primary key,
    laptop_id uuid default uuid_generate_v4(),
    brand varchar not null,
    name varchar not null,
    cpu_id int references CPU(id),
    ram json not null,
    gpu json not null,
    storage json not null,
    screen_id int references Screen(id),
    keyboard json not null,
    weight json not null,
    price_usd float not null,
    release_year int, 
    update_at timestamp,
    created_at timestamp default CURRENT_TIMESTAMP
);