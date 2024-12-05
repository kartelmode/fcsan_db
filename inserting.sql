INSERT INTO user (first_name, last_name, birth_date, password, email) 
	VALUES ("Aliaksandr", "Loseu", "2005-02-17", "admin", "losevgg@gmail.com");
INSERT INTO admin (user_id) 
	VALUES (1);
INSERT INTO basket () VALUES ();
INSERT INTO user (first_name, last_name, birth_date, password, email, basket_id)
	VALUES ("username1", "userlastname1", "2002-02-02", "lol", "user1@gmail.com", 1);

INSERT INTO basket () VALUES ();
INSERT INTO user (first_name, last_name, birth_date, password, email, basket_id)
	VALUES ("username2", "userlastname2", "2002-03-04", "lol2", "user2@gmail.com", 2);
INSERT INTO basket () VALUES ();
INSERT INTO user (first_name, last_name, birth_date, password, email, basket_id)
	VALUES ("username3", "userlastname3", "2002-03-03", "lol3", "user3@gmail.com", 3);

INSERT INTO user_wallet (user_id) VALUES (1);
INSERT INTO user_wallet (user_id) VALUES (2);
INSERT INTO user_wallet (user_id) VALUES (3);
INSERT INTO user_wallet (user_id) VALUES (4);

INSERT INTO basket () VALUES ();
INSERT INTO user (first_name, last_name, birth_date, password, email, basket_id)
	VALUES ("organizator1", "organizatorlastname1", "2000-03-04", "zheskiy_org", "organizator@gmail.com", 4);
INSERT INTO user_wallet (user_id) VALUES (4);

INSERT INTO organization (name, description, phone_number, user_id) 
	VALUES ("Apple", "BIG TECH SCAM CORPORATION", "8(800)5553535", 5);
INSERT INTO organization_request (organization_id, description, document, status)
	VALUES (1, "ДОБАВЬТЕ ПЖПЖПЖПЖПЖПЖ", LOAD_FILE("C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/CV_Aliaksandr_Loseu_RU.pdf"),  "Accepted");
INSERT INTO organization_manager (user_id, organization_id)
	VALUES (5, 1);

INSERT INTO product (name, description, cost, organization_id)
	VALUES ("Iphone 17 super pro max (not scam)", "Лучший телефон на планете Земля. Все ведроиды -- шлак, покупайте нашу продукцию", 2000, 1);
INSERT INTO product (name, description, cost, organization_id)
	VALUES ("Iphone 18 super pro max (not scam)", "Лучший телефон на планете Земля. Все ведроиды -- шлак, покупайте нашу продукцию", 3000, 1);
INSERT INTO product (name, description, cost, organization_id)
	VALUES ("Iphone 19 super pro max (not scam)", "Лучший телефон на планете Земля. Все ведроиды -- шлак, покупайте нашу продукцию", 4000, 1);
INSERT INTO product (name, description, cost, organization_id)
	VALUES ("Iphone 20 super pro max (not scam)", "Лучший телефон на планете Земля. Все ведроиды -- шлак, покупайте нашу продукцию", 5000, 1);
INSERT INTO product (name, description, cost, organization_id)
	VALUES ("Iphone 21 super pro max (not scam)", "Лучший телефон на планете Земля. Все ведроиды -- шлак, покупайте нашу продукцию", 6000, 1);
INSERT INTO product (name, description, cost, organization_id)
	VALUES ("Iphone 22 super pro max (not scam)", "Лучший телефон на планете Земля. Все ведроиды -- шлак, покупайте нашу продукцию", 7000, 1);
INSERT INTO product (name, description, cost, organization_id)
	VALUES ("Iphone 23 super pro max (not scam)", "Лучший телефон на планете Земля. Все ведроиды -- шлак, покупайте нашу продукцию", 8000, 1);
INSERT INTO product (name, description, cost, organization_id)
	VALUES ("Iphone 24 super pro max (not scam)", "Лучший телефон на планете Земля. Все ведроиды -- шлак, покупайте нашу продукцию", 9000, 1);
    
INSERT INTO basket_product (basket_id, product_id)
	VALUES (1, 1);
INSERT INTO basket_product (basket_id, product_id)
	VALUES (1, 2);
INSERT INTO basket_product (basket_id, product_id)
	VALUES (2, 3);
INSERT INTO user_order (status, basket_id, user_id)
	VALUES ("Accepted", 1, 2);
INSERT INTO transaction 
	(wallet_id, amount, time)
	VALUES (2, (
				SELECT sum(p.cost) 
				FROM product as p
                INNER JOIN basket_product as bp
                ON bp.product_id = p.id
                AND basket_id = (SELECT basket_id FROM user_order WHERE user_order.user_id = 2)), NOW());
    
INSERT INTO basket () VALUES ();
UPDATE user 
	SET basket_id = 5
    WHERE id = 2;
INSERT INTO basket_product (basket_id, product_id)
	VALUES (5, 4);
INSERT INTO basket_product (basket_id, product_id)
	VALUES (5, 6);
INSERT INTO basket_product (basket_id, product_id)
	VALUES (3, 5);