DELIMITER //
CREATE FUNCTION get_basket_id_for_user(user_id int) RETURNS INT DETERMINISTIC
BEGIN
	DECLARE
		basket_id INT;
    SELECT u.basket_id INTO basket_id FROM user as u WHERE u.id = user_id;
    
    RETURN basket_id;
END
//
DELIMITER ;

DELIMITER //
CREATE FUNCTION get_total_basket_cost(basket_id int) RETURNS INT DETERMINISTIC
BEGIN
	DECLARE
		total_cost INT;
    SELECT b.total_cost INTO total_cost
	  FROM basket as b
      WHERE b.id = basket_id;
    
    RETURN total_cost;
END
//
DELIMITER ;

DELIMITER //
CREATE FUNCTION get_product_cost(product_id int) RETURNS INT DETERMINISTIC
BEGIN
	DECLARE
		cost INT;
    SELECT p.cost INTO cost
	  FROM product as p
      WHERE p.id = product_id;
    
    RETURN cost;
END
//
DELIMITER ;

DELIMITER //
CREATE FUNCTION get_wallet_id(user_id int) RETURNS INT DETERMINISTIC
BEGIN
	DECLARE
		wallet_id INT;
    SELECT w.id INTO wallet_id
	  FROM user_wallet as w
      WHERE w.user_id = user_id;
    
    RETURN wallet_id;
END
//
DELIMITER ;

DELIMITER //
CREATE FUNCTION get_count_of_products_in_basket(product_id int, basket_id int) RETURNS INT DETERMINISTIC
BEGIN
	DECLARE
		cnt INT;
    SELECT count(*) INTO cnt 
		FROM basket_product as bp 
        WHERE bp.product_id = product_id AND bp.basket_id = basket_id;
    RETURN cnt;
END
//
DELIMITER ;

DELIMITER //
CREATE PROCEDURE update_basket_total_cost(basket_id int, product_id int)
BEGIN
	UPDATE basket as b
		SET b.total_cost = b.total_cost + get_product_cost(product_id) 
        WHERE b.id = basket_id;
END
//
DELIMITER ;

DELIMITER //
CREATE PROCEDURE add_order_transaction(total_cost int, user_id int)
BEGIN
	INSERT INTO transaction (amount, time, wallet_id) VALUES (total_cost, NOW(), get_wallet_id(user_id));
END
//
DELIMITER ;

DELIMITER //
CREATE PROCEDURE update_baskets_cost(product_id int, diff_cost int)
BEGIN
    UPDATE basket as b
		SET b.total_cost = (b.total_cost + (diff_cost * get_count_of_products_in_basket(product_id, b.id)))
        WHERE get_count_of_products_in_basket(product_id, b.id) > 0;
END
//
DELIMITER ;

DELIMITER //
CREATE PROCEDURE add_organization_manager(user_id_v int, organization_id_v int)
BEGIN
    INSERT INTO organization_manager
		(organization_id, user_id)
        VALUES (organization_id_v, user_id_v);
END
//
DELIMITER ;

DELIMITER //
CREATE PROCEDURE add_empty_basket_for_user(user_id int)
BEGIN
	INSERT INTO basket () VALUES ();
  UPDATE user
		SET basket_id = last_insert_id()
        WHERE id = user_id;
END
//
DELIMITER ;

DELIMITER //
CREATE PROCEDURE add_empty_wallet_for_user(user_id int)
BEGIN
	INSERT INTO user_wallet (user_id) VALUES (user_id);
END
//
DELIMITER ;