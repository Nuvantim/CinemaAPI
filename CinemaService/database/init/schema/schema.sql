CREATE  TABLE "public".cinema ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	address              text  NOT NULL  ,
	city                 varchar(100)  NOT NULL  ,
	CONSTRAINT pk_cinema PRIMARY KEY ( id ),
	CONSTRAINT unq_cinema_name UNIQUE ( name ) 
 );

CREATE  TABLE "public".genre ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	CONSTRAINT pk_genre PRIMARY KEY ( id ),
	CONSTRAINT unq_genre UNIQUE ( name ) 
 );

CREATE  TABLE "public".screen_type ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	CONSTRAINT pk_screen_type PRIMARY KEY ( id ),
	CONSTRAINT unq_screen_type_name UNIQUE ( name ) 
 );


CREATE  TABLE "public".film ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	title                varchar(100)  NOT NULL  ,
	director             varchar(100)  NOT NULL  ,
	genre_id             integer  NOT NULL  ,
	duration             varchar  NOT NULL  ,
	CONSTRAINT pk_films PRIMARY KEY ( id ),
	CONSTRAINT unq_film UNIQUE ( title ) ,
	CONSTRAINT fk_film_genre FOREIGN KEY ( genre_id ) REFERENCES "public".genre( id ) ON DELETE CASCADE  
 );

CREATE  TABLE "public".screen ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	cinema_id            integer  NOT NULL  ,
	name                 varchar(100)  NOT NULL  ,
	screen_type_id       integer  NOT NULL  ,
	CONSTRAINT pk_screen PRIMARY KEY ( id ),
	CONSTRAINT fk_screen_cinema FOREIGN KEY ( cinema_id ) REFERENCES "public".cinema( id ) ON DELETE CASCADE  ,
	CONSTRAINT fk_screen_screen FOREIGN KEY ( screen_type_id ) REFERENCES "public".screen_type( id ) ON DELETE CASCADE  
 );

CREATE  TABLE "public".seat ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	screen_id            integer    ,
	seat_row             char(1)  NOT NULL  ,
	seat_number          integer  NOT NULL  ,
	seat_price_modifier  double precision  NOT NULL  ,
	CONSTRAINT pk_seats PRIMARY KEY ( id ),
	CONSTRAINT unq_seat UNIQUE ( screen_id, seat_row, seat_number ) ,
	CONSTRAINT fk_seat_screen FOREIGN KEY ( screen_id ) REFERENCES "public".screen( id ) ON DELETE CASCADE  
 );

CREATE  TABLE "public".showtime ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	film_id              integer  NOT NULL  ,
	screen_id            integer  NOT NULL  ,
	start_time           varchar(100)  NOT NULL  ,
	base_price           integer  NOT NULL  ,
	CONSTRAINT pk_showtime PRIMARY KEY ( id ),
	CONSTRAINT fk_showtime_film FOREIGN KEY ( film_id ) REFERENCES "public".film( id ) ON DELETE CASCADE  ,
	CONSTRAINT fk_showtime_screen FOREIGN KEY ( screen_id ) REFERENCES "public".screen( id ) ON DELETE CASCADE  
 );