CREATE SCHEMA IF NOT EXISTS "public";

CREATE  TABLE "public".cinema ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	address              text  NOT NULL  ,
	city                 varchar(100)  NOT NULL  ,
	CONSTRAINT pk_cinema PRIMARY KEY ( id ),
	CONSTRAINT unq_cinema_name UNIQUE ( name ) 
 );

CREATE  TABLE "public".genre ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	CONSTRAINT pk_genre PRIMARY KEY ( id ),
	CONSTRAINT unq_genre UNIQUE ( name ) 
 );

CREATE  TABLE "public".screen_type ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	CONSTRAINT pk_screen_type PRIMARY KEY ( id ),
	CONSTRAINT unq_screen_type_name UNIQUE ( name ) 
 );

CREATE  TABLE "public".film ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	title                varchar(100)  NOT NULL  ,
	director             varchar(100)  NOT NULL  ,
	genre_id             bigint  NOT NULL  ,
	duration             varchar  NOT NULL  ,
	CONSTRAINT pk_films PRIMARY KEY ( id ),
	CONSTRAINT unq_film UNIQUE ( title ) ,
	CONSTRAINT fk_film_genre FOREIGN KEY ( genre_id ) REFERENCES "public".genre( id ) ON DELETE CASCADE  
 );

CREATE  TABLE "public".screen ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	cinema_id            bigint  NOT NULL  ,
	name                 varchar(100)  NOT NULL  ,
	screen_type_id       bigint  NOT NULL  ,
	CONSTRAINT pk_screen PRIMARY KEY ( id ),
	CONSTRAINT fk_screen_cinema FOREIGN KEY ( cinema_id ) REFERENCES "public".cinema( id ) ON DELETE CASCADE  ,
	CONSTRAINT fk_screen_screen FOREIGN KEY ( screen_type_id ) REFERENCES "public".screen_type( id ) ON DELETE CASCADE  
 );

CREATE  TABLE "public".seat ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	screen_id            bigint  NOT NULL  ,
	seat_row             char(1)  NOT NULL  ,
	seat_number          integer  NOT NULL  ,
	seat_price_modifier  double precision  NOT NULL  ,
	CONSTRAINT pk_seats PRIMARY KEY ( id ),
	CONSTRAINT unq_seat UNIQUE ( screen_id, seat_row, seat_number ) ,
	CONSTRAINT fk_seat_screen FOREIGN KEY ( screen_id ) REFERENCES "public".screen( id ) ON DELETE CASCADE  
 );

CREATE  TABLE "public".showtime ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	film_id              bigint  NOT NULL  ,
	screen_id            bigint  NOT NULL  ,
	start_time           varchar(100)  NOT NULL  ,
	base_price           integer  NOT NULL  ,
	CONSTRAINT pk_showtime PRIMARY KEY ( id ),
	CONSTRAINT fk_showtime_film FOREIGN KEY ( film_id ) REFERENCES "public".film( id ) ON DELETE CASCADE  ,
	CONSTRAINT fk_showtime_screen FOREIGN KEY ( screen_id ) REFERENCES "public".screen( id ) ON DELETE CASCADE  
 );

INSERT INTO "public".cinema( name, address, city ) VALUES ( 'CineMaxx Pasuruan', 'Jl. Raya Panglima Sudirman No.10', 'Pasuruan');
INSERT INTO "public".cinema( name, address, city ) VALUES ( 'XXI Plaza Surabaya', 'Jl. Pemuda No.1', 'Surabaya');
INSERT INTO "public".cinema( name, address, city ) VALUES ( 'Trans Studio Premiere', 'Jl. Ahmad Yani No.45', 'Malang');
INSERT INTO "public".cinema( name, address, city ) VALUES ( 'Cinepolis Gresik Mall', 'Jl. Veteran No.7', 'Gresik');
INSERT INTO "public".cinema( name, address, city ) VALUES ( 'XXI Royal Plaza', 'Jl. A. Yani No.16', 'Surabaya');
INSERT INTO "public".cinema( name, address, city ) VALUES ( 'CGV Marvel City', 'Jl. Ngagel No.23', 'Surabaya');
INSERT INTO "public".cinema( name, address, city ) VALUES ( 'Movimax Jember', 'Jl. Gajah Mada No.9', 'Jember');
INSERT INTO "public".cinema( name, address, city ) VALUES ( 'XXI Lippo Plaza Batu', 'Jl. Diponegoro No.22', 'Batu');
INSERT INTO "public".cinema( name, address, city ) VALUES ( 'Cinepolis Tunjungan Plaza', 'Jl. Basuki Rahmat No.8', 'Surabaya');
INSERT INTO "public".cinema( name, address, city ) VALUES ( 'XXI City of Tomorrow', 'Jl. Ahmad Yani No.288', 'Surabaya');
INSERT INTO "public".genre( name ) VALUES ( 'Action');
INSERT INTO "public".genre( name ) VALUES ( 'Comedy');
INSERT INTO "public".genre( name ) VALUES ( 'Drama');
INSERT INTO "public".genre( name ) VALUES ( 'Horror');
INSERT INTO "public".genre( name ) VALUES ( 'Romance');
INSERT INTO "public".genre( name ) VALUES ( 'Thriller');
INSERT INTO "public".genre( name ) VALUES ( 'Animation');
INSERT INTO "public".genre( name ) VALUES ( 'Fantasy');
INSERT INTO "public".genre( name ) VALUES ( 'Sci-Fi');
INSERT INTO "public".genre( name ) VALUES ( 'Documentary');
INSERT INTO "public".screen_type( name ) VALUES ( 'Standard');
INSERT INTO "public".screen_type( name ) VALUES ( 'IMAX');
INSERT INTO "public".screen_type( name ) VALUES ( '4DX');
INSERT INTO "public".screen_type( name ) VALUES ( 'Dolby Atmos');
INSERT INTO "public".screen_type( name ) VALUES ( 'VIP');
INSERT INTO "public".screen_type( name ) VALUES ( 'Deluxe');
INSERT INTO "public".screen_type( name ) VALUES ( 'Gold Class');
INSERT INTO "public".screen_type( name ) VALUES ( 'Private Box');
INSERT INTO "public".screen_type( name ) VALUES ( 'Premium');
INSERT INTO "public".screen_type( name ) VALUES ( 'Classic');
INSERT INTO "public".film( title, director, genre_id, duration ) VALUES ( 'The Silent Code', 'Christopher Nolan', 6, '02:10:00');
INSERT INTO "public".film( title, director, genre_id, duration ) VALUES ( 'Love in Kyoto', 'Makoto Shinkai', 5, '01:55:00');
INSERT INTO "public".film( title, director, genre_id, duration ) VALUES ( 'Haunted Forest', 'James Wan', 4, '01:45:00');
INSERT INTO "public".film( title, director, genre_id, duration ) VALUES ( 'Beyond the Stars', 'Denis Villeneuve', 9, '02:20:00');
INSERT INTO "public".film( title, director, genre_id, duration ) VALUES ( 'The Last Samurai', 'Edward Zwick', 1, '02:30:00');
INSERT INTO "public".film( title, director, genre_id, duration ) VALUES ( 'Journey to Evermore', 'Peter Jackson', 8, '02:25:00');
INSERT INTO "public".film( title, director, genre_id, duration ) VALUES ( 'Laugh It Up', 'Taika Waititi', 2, '01:40:00');
INSERT INTO "public".film( title, director, genre_id, duration ) VALUES ( 'City Tears', 'Greta Gerwig', 3, '02:00:00');
INSERT INTO "public".film( title, director, genre_id, duration ) VALUES ( 'The Truth Unfolds', 'David Fincher', 6, '02:15:00');
INSERT INTO "public".film( title, director, genre_id, duration ) VALUES ( 'Animal Kingdom', 'Pixar Studio', 7, '01:50:00');
INSERT INTO "public".screen( cinema_id, name, screen_type_id ) VALUES ( 1, 'Studio 1', 1);
INSERT INTO "public".screen( cinema_id, name, screen_type_id ) VALUES ( 1, 'Studio 2', 2);
INSERT INTO "public".screen( cinema_id, name, screen_type_id ) VALUES ( 2, 'Studio 1', 1);
INSERT INTO "public".screen( cinema_id, name, screen_type_id ) VALUES ( 3, 'Theatre A', 4);
INSERT INTO "public".screen( cinema_id, name, screen_type_id ) VALUES ( 4, 'Hall 1', 3);
INSERT INTO "public".screen( cinema_id, name, screen_type_id ) VALUES ( 5, 'Screen 5', 5);
INSERT INTO "public".screen( cinema_id, name, screen_type_id ) VALUES ( 6, 'IMAX Hall', 2);
INSERT INTO "public".screen( cinema_id, name, screen_type_id ) VALUES ( 7, 'Screen A', 1);
INSERT INTO "public".screen( cinema_id, name, screen_type_id ) VALUES ( 8, 'Studio 3', 6);
INSERT INTO "public".screen( cinema_id, name, screen_type_id ) VALUES ( 9, 'VIP Room', 7);
INSERT INTO "public".seat( screen_id, seat_row, seat_number, seat_price_modifier ) VALUES ( 1, 'A', 1, 1.25);
INSERT INTO "public".seat( screen_id, seat_row, seat_number, seat_price_modifier ) VALUES ( 1, 'A', 2, 0.95);
INSERT INTO "public".seat( screen_id, seat_row, seat_number, seat_price_modifier ) VALUES ( 1, 'B', 1, 1.25);
INSERT INTO "public".seat( screen_id, seat_row, seat_number, seat_price_modifier ) VALUES ( 1, 'B', 2, 1.1);
INSERT INTO "public".seat( screen_id, seat_row, seat_number, seat_price_modifier ) VALUES ( 2, 'A', 1, 1.2);
INSERT INTO "public".seat( screen_id, seat_row, seat_number, seat_price_modifier ) VALUES ( 2, 'B', 2, 0.85);
INSERT INTO "public".seat( screen_id, seat_row, seat_number, seat_price_modifier ) VALUES ( 3, 'C', 3, 1.1);
INSERT INTO "public".seat( screen_id, seat_row, seat_number, seat_price_modifier ) VALUES ( 4, 'A', 1, 1.0);
INSERT INTO "public".seat( screen_id, seat_row, seat_number, seat_price_modifier ) VALUES ( 5, 'A', 1, 1.2);
INSERT INTO "public".seat( screen_id, seat_row, seat_number, seat_price_modifier ) VALUES ( 6, 'B', 2, 0.9);
INSERT INTO "public".showtime( film_id, screen_id, start_time, base_price ) VALUES ( 1, 1, '2025-11-05 10:00', 35000);
INSERT INTO "public".showtime( film_id, screen_id, start_time, base_price ) VALUES ( 2, 1, '2025-11-05 12:30', 40000);
INSERT INTO "public".showtime( film_id, screen_id, start_time, base_price ) VALUES ( 3, 2, '2025-11-05 14:00', 50000);
INSERT INTO "public".showtime( film_id, screen_id, start_time, base_price ) VALUES ( 4, 3, '2025-11-05 16:30', 40000);
INSERT INTO "public".showtime( film_id, screen_id, start_time, base_price ) VALUES ( 5, 4, '2025-11-05 18:00', 35000);
INSERT INTO "public".showtime( film_id, screen_id, start_time, base_price ) VALUES ( 6, 5, '2025-11-05 20:00', 40000);
INSERT INTO "public".showtime( film_id, screen_id, start_time, base_price ) VALUES ( 7, 6, '2025-11-05 10:30', 35000);
INSERT INTO "public".showtime( film_id, screen_id, start_time, base_price ) VALUES ( 8, 7, '2025-11-05 13:00', 40000);
INSERT INTO "public".showtime( film_id, screen_id, start_time, base_price ) VALUES ( 9, 8, '2025-11-05 15:30', 40000);
INSERT INTO "public".showtime( film_id, screen_id, start_time, base_price ) VALUES ( 10, 9, '2025-11-05 19:00', 50000);