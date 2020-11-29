-- Create tables.
DROP TABLE IF EXISTS "menu" CASCADE;
CREATE TABLE "menu"
(
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL UNIQUE,
    "price" INT NOT NULL
);

DROP TABLE IF EXISTS "orders" CASCADE;
CREATE TABLE "orders"
(
    "id"           SERIAL PRIMARY KEY,
    "menu_item_id" SERIAL,
    "table_number" INT
);

alter table "orders"
	add constraint orders_menu_id_fk
		foreign key (menu_item_id) references "menu";

-- Insert demo data.
INSERT INTO "menu" (name, price) VALUES ('BigMac', 62);
INSERT INTO "menu" (name, price) VALUES ('Cheeseburger', 32);
INSERT INTO "menu" (name, price) VALUES ('Hamburger', 28);
INSERT INTO "menu" (name, price) VALUES ('McChicken', 52);

INSERT INTO "orders" (menu_item_id, table_number) VALUES (1, 1);
INSERT INTO "orders" (menu_item_id, table_number) VALUES (2, 1);
INSERT INTO "orders" (menu_item_id, table_number) VALUES (3, 2);
INSERT INTO "orders" (menu_item_id, table_number) VALUES (4, 3);