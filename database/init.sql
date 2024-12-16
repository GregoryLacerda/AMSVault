CREATE DATABASE IF NOT EXISTS amsvault;

USE amsvault;

CREATE TABLE IF NOT EXISTS Users (
    ID SERIAL PRIMARY KEY,         
    Name VARCHAR(255) NOT NULL,    
    Email VARCHAR(255) NOT NULL UNIQUE, 
    Password VARCHAR(255) NOT NULL 
);

CREATE TABLE IF NOT EXISTS Stories (
    ID SERIAL PRIMARY KEY,
    User VARCHAR(255) NOT NULL,
    Name VARCHAR(255) NOT NULL,
    Source VARCHAR(255),
    Description TEXT,
    Season INT,
    Episode INT,
    Volume INT,
    Chapter INT,
    Status VARCHAR(255) NOT NULL,
    MediumImage VARCHAR(255),
    LargeImage VARCHAR(255),
    CreatedAt VARCHAR(255),
    UpdatedAt VARCHAR(255),
    DeletedAt VARCHAR(255)
);
