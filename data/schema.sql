-- Due to simplistic DB initialization within the program, multi-line statements are not allowed.
create table players (id integer not null primary key, identifier text not null unique, name text not null, rating real not null, played integer not null, wins integer not null, losses integer not null, ties integer not null);
create table clubs (id integer not null primary key, name text not null, league text not null, country text not null, stars real not null);
create table teams (id integer not null primary key, type text not null);
create table teams_1 (id integer not null primary key, team_id integer not null, player_id integer not null);
create table teams_2 (id integer not null primary key, team_id integer not null, player_id integer not null);
create table teams_3 (id integer not null primary key, team_id integer not null, player_id integer not null);
create table matches (id integer not null primary key, home_team_id integer not null, away_team_id integer not null, home_club_id integer not null, away_club_id integer not null, home_score integer not null, away_score integer not null, timestamp timestamp not null);