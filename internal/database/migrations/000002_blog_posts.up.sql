CREATE TABLE blog_posts(
    id SERIAL PRIMARY KEY UNIQUE,
    title VARCHAR (255),
    short_description VARCHAR (255),
    long_description VARCHAR (255)
    
);