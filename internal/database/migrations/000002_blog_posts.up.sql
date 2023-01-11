CREATE TABLE blog_posts(
    id SERIAL PRIMARY KEY UNIQUE,
    title VARCHAR (255),
    type VARCHAR (255),
    description VARCHAR (2500)
    
);