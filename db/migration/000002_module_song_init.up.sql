CREATE TABLE "albums" (
  "id" serial PRIMARY KEY,
  "mask_id" varchar(255) NOT NULL,
  "title" varchar(255) NOT NULL,
  "is_offical" bool DEFAULT true,
  "thumbnail" varchar(255),
  "sortDescription" varchar(255),
  "release_at" int
);

CREATE TABLE "songs" (
  "id" serial PRIMARY KEY,
  "mask_id" varchar(255) NOT NULL,
  "title" varchar(255) NOT NULL,
  "alias" varchar(255),
  "is_offical" bool DEFAULT true,
  "thumbnail" varchar(255),
  "thumbnailM" varchar(255),
  "duration" int NOT NULL DEFAULT 0,
  "releaseDate" int,
  "album" serial,
  "distributor" varchar(255),
  "hasLyric" bool NOT NULL DEFAULT false,
  "like" int NOT NULL DEFAULT 0,
  "listen" int NOT NULL DEFAULT 0,
  "comment" int NOT NULL DEFAULT 0,
  "audio_file" varchar
);

CREATE TABLE "artists" (
  "id" serial PRIMARY KEY,
  "mask_id" varchar(255) NOT NULL,
  "name" varchar(255),
  "spotlight" boolean,
  "alias" varchar(255),
  "thumbnail" varchar(255),
  "thumbnailM" varchar(255),
  "playlistId" varchar(20),
  "totalFollow" int NOT NULL DEFAULT 0
);

CREATE TABLE "composers" (
  "id" serial PRIMARY KEY,
  "mask_id" varchar(255) NOT NULL,
  "name" varchar(255),
  "spotlight" boolean,
  "alias" varchar(255),
  "thumbnail" varchar(255),
  "cover" varchar(255)
);


CREATE TABLE "genres" (
  "id" serial PRIMARY KEY,
  "mask_id" varchar(255) NOT NULL,
  "title" varchar(255) NOT NULL,
  "name" varchar(255) NOT NULL,
  "alias" varchar(255)
);

CREATE TABLE "songs_artists" (
  "songs_id" serial,
  "artists_id" serial,
  PRIMARY KEY ("songs_id", "artists_id")
);

ALTER TABLE "songs_artists" ADD FOREIGN KEY ("songs_id") REFERENCES "songs" ("id");

ALTER TABLE "songs_artists" ADD FOREIGN KEY ("artists_id") REFERENCES "artists" ("id");


CREATE TABLE "songs_genres" (
  "songs_id" serial,
  "genres_id" serial,
  PRIMARY KEY ("songs_id", "genres_id")
);

ALTER TABLE "songs_genres" ADD FOREIGN KEY ("songs_id") REFERENCES "songs" ("id");

ALTER TABLE "songs_genres" ADD FOREIGN KEY ("genres_id") REFERENCES "genres" ("id");


CREATE TABLE "songs_composers" (
  "songs_id" serial,
  "composers_id" serial,
  PRIMARY KEY ("songs_id", "composers_id")
);

ALTER TABLE "songs_composers" ADD FOREIGN KEY ("songs_id") REFERENCES "songs" ("id");

ALTER TABLE "songs_composers" ADD FOREIGN KEY ("composers_id") REFERENCES "composers" ("id");


ALTER TABLE "songs" ADD FOREIGN KEY ("album") REFERENCES "albums" ("id") ON DELETE SET NULL;

CREATE TABLE "albums_artists" (
  "albums_id" serial,
  "artists_id" serial,
  PRIMARY KEY ("albums_id", "artists_id")
);

ALTER TABLE "albums_artists" ADD FOREIGN KEY ("albums_id") REFERENCES "albums" ("id");

ALTER TABLE "albums_artists" ADD FOREIGN KEY ("artists_id") REFERENCES "artists" ("id");


CREATE TABLE "albums_genres" (
  "albums_id" serial,
  "genres_id" serial,
  PRIMARY KEY ("albums_id", "genres_id")
);

ALTER TABLE "albums_genres" ADD FOREIGN KEY ("albums_id") REFERENCES "albums" ("id");

ALTER TABLE "albums_genres" ADD FOREIGN KEY ("genres_id") REFERENCES "genres" ("id");