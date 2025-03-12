-- Xóa các foreign key trước
ALTER TABLE "play_list" DROP CONSTRAINT IF EXISTS "play_list_users_fkey";

ALTER TABLE "play_list_songs" DROP CONSTRAINT IF EXISTS "play_list_songs_songs_id_fkey";
ALTER TABLE "play_list_songs" DROP CONSTRAINT IF EXISTS "play_list_songs_play_list_id_fkey";

ALTER TABLE "song_lyric" DROP CONSTRAINT IF EXISTS "song_lyric_song_fkey";

ALTER TABLE "song_file_mp3" DROP CONSTRAINT IF EXISTS "song_file_mp3_song_fkey";

-- Xóa các bảng theo thứ tự ngược lại để tránh lỗi foreign key
DROP TABLE IF EXISTS "play_list_songs";

DROP TABLE IF EXISTS "play_list";

DROP TABLE IF EXISTS "song_lyric";

DROP TABLE IF EXISTS "song_file_mp3";
