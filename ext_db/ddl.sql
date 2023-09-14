CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    deposit_amount FLOAT,
    jwt_token VARCHAR(255)
);

-- Create Games table
CREATE TABLE Video_Games (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    description VARCHAR(255),
    availability BOOLEAN,
    rental_cost FLOAT
);

INSERT INTO Video_Games (title, description, availability, rental_cost)
VALUES
    ('The Witcher 3: Wild Hunt', 'Open-world RPG', true, 4.99),
    ('Grand Theft Auto V', 'Open-world action-adventure', true, 3.99),
    ('FIFA 22', 'Football simulation', true, 2.99),
    ('Red Dead Redemption 2', 'Western action-adventure', true, 4.49),
    ('Minecraft', 'Sandbox building', true, 2.49),
    ('Call of Duty: Warzone', 'Battle royale shooter', true, 1.99),
    ('Among Us', 'Social deduction', true, 1.49),
    ('Cyberpunk 2077', 'Open-world RPG', true, 3.99),
    ('Animal Crossing: New Horizons', 'Life simulation', true, 2.99),
    ('The Legend of Zelda: Breath of the Wild', 'Action-adventure', true, 4.99),
    ('Fortnite', 'Battle royale shooter', true, 0.99),
    ('Rocket League', 'Soccer with rocket-powered cars', true, 1.49),
    ('Assassins Creed Valhalla', 'Open-world action RPG', true, 3.99),
    ('Overwatch', 'Team-based shooter', true, 2.49),
    ('Hades', 'Roguelike action', true, 1.99),
    ('The Elder Scrolls V: Skyrim', 'Open-world RPG', true, 4.49),
    ('Super Mario Odyssey', 'Platformer', true, 2.99),
    ('Valorant', 'Tactical shooter', true, 0.99),
    ('Rainbow Six Siege', 'Tactical shooter', true, 1.99),
    ('Mortal Kombat 11', 'Fighting game', true, 3.49);
    
-- Create history table
CREATE TABLE Histories (
    id SERIAL PRIMARY KEY,
   	user_id INTEGER REFERENCES Users(id),
   	video_game_id INTEGER REFERENCES Video_Games(id),
   	start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   	due_date TIMESTAMP,
   	status VARCHAR(50)
);

-- status = in-progress/done
