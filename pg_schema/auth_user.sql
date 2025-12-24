CREATE SCHEMA IF NOT EXISTS "public";

CREATE  TABLE "public".permission ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_permission PRIMARY KEY ( id ),
	CONSTRAINT permission_name UNIQUE ( name ) 
 );

CREATE  TABLE "public"."role" ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_role PRIMARY KEY ( id ),
	CONSTRAINT role_name UNIQUE ( name ) 
 );

CREATE  TABLE "public".role_permission ( 
	id_role              bigint  NOT NULL  ,
	id_permission        bigint  NOT NULL  ,
	CONSTRAINT fk_role_permission_permission FOREIGN KEY ( id_permission ) REFERENCES "public".permission( id ) ON DELETE CASCADE  ,
	CONSTRAINT fk_role_permission_role FOREIGN KEY ( id_role ) REFERENCES "public"."role"( id ) ON DELETE CASCADE  
 );

CREATE  TABLE "public".user_account ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	name                 varchar(100)  NOT NULL  ,
	email                varchar(100)  NOT NULL  ,
	"password"           varchar(100)  NOT NULL  ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP   ,
	CONSTRAINT pk_user_account PRIMARY KEY ( id )
 );

CREATE  TABLE "public".user_profile ( 
	id                   bigint  NOT NULL GENERATED  BY DEFAULT AS IDENTITY ,
	user_id              bigint  NOT NULL  ,
	age                  integer    ,
	phone                integer    ,
	district             varchar(100)    ,
	city                 varchar(100)    ,
	country              varchar(100)    ,
	created_at           timestamptz DEFAULT CURRENT_TIMESTAMP   ,
	CONSTRAINT pk_user_profile PRIMARY KEY ( id ),
	CONSTRAINT fk_user_profile_user_account FOREIGN KEY ( user_id ) REFERENCES "public".user_account( id ) ON DELETE CASCADE  
 );

CREATE  TABLE "public".user_role ( 
	id_user              bigint  NOT NULL  ,
	id_role              bigint  NOT NULL  ,
	CONSTRAINT fk_user_role_role FOREIGN KEY ( id_role ) REFERENCES "public"."role"( id ) ON DELETE CASCADE  ,
	CONSTRAINT fk_user_role_user_account FOREIGN KEY ( id_user ) REFERENCES "public".user_account( id ) ON DELETE CASCADE  
 );

INSERT INTO "public".permission( name, created_at ) VALUES ( 'handle permission', '2025-08-08 05:02:38 PM');
INSERT INTO "public".permission( name, created_at ) VALUES ( 'handle role', '2025-08-08 05:02:48 PM');
INSERT INTO "public".permission( name, created_at ) VALUES ( 'handle genre', '2025-11-09 08:08:30 PM');
INSERT INTO "public".permission( name, created_at ) VALUES ( 'handle film', '2025-11-09 08:08:53 PM');
INSERT INTO "public".permission( name, created_at ) VALUES ( 'handle cinema', '2025-11-09 08:09:09 PM');
INSERT INTO "public".permission( name, created_at ) VALUES ( 'handle screen type', '2025-11-09 08:09:18 PM');
INSERT INTO "public".permission( name, created_at ) VALUES ( 'handle screen', '2025-11-09 08:09:25 PM');
INSERT INTO "public".permission( name, created_at ) VALUES ( 'handle seat', '2025-11-09 08:09:41 PM');
INSERT INTO "public".permission( name, created_at ) VALUES ( 'handle showtime', '2025-11-09 08:09:49 PM');
INSERT INTO "public"."role"( name, created_at ) VALUES ( 'admin', '2025-08-08 05:03:16 PM');
INSERT INTO "public".role_permission( id_role, id_permission ) VALUES ( 1, 1);
INSERT INTO "public".role_permission( id_role, id_permission ) VALUES ( 1, 2);
INSERT INTO "public".role_permission( id_role, id_permission ) VALUES ( 1, 3);
INSERT INTO "public".role_permission( id_role, id_permission ) VALUES ( 1, 4);
INSERT INTO "public".role_permission( id_role, id_permission ) VALUES ( 1, 5);
INSERT INTO "public".role_permission( id_role, id_permission ) VALUES ( 1, 6);
INSERT INTO "public".role_permission( id_role, id_permission ) VALUES ( 1, 7);
INSERT INTO "public".role_permission( id_role, id_permission ) VALUES ( 1, 8);
INSERT INTO "public".role_permission( id_role, id_permission ) VALUES ( 1, 9);
INSERT INTO "public".user_account( name, email, "password", created_at ) VALUES ( 'nuvantim', 'nuvantim@gmail.com', '$2a$10$7CfzVVG0RoADtUxfWRC9COjj01d/fk2vlZTz.TsuFpw3HoOZlwVeG', '2025-08-08 05:06:52 PM');
INSERT INTO "public".user_profile( user_id, age, phone, district, city, country, created_at ) VALUES ( 1, 23, 891290202, 'nguling', 'pasuruan', 'indonesia', '2025-08-08 05:06:52 PM');
INSERT INTO "public".user_role( id_user, id_role ) VALUES ( 1, 1);