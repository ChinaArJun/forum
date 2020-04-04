create table users (
    id serial primary key,
    uuid varchar(64) not null unique,
    name varchar(255),
    email varchar(255) not null,
    password varchar(255) not null,
    created_at timestamp not null
);

create table sessions (
    id serial primary key,
    uuid varchar(64) not null unique,
    email varchar(255),
    user_id integer references users(id),  -- 关联用户表
    created_at timestamp not null
);

create table threads (
    id serial primary key,
    uuid varchar(64) not null unique,
    user_id integer references users(id),
    topic text,     -- 长文本
    created_at timestamp not null
);

create table posts (
    id serial primary key,
    uuid varchar(64) not null unique,
    body text,
    user_id integer references users(id),
    threads_id integer references threads(id),
    created_at timestamp not null
);