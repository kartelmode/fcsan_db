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
	if NEW.status != 'Accepted' then
		set new.total_cost = get_total_basket_cost(new.basket_id);
    end if;
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

select * from user_order;
select * from user;
select * from basket;
select * from transaction;
select * from organization_manager;
select * from user_wallet;
select * from product;
SELECT * FROM basket_product;

INSERT INTO basket_product (product_id, basket_id) VALUES (7, 15);
INSERT INTO basket_product (product_id, basket_id) VALUES (2, 14);

INSERT INTO transaction (amount, time, wallet_id) VALUES (2000, now(), 3);

INSERT INTO user_order (basket_id, user_id, status) VALUES (14, 3, 'Waiting');
UPDATE user_order SET status = 'Accepted' WHERE id = 8;

UPDATE product SET cost = 3000 WHERE id = 7;