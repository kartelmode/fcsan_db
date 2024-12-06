CREATE TABLE IF NOT EXISTS user(
	id int UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY,
	first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    birth_date date NOT NULL,
    password varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    basket_id int
);

ALTER TABLE user
	ADD UNIQUE (id, basket_id);

CREATE TABLE IF NOT EXISTS user_wallet(
	id int UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY,
    balance int UNSIGNED,
    user_id int NOT NULL,
    FOREIGN KEY (user_id) 
		REFERENCES user(id)
		ON DELETE CASCADE
);

ALTER TABLE user_wallet 
	CHANGE balance 
    balance int NOT NULL DEFAULT(0);

CREATE TABLE IF NOT EXISTS transaction(
	id int UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY,
    amount int,
    time timestamp NOT NULL,
    wallet_id int NOT NULL,
    FOREIGN KEY (wallet_id) 
		REFERENCES user_wallet(id)
		ON DELETE CASCADE
);

ALTER TABLE transaction 
	CHANGE amount
		   amount int;

CREATE TABLE IF NOT EXISTS admin(
	id int UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    FOREIGN KEY (user_id) 
		REFERENCES user(id)
		ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS organization(
	id int UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(200) NOT NULL,
    description varchar(1000) NOT NULL,
    phone_number varchar(45) NOT NULL,
	user_id int NOT NULL,
    FOREIGN KEY (user_id) 
		REFERENCES user(id)
		ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS organization_request(
	id int UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY,
    description varchar(1000) NOT NULL,
    document longblob not null,
    status ENUM("Waiting", "Rejected", "Accepted") NOT NULL,
    organization_id int NOT NULL,
    FOREIGN KEY (organization_id) 
		REFERENCES organization(id)
		ON DELETE CASCADE
);

ALTER TABLE organization_request 
	CHANGE document 
		   document longblob NOT NULL;

CREATE TABLE IF NOT EXISTS organization_manager(
	id int UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY,
    organization_id int NOT NULL,
    user_id int NOT NULL,
	FOREIGN KEY (organization_id) 
		REFERENCES organization(id)
		ON DELETE CASCADE,
	FOREIGN KEY (user_id) 
		REFERENCES user(id)
		ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS product(
	id int UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(100) NOT NULL,
    description varchar(1000) NOT NULL,
    cost int UNSIGNED NOT NULL,
    organization_id int NOT NULL,
	FOREIGN KEY (organization_id) 
		REFERENCES organization(id)
		ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS basket(
	id int UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY
);

ALTER TABLE basket
	ADD total_cost INT NOT NULL DEFAULT 0;

ALTER TABLE user 
	ADD FOREIGN KEY (basket_id) 
			REFERENCES basket(id)
			ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS basket_product(
	id int UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_id int NOT NULL,
    basket_id int NOT NULL,
    FOREIGN KEY (basket_id) 
		REFERENCES basket(id)
		ON DELETE CASCADE,
	FOREIGN KEY (product_id) 
		REFERENCES product(id)
		ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_order(
	id int UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY,
    status ENUM("Waiting", "Rejected", "Accepted") NOT NULL,
    basket_id int NOT NULL,
    user_id int NOT NULL,
    FOREIGN KEY (basket_id) 
		REFERENCES basket(id)
		ON DELETE CASCADE,
	FOREIGN KEY (user_id) 
		REFERENCES user(id)
		ON DELETE CASCADE
);

ALTER TABLE user_order
	ADD UNIQUE(basket_id, user_id);
    
ALTER TABLE user_order
	ADD total_cost INT NOT NULL DEFAULT 0;