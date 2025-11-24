CREATE SCHEMA IF NOT EXISTS "public";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE  TABLE "public".booking ( 
	id                   bigint  NOT NULL  ,
	user_id              bigint  NOT NULL  ,
	showtime_id          bigint  NOT NULL  ,
	booking_time         timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	total_amount         double precision DEFAULT 0 NOT NULL  ,
	CONSTRAINT pk_booking PRIMARY KEY ( id ),
	CONSTRAINT unq_booking UNIQUE ( user_id, showtime_id ) 
 );

CREATE  TABLE "public".booking_seat ( 
	id                   uuid DEFAULT uuid_generate_v4() NOT NULL  ,
	booking_id           bigint  NOT NULL  ,
	seat_id              bigint  NOT NULL  ,
	price_paid           double precision  NOT NULL  ,
	CONSTRAINT booking_seat_pkey PRIMARY KEY ( id ),
	CONSTRAINT unq_booking_seat_booking_id UNIQUE ( booking_id, seat_id ) ,
	CONSTRAINT fk_booking_seat_booking FOREIGN KEY ( booking_id ) REFERENCES "public".booking( id ) ON DELETE CASCADE  
 );

CREATE  TABLE "public".payment ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	user_id              bigint  NOT NULL  ,
	booking_id           bigint  NOT NULL  ,
	payment_method       varchar(100)  NOT NULL  ,
	payment_status       varchar(10)  NOT NULL  ,
	transaction_amount   double precision  NOT NULL  ,
	payment_time         timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_payment PRIMARY KEY ( id )
 );
