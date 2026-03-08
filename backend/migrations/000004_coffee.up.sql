CREATE TABLE coffee (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  roast_level SMALLINT CHECK (roast_level BETWEEN 1 AND 5),
  roaster_name VARCHAR(255) NOT NULL,
  origin_country VARCHAR(255) NOT NULL,
  processing VARCHAR(255) NOT NULL,
  varietal VARCHAR(255) NOT NULL DEFAULT 'Arabica',
  image VARCHAR(255) NOT NULL,
  brew_date TIMESTAMPTZ,
  description TEXT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);


CREATE TABLE coffee_cup (
  id SERIAL PRIMARY KEY,
  coffee_id INTEGER REFERENCES coffee(id),
  temperature SMALLINT,
  date_drank TIMESTAMPTZ,
  acidity SMALLINT CHECK (acidity BETWEEN 1 and 10),
  body SMALLINT CHECK (body BETWEEN 1 and 10),
  sweetness SMALLINT CHECK (sweetness BETWEEN 1 AND 10),
  water_type VARCHAR(255),
  grind_size SMALLINT CHECK (grind_size BETWEEN 1 and 30),
  method VARCHAR(255),
  rating SMALLINT CHECK (rating BETWEEN 1 and 10)
);
