CREATE TABLE songs (
    id SERIAl PRIMARY KEY, 
    group_name VARCHAR(255) NOT NULL, 
    song_name VARCHAR(255) NOT NULL,
    release_date DATE, 
    lyrics TEXT, 
    youtube_link TEXT, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 

);