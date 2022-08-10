-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    id           uuid      DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at   timestamp DEFAULT now(),
    updated_at   timestamp,
    first_name   text,
    last_name    text,
    username     text,
    password_hash     text NOT NULL,
    status       text NOT NULL,
    role         text NOT NULL,
    is_anonymous boolean   DEFAULT false
-- TODO: add isGoogle, isApple, etc.
);

CREATE TABLE IF NOT EXISTS news
(
    id                   uuid      DEFAULT gen_random_uuid() PRIMARY KEY,
    created_at           timestamp DEFAULT now(),
    updated_at           timestamp,
    hide_metrics         boolean   DEFAULT false,
    views_count          bigint    DEFAULT 0,
    reactions_count      bigint    DEFAULT 0,
    hide_views_count     boolean   DEFAULT false,
    hide_reactions_count boolean   DEFAULT false,
    media                jsonb,
    author_id            uuid REFERENCES users (id),
    -- used for full-text search (might later become deprecated, if moved to some other search engine)
    -- tsv_text_[locale-iso2]
    tsv_text_en          tsvector,
    -- ru used for both ru and ua here (postgres only supports one locale)
    tsv_text_ru          tsvector
);

CREATE TABLE IF NOT EXISTS email
(
    id          uuid    DEFAULT gen_random_uuid() PRIMARY KEY,
    is_verified boolean DEFAULT false,
    is_primary  boolean DEFAULT false,
    email       text NOT NULL,
    user_id     uuid REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS phone
(
    id                uuid    DEFAULT gen_random_uuid() PRIMARY KEY,
    subscriber_number text NOT NULL,
    extension         text,
    country_code      text,
    is_verified       boolean DEFAULT false,
    is_primary        boolean DEFAULT false,
    user_id           uuid REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS country
(
    code character(3) PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE IF NOT EXISTS location
(
    id           uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    country_code text REFERENCES country (code),
    region       text,
    city         text,
    street       text
);

CREATE TABLE IF NOT EXISTS user_location
(
    id          uuid      DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id     uuid REFERENCES users (id),
    location_id uuid REFERENCES location (id),
    created_at  timestamp DEFAULT now()
);

-- indexes
-- new search indexes should be created for other locales
CREATE INDEX news_text_search_en_idx ON news USING GIN (tsv_text_en);
CREATE INDEX news_text_search_ru_idx ON news USING GIN (tsv_text_ru);

-- +migrate Down
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS news CASCADE;
DROP TABLE IF EXISTS email CASCADE;
DROP TABLE IF EXISTS phone CASCADE;
DROP TABLE IF EXISTS user_location CASCADE;
DROP TABLE IF EXISTS location CASCADE;
DROP TABLE IF EXISTS country CASCADE;