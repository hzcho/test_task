create table if not exists buildings(
    id bigserial primary key,
    name varchar(255) not null,
    city varchar(255) not null,
    year int not null,
    floors int not null
);
