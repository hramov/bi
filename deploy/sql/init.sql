create database bi
    with owner postgres;

\connect bi

create table dashboards (
id serial unique,
dash_id uuid unique not null default gen_random_uuid(),
title varchar(50),
description text,
created_at timestamp default now(),
updated_at timestamp,
deleted_at timestamp
);

create table dashboard_item_types (
          id serial unique,
          title varchar(50),
          name varchar(50)
);

insert into dashboard_item_types (title, name) values
('График', 'chart'),
('Таблица', 'table');

create table dashboard_items
(
    id           serial
        unique,
    dash_id      uuid not null,
    item_type    integer
        references dashboard_item_types (id),
    position     varchar(2),
    title        varchar(50),
    description  text,
    data_queries jsonb,
    raw_options  jsonb,
    created_at   timestamp default now(),
    updated_at   timestamp,
    deleted_at   timestamp
);

alter table dashboard_items
    owner to postgres;


create table users (
id serial unique,
user_id uuid unique not null,
fio varchar(50),
ad varchar(50),
is_active bool default false,
created_at timestamp default now(),
updated_at timestamp,
deleted_at timestamp
);

alter table users add constraint idx_users_user_id unique (user_id);

create table item_types (
id serial unique,
title varchar(50)
);

create table acl (
id serial unique,
item_type integer references item_types(id),
item_id uuid,
reading bool default false,
creating bool default false,
updating bool default false,
deleting bool default false,
created_at timestamp default now(),
updated_at timestamp,
deleted_at timestamp
);

create table user_acl (
user_id uuid references users(user_id),
acl_id integer references acl(id)
);

create table available_drivers (
       id serial unique,
       title varchar(50),
       code varchar(50),
       date_created timestamp default now()
);

create table drivers
(
    id serial unique,
    title        varchar(50),
    code         varchar(50),
    date_created timestamp default now(),
    db_need      boolean   default true
);

alter table drivers
    owner to postgres;

insert into drivers (title, code) values
('Postgres', 'pg'),
('SQL Server', 'sqlserver');

create table data_sources (
  id serial unique,
  driver_id integer references drivers(id),
  title varchar(50),
  dsn varchar(100),
  checked bool default false,
  date_created timestamp default now()
);

create table format_functions (
      id serial unique,
      title varchar(100) unique,
      name varchar(50) unique
);