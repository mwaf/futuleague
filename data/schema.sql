-- Due to simplistic DB initialization within the program, multi-line statements are not allowed.
create table players (id integer not null primary key, identifier text not null unique, name text not null, rating real not null, played integer not null, wins integer not null, losses integer not null, ties integer not null);
create table clubs (id integer not null primary key, name text not null, league text not null, country text not null, stars real not null);
