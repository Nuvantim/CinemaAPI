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
CREATE  TABLE "public".cinema ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	address              text  NOT NULL  ,
	city                 varchar(100)  NOT NULL  ,
	CONSTRAINT pk_cinema PRIMARY KEY ( id ),
	CONSTRAINT unq_cinema_name UNIQUE ( name ) 
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
CREATE  TABLE "public".user_account ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	email                varchar(100)  NOT NULL  ,
	"password"           varchar(100)  NOT NULL  ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP   ,
	CONSTRAINT pk_user_account PRIMARY KEY ( id ),
	CONSTRAINT unq_user_email UNIQUE ( email ) 
 );

CREATE  TABLE "public".booking ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	user_id              integer  NOT NULL  ,
	showtime_id          integer  NOT NULL  ,
	booking_time         timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	total_amount         double precision    ,
	CONSTRAINT pk_booking PRIMARY KEY ( id ),
	CONSTRAINT unq_booking UNIQUE ( user_id, showtime_id ) ,
	CONSTRAINT fk_booking_user_account FOREIGN KEY ( user_id ) REFERENCES "public".user_account( id ) ON DELETE CASCADE  ,
	CONSTRAINT fk_booking_showtime FOREIGN KEY ( showtime_id ) REFERENCES "public".showtime( id )   
 );

CREATE  TABLE "public".booking_seat ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	booking_id           integer  NOT NULL  ,
	seat_id              integer  NOT NULL  ,
	price_paid           double precision  NOT NULL  ,
	CONSTRAINT pk_booking_seat PRIMARY KEY ( id ),
	CONSTRAINT unq_booking_seat_booking_id UNIQUE ( booking_id, seat_id ) ,
	CONSTRAINT fk_booking_seat_booking FOREIGN KEY ( booking_id ) REFERENCES "public".booking( id ) ON DELETE CASCADE  ,
	CONSTRAINT fk_booking_seat_permission FOREIGN KEY ( seat_id ) REFERENCES "public".seat( id ) ON DELETE CASCADE  
 );

CREATE  TABLE "public".payment ( 
	id                   integer  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	booking_id           integer  NOT NULL  ,
	payment_method       varchar(100)  NOT NULL  ,
	payment_status       varchar(10)    ,
	transaction_amount  double precision  NOT NULL  ,
	payment_time         date DEFAULT CURRENT_DATE   ,
	CONSTRAINT pk_payment PRIMARY KEY ( id ),
	CONSTRAINT fk_payment_booking FOREIGN KEY ( booking_id ) REFERENCES "public".booking( id ) ON DELETE CASCADE  
 );