CREATE TABLE roast (
  id SERIAL PRIMARY KEY,
  roast_level SMALLINT CHECK (roast_level BETWEEN 1 AND 5),
  roaster_name VARCHAR(255) NOT NULL,
  roast_date TIMESTAMPTZ NOT NULL,
  image VARCHAR(255),
  coffee_id INTEGER REFERENCES coffee(id)
);

INSERT INTO roast (roast_level, roaster_name, roast_date, image, coffee_id)
SELECT roast_level, roaster_name, brew_date, image, id FROM coffee;

ALTER TABLE roast ALTER COLUMN coffee_id SET NOT NULL;

ALTER TABLE coffee DROP COLUMN IF EXISTS roast_level;
ALTER TABLE coffee DROP COLUMN IF EXISTS roaster_name;
ALTER TABLE coffee DROP COLUMN IF EXISTS brew_date;
ALTER TABLE coffee DROP COLUMN IF EXISTS image;

ALTER TABLE coffee_cup ADD COLUMN roast_id INTEGER;

UPDATE coffee_cup cc
SET roast_id = r.id
FROM roast r
WHERE r.coffee_id = cc.coffee_id;

ALTER TABLE coffee_cup ALTER COLUMN roast_id SET NOT NULL;
ALTER TABLE coffee_cup DROP CONSTRAINT coffee_cup_coffee_id_fkey;
ALTER TABLE coffee_cup DROP COLUMN coffee_id;

ALTER TABLE coffee_cup
ADD CONSTRAINT coffee_cup_roast_id_fkey
FOREIGN KEY (roast_id) REFERENCES roast(id);

