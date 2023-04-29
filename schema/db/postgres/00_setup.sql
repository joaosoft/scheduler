
-- migrate up

-- schema
CREATE SCHEMA IF NOT EXISTS "scheduler";

-- continent
CREATE TABLE scheduler.continent (
    id_continent serial NOT NULL PRIMARY KEY,
    "name" varchar(100) NOT NULL DEFAULT '',
    code varchar(2) NOT NULL UNIQUE,
    active bool NOT NULL DEFAULT true
);

CREATE INDEX ON scheduler.continent (code);

-- region
CREATE TABLE scheduler.region (
    id_region serial NOT NULL PRIMARY KEY,
    "name" varchar(50) NOT NULL UNIQUE,
    active bool NOT NULL DEFAULT true
);

CREATE INDEX ON scheduler.region (name);

-- country
CREATE TABLE scheduler.country (
    id_country serial NOT NULL PRIMARY KEY,
    "name" varchar(100) NOT NULL,
    code varchar(2) NOT NULL UNIQUE,
    phone_prefix varchar(20) NOT NULL DEFAULT '',
    fk_continent int4 NOT NULL REFERENCES scheduler.continent(id_continent),
    fk_region int4 NULL REFERENCES scheduler.region(id_region),
    fk_timezone int4 NULL,
    is_european_union bool NOT NULL DEFAULT false,
    lat numeric(14,11) NOT NULL DEFAULT 0,
    long numeric(14,11) NOT NULL DEFAULT 0,
    active bool NOT NULL DEFAULT true
);

CREATE INDEX ON scheduler.country (code);
CREATE INDEX ON scheduler.country (fk_continent);
CREATE INDEX ON scheduler.country (fk_region);
CREATE INDEX ON scheduler.country (fk_timezone);

-- timezone
CREATE TABLE scheduler.timezone (
    id_timezone serial NOT NULL PRIMARY KEY,
    "name" varchar(100) NOT NULL UNIQUE,
    fk_country int4 NOT NULL REFERENCES scheduler.country(id_country),
    "offset" varchar(10) NOT NULL,
    offset_dst varchar(10) NOT NULL,
    active bool NOT NULL DEFAULT true
);

CREATE INDEX ON scheduler.timezone (fk_country);
CREATE INDEX ON scheduler.timezone (name);

-- schedule_status
CREATE TABLE scheduler.schedule_status (
    id_schedule_status serial NOT NULL PRIMARY KEY,
    "name" varchar(50) NOT NULL,
    "key" varchar(50) NOT NULL UNIQUE,
    visible bool NOT NULL DEFAULT true,
    active bool NOT NULL DEFAULT true
);

CREATE INDEX ON scheduler.schedule_status (key, active);

-- schedule
CREATE TABLE scheduler.schedule (
    id_schedule serial NOT NULL PRIMARY KEY,
    hashed_id uuid NOT NULL UNIQUE,
    subject text NOT NULL,
    description text NOT NULL,
    fk_user text NOT NULL,
    fk_timezone int4 NOT NULL REFERENCES scheduler.timezone(id_timezone),
    external_id text NULL,
    start_url text NULL,
    join_url text NULL,
    fk_schedule_status int4 NOT NULL REFERENCES scheduler.schedule_status(id_schedule_status),
    fk_previous_schedule_status int4 NULL REFERENCES scheduler.schedule_status(id_schedule_status),
    status_updated_by int4 NOT NULL DEFAULT 1,
    status_updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_by INTEGER ARRAY NULL,
    updated_by int4 NOT NULL DEFAULT 1,
    updated_at timestamptz NOT NULL DEFAULT now(),
    created_by int4 NOT NULL DEFAULT 1,
    created_at timestamptz NOT NULL DEFAULT now(),
    reminded bool NOT NULL DEFAULT false
);

CREATE INDEX ON scheduler.schedule (hashed_id);
CREATE INDEX ON scheduler.schedule (fk_user);
CREATE INDEX ON scheduler.schedule (external_id);
CREATE INDEX ON scheduler.schedule (fk_schedule_status);
CREATE INDEX ON scheduler.schedule (fk_previous_schedule_status);

-- schedule_time_slot
CREATE TABLE scheduler.schedule_time_slot (
    fk_schedule int4 NOT NULL REFERENCES scheduler.schedule(id_schedule),
    "time" timestamptz NOT NULL DEFAULT now(),
    "position" int4 NOT NULL DEFAULT 1,
    active bool NOT NULL DEFAULT true,
    updated_by int4 NOT NULL DEFAULT 1,
    updated_at timestamptz NOT NULL DEFAULT now(),
    created_by int4 NOT NULL DEFAULT 1,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT schedule_time_slot_fk_schedule_position_unique UNIQUE (fk_schedule, "position")
);

CREATE INDEX ON scheduler.schedule_time_slot (fk_schedule);
CREATE INDEX ON scheduler.schedule_time_slot ("time", active, "position");

-- user
CREATE TABLE scheduler.user (
    fk_user text NOT NULL PRIMARY KEY,
    fk_timezone int4 NOT NULL REFERENCES scheduler.timezone(id_timezone),
    fk_country int4 NOT NULL REFERENCES scheduler.country(id_country),
    active bool NOT NULL DEFAULT true
);


-- migrate down

-- user
DROP TABLE scheduler.user;

-- schedule_time_slot
DROP TABLE scheduler.schedule_time_slot;

-- schedule_status
DROP TABLE scheduler.schedule_status;

-- schedule
DROP TABLE scheduler.schedule;

-- timezone
DROP TABLE scheduler.timezone;

-- country
DROP TABLE scheduler.country;

-- region
DROP TABLE scheduler.region;

-- continent
DROP TABLE scheduler.continent;

-- schema
DROP SCHEMA "scheduler";