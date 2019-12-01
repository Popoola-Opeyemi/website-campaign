CREATE TABLE "customers" (
  "id" serial NOT NULL PRIMARY KEY,
  "fullname" varchar NOT NULL,
  "mobile_number" varchar,
  "order_type" int,
  "type" int,
  "plan_type" varchar,
  "description" varchar, 
  "email_address" varchar NOT NULL,
  "created_at" timestamp not null default current_timestamp
);

CREATE TABLE "authentications" (
  "id" serial not null PRIMARY KEY,
  "role_id" bigint NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp default current_timestamp
);

CREATE TABLE "logger"(
  "id" serial not null PRIMARY key,
  "admin_id" int not null references authentications(id),
  "action" varchar not null,
  "data" json,
  "user_agent" text,
  "ip_address" character varying(200),
  "created_at" timestamp not null default current_timestamp
);

create table "gender"(
  "id" serial not null PRIMARY key,
  "name" varchar not null
);


create table "order_types"(
  "id" serial not null PRIMARY key,
  "name" varchar not null
);


INSERT into "gender" VALUES (1, 'male'), (2, 'female');
INSERT into "order_types" VALUES (1, 'web application development'), (2, 'mobile application development'), (3, 'others');

INSERT INTO "authentications" 
	VALUES (1, 1, 'o.davies@pbgdigital.co.uk','$2y$12$ymawoGhVggafmo138rcc6uA3jyUK6mvArl/SFq3fxrH206s7891fW');
  INSERT INTO "authentications"
	VALUES (2, 1, 'o.popoola@pbgdigital.co.uk','$2y$12$bEuv2Hz18t4V.9F0ItG/de2ZfBE.3uufLeB7tcaUei6r40a5DHHzq');

-- INSERT INTO "students" 
-- VALUES (1, 'Emmanuel', 'Adebayor', 'john', 'Danbo Schools', 0 , '2006-04-10' , 12, 'No 3 shadia Street Gbagada Lagos', 'John Ali', '09033003011','johnali@gmail.com', 'Opeyemi Adebayor', '093123123111' ,'adebayor1@gmail.com', 1 , 1, '2019-08-03', 'LS001' );

-- GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO sprintstar;

-- GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public to sprintstar;

-- admin 
-- secret