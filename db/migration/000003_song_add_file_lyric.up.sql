CREATE TABLE "song_file_mp3" (
  "id" serial PRIMARY KEY,
  "title" varchar(255),
  "url" varchar(255) NOT NULL,
  "description" varchar(255),
  "is_vip" bool DEFAULT false,
  "song" serial
);

CREATE TABLE "song_lyric" (
  "id" serial PRIMARY KEY,
  "song" serial UNIQUE,
  "file" varchar(255),
  "data" jsonb
);

CREATE TABLE "play_list" (
  "id" serial PRIMARY KEY,
  "title" varchar(255),
  "users" serial,
  "thumbnail" varchar,
  "sort_description" varchar,
  "created_at" timestamp DEFAULT (now()),
  "is_private" bool DEFAULT false
);

ALTER TABLE "song_file_mp3" ADD FOREIGN KEY ("song") REFERENCES "songs" ("id") ON DELETE CASCADE;

ALTER TABLE "song_lyric" ADD FOREIGN KEY ("song") REFERENCES "songs" ("id") ON DELETE CASCADE;

CREATE TABLE "play_list_songs" (
  "play_list_id" serial,
  "songs_id" serial,
  PRIMARY KEY ("play_list_id", "songs_id")
);

ALTER TABLE "play_list_songs" ADD FOREIGN KEY ("play_list_id") REFERENCES "play_list" ("id");

ALTER TABLE "play_list_songs" ADD FOREIGN KEY ("songs_id") REFERENCES "songs" ("id");


ALTER TABLE "play_list" ADD FOREIGN KEY ("users") REFERENCES "users" ("id") ON DELETE CASCADE;
