-- Due to simplistic DB initialization within the program, multi-line statements are not allowed.

-- Games
insert into games values (1, "FIFA14");
insert into games values (2, "NHL14");

-- FIFA14 Clubs
insert into clubs (game, name, league, stars) values (1, "Bor. Dortmund", "Bundesliga", 5);
insert into clubs (game, name, league, stars) values (1, "FC Bayern Münich", "Bundesliga", 5);
insert into clubs (game, name, league, stars) values (1, "Paris SG", "Ligue 1", 5);
insert into clubs (game, name, league, stars) values (1, "Real Madrid", "La Liga", 5);
insert into clubs (game, name, league, stars) values (1, "Arsenal", "Premier League", 5);
insert into clubs (game, name, league, stars) values (1, "Inter Milan", "Seria A", 5);
-- TODO the rest and proper ratings

-- NHL14 Clubs
-- TODO

