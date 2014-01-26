-- Due to simplistic DB initialization within the program, multi-line statements are not allowed.
create table players (id integer not null primary key, identifier text not null, name text not null, rating real not null);
create table games (id integer not null primary key, name text not null);
create table clubs (id integer not null primary key, name text not null, league text not null, stars real not null);
