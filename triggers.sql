DELIMITER //

CREATE TRIGGER after_insert_product_to_basket
AFTER INSERT ON basket_product
FOR EACH ROW
BEGIN
    call update_basket_total_cost(NEW.basket_id, NEW.product_id);
    
    if (SELECT count(*) FROM user_order WHERE basket_id = NEW.basket_id AND status = 'Accepted') > 0 then
		update user_order set total_cost = 1 / 0 WHERE basket_id = new.basket_id;
	end if;
END//

DELIMITER ;

DELIMITER //

CREATE TRIGGER before_insert_order
BEFORE INSERT ON user_order
FOR EACH ROW
BEGIN
	set new.total_cost = get_total_basket_cost(new.basket_id);
	if NEW.status = 'Accepted' then
		call add_order_transaction(-NEW.total_cost, NEW.user_id);
        call add_empty_basket_for_user(NEW.user_id);
    end if;
END//

DELIMITER ;

DELIMITER //

CREATE TRIGGER before_insert_user
BEFORE INSERT ON user
FOR EACH ROW
BEGIN
	INSERT INTO basket () VALUES ();
  	set NEW.basket_id = last_insert_id();
END//

DELIMITER ;

DELIMITER //

CREATE TRIGGER after_insert_user
AFTER INSERT ON user
FOR EACH ROW
BEGIN
	call add_empty_wallet_for_user(NEW.id);
END//

DELIMITER ;

DELIMITER //

CREATE TRIGGER before_update_order
BEFORE UPDATE ON user_order
FOR EACH ROW
BEGIN
	if old.status = 'Accepted' then 
		UPDATE user SET basket_id = 1 / 0 where id = 1;
    end if;
	SET NEW.total_cost = get_total_basket_cost(NEW.basket_id);
	if NEW.status = 'Accepted' then
		call add_order_transaction(-NEW.total_cost, NEW.user_id);
        call add_empty_basket_for_user(NEW.user_id);
    end if;
END//

DELIMITER ;

DELIMITER //

CREATE TRIGGER after_update_product
AFTER UPDATE ON product
FOR EACH ROW
BEGIN
	call update_baskets_cost(NEW.id, CAST(NEW.cost AS SIGNED)-CAST(OLD.cost as SIGNED));
END//

DELIMITER ;

DELIMITER //

CREATE TRIGGER after_insert_organization
AFTER INSERT ON organization
FOR EACH ROW
BEGIN
	call add_organization_manager(NEW.user_id, NEW.id);
END//

DELIMITER ;

DELIMITER //

CREATE TRIGGER after_update_wallet
AFTER UPDATE ON user_wallet
FOR EACH ROW
BEGIN
	if new.balance < 0 then
		update user set basket_id = 1 / 0 WHERE id = new.id;
	end if;
END//

DELIMITER ;

DELIMITER //

CREATE TRIGGER after_insert_transaction
AFTER INSERT ON transaction
FOR EACH ROW
BEGIN
	UPDATE user_wallet as uw 
		SET uw.balance = uw.balance + new.amount
        WHERE uw.id = new.wallet_id;
END//

DELIMITER ;