ALTER TABLE coffee ADD COLUMN roast_level SMALLINT CHECK (roast_level BETWEEN 1 AND 5);
ALTER TABLE coffee ADD COLUMN roaster_name VARCHAR(255) NOT NULL;
ALTER TABLE coffee ADD COLUMN roast_date TIMESTAMPTZ NOT NULL;
ALTER TABLE coffee ADD COLUMN image VARCHAR(255);

UPDATE coffee c
SET roast_level = r.roast_level,
    roaster_name = r.roaster_name,
    roast_date = r.roast_date,
    image = r.image
FROM roast r
WHERE r.coffee_id = c.id;

ALTER TABLE coffee_cup ADD COLUMN coffee_id INTEGER;

UPDATE coffee_cup cc
SET coffee_id = r.coffee_id
FROM roast r
WHERE r.id = cc.roast_id;

ALTER TABLE coffee_cup ALTER COLUMN coffee_id SET NOT NULL;

ALTER TABLE coffee_cup
ADD CONSTRAINT coffee_cup_coffee_id_fkey
FOREIGN KEY (coffee_id) REFERENCES coffee(id);

ALTER TABLE coffee_cup DROP COLUMN roast_id;


DROP TABLE roast;
