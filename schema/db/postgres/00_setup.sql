
-- migrate up

-- schema
CREATE SCHEMA IF NOT EXISTS "scheduler";

-- functions
CREATE OR REPLACE FUNCTION "scheduler".function_updated_at()
  RETURNS TRIGGER AS $$
  BEGIN
   NEW.updated_at = now();
   RETURN NEW;
  END;
  $$ LANGUAGE 'plpgsql';

-- sections
CREATE TABLE "scheduler"."section" (
	id_section 		    TEXT PRIMARY KEY,
	"key"               TEXT NOT NULL UNIQUE,
	"name"    		    TEXT NOT NULL,
	description			TEXT NOT NULL,
	position            INTEGER NOT NULL UNIQUE,
	"active"			BOOLEAN DEFAULT TRUE NOT NULL,
	created_at			TIMESTAMP DEFAULT NOW(),
	updated_at			TIMESTAMP DEFAULT NOW()
);


-- content type
CREATE TABLE "scheduler"."content_type" (
    id_content_type 		    TEXT PRIMARY KEY,
    "key"                       TEXT NOT NULL UNIQUE,
	"name"                      TEXT NOT NULL,
	"active"			        BOOLEAN DEFAULT TRUE NOT NULL,
	created_at			        TIMESTAMP DEFAULT NOW(),
	updated_at			        TIMESTAMP DEFAULT NOW()
);

-- section contents
CREATE TABLE "scheduler"."content" (
    id_content 		            TEXT PRIMARY KEY,
    "key"                       TEXT NOT NULL UNIQUE,
	fk_section                  TEXT NOT NULL REFERENCES "scheduler"."section" (id_section),
	fk_content_type             TEXT NOT NULL REFERENCES "scheduler"."content_type" (id_content_type),
	"content"                   JSONB NOT NULL,
    position                    INTEGER NOT NULL,
	"active"			        BOOLEAN DEFAULT TRUE NOT NULL,
	created_at			        TIMESTAMP DEFAULT NOW(),
	updated_at			        TIMESTAMP DEFAULT NOW(),
	UNIQUE(fk_section, "position")
);

-- triggers
CREATE TRIGGER trigger_section_updated_at BEFORE UPDATE
  ON "scheduler"."section" FOR EACH ROW EXECUTE PROCEDURE "scheduler".function_updated_at();

CREATE TRIGGER trigger_content_updated_at BEFORE UPDATE
  ON "scheduler"."content" FOR EACH ROW EXECUTE PROCEDURE "scheduler".function_updated_at();






-- migrate down

-- triggers
DROP TRIGGER trigger_section_updated_at ON scheduler."section";
DROP TRIGGER trigger_content_updated_at ON scheduler."content";

-- tables
DROP TABLE "scheduler"."section";
DROP TABLE "scheduler"."content";

-- functions
DROP FUNCTION "scheduler".function_updated_at;

-- schema
DROP SCHEMA "scheduler";
