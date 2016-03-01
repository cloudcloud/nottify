
-- version will track history of nottify versions
CREATE SEQUENCE version_id_seq;--end
CREATE TABLE IF NOT EXISTS version (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('version_id_seq'),
    version VARCHAR(15) UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE
);--end

-- generic users
CREATE SEQUENCE users_id_seq;--end
CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('users_id_seq'),
    username VARCHAR(40) NOT NULL UNIQUE,
    email VARCHAR(150) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    token VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    roles TEXT
);--end

-- allow for defined roles
CREATE SEQUENCE roles_id_seq;--end
CREATE TABLE IF NOT EXISTS roles (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('roles_id_seq'),
    name VARCHAR(50) NOT NULL UNIQUE,
    action VARCHAR(60) NOT NULL
);--end

-- track each individual artist
CREATE SEQUENCE artist_id_seq;--end
CREATE TABLE IF NOT EXISTS artist (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('artist_id_seq'),
    name TEXT NOT NULL UNIQUE,
    website TEXT,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);--end

-- track each known album
CREATE SEQUENCE album_id_seq;--end
CREATE TABLE IF NOT EXISTS album (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('album_id_seq'),
    name TEXT NOT NULL,
    artist_id INTEGER REFERENCES artist(id),
    cover_art TEXT,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);--end

-- track each song that is known
CREATE SEQUENCE song_id_seq;--end
CREATE TABLE IF NOT EXISTS song (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('song_id_seq'),
    title TEXT,
    artist_id INTEGER REFERENCES artist(id),
    album_id INTEGER REFERENCES album(id),
    length INTEGER,
    genre TEXT,
    track_number INTEGER,
    release_year INTEGER,
    filename TEXT,
    filesize INTEGER,
    id3_comment TEXT,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);--end

-- provide mechanism for storing sequences of songs
CREATE SEQUENCE playlist_id_seq;--end
CREATE TABLE IF NOT EXISTS playlist (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('playlist_id_seq'),
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);--end

-- also put songs into a playlist
CREATE SEQUENCE playlist_content_id_seq;--end
CREATE TABLE IF NOT EXISTS playlist_content (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('playlist_content_id_seq'),
    playlist_id INTEGER REFERENCES playlist(id),
    song_id INTEGER REFERENCES song(id)
);--end

-- allocate ownership to meta items
CREATE SEQUENCE ownership_id_seq;--end
CREATE TABLE IF NOT EXISTS ownership (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('ownership_id_seq'),
    user_id INTEGER REFERENCES users(id),
    item VARCHAR(20) NOT NULL,
    item_id INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);--end

-- track when modifications are through the interface
CREATE SEQUENCE history_id_seq;--end
CREATE TABLE IF NOT EXISTS history (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('history_id_seq'),
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    action TEXT
);--end

-- track as ingest processes
CREATE SEQUENCE ingest_id_seq;--end
CREATE TABLE IF NOT EXISTS ingest (
    id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('ingest_id_seq'),
    count INTEGER NOT NULL DEFAULT 0,
    completed BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);--end

-- push in the version that has now been loaded
INSERT INTO version (
    version, created_at
) VALUES (
    'v1', CURRENT_TIMESTAMP
);--end

-- function for timestamping updates
CREATE OR REPLACE FUNCTION upd_timestamp() RETURNS TRIGGER AS $upd_timestamp$
BEGIN
    IF (NEW != OLD) THEN
        NEW.updated_at = CURRENT_TIMESTAMP;
        RETURN NEW;
    END IF;
    RETURN OLD;
END;
$upd_timestamp$ LANGUAGE plpgsql;--end

-- implementation of trigger to use function
CREATE TRIGGER upd_timestamp
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE PROCEDURE upd_timestamp();--end
CREATE TRIGGER upd_timestamp
    BEFORE UPDATE ON artist
    FOR EACH ROW
    EXECUTE PROCEDURE upd_timestamp();--end
CREATE TRIGGER upd_timestamp
    BEFORE UPDATE ON album
    FOR EACH ROW
    EXECUTE PROCEDURE upd_timestamp();--end
CREATE TRIGGER upd_timestamp
    BEFORE UPDATE ON song
    FOR EACH ROW
    EXECUTE PROCEDURE upd_timestamp();--end
CREATE TRIGGER upd_timestamp
    BEFORE UPDATE ON playlist
    FOR EACH ROW
    EXECUTE PROCEDURE upd_timestamp();--end
CREATE TRIGGER upd_timestamp
    BEFORE UPDATE ON ownership
    FOR EACH ROW
    EXECUTE PROCEDURE upd_timestamp();--end
CREATE TRIGGER upd_timestamp
    BEFORE UPDATE ON history
    FOR EACH ROW
    EXECUTE PROCEDURE upd_timestamp();--end
CREATE TRIGGER upd_timestamp
    BEFORE UPDATE ON ingest
    FOR EACH ROW
    EXECUTE PROCEDURE upd_timestamp();--end

