-- checking, if user with id = 1 is admin
SELECT count(*) > 0
FROM admin
WHERE user_id = 1;

-- update user wallet
UPDATE user_wallet
	SET balance = balance + 100
    WHERE user_id = 3;
INSERT INTO transaction (amount, time, wallet_id) VALUES (100, now(), 3);

UPDATE user_wallet
	SET balance = balance - 100
    WHERE user_id = 3;
INSERT INTO transaction (amount, time, wallet_id) VALUES (-100, now(), 3);

SELECT strcmp(first_name, "BLALBALBAL")
FROM user
WHERE id = 1;

-- fetching all users, who has at least 1 order
SELECT u.*
FROM user as u
WHERE EXISTS (SELECT * FROM user_order as o WHERE o.user_id = u.id);

-- fetching user balance
SELECT balance
FROM user_wallet
WHERE user_id = 2;

SELECT t.*, avg(t.amount) over(partition by wallet_id) as avg_transaction
FROM transaction as t
WHERE t.amount > 0;

SELECT *
FROM transaction;

-- fetching user's current basket total cost 
SELECT sum(p.cost)
FROM product as p
INNER JOIN basket_product as bp
ON bp.product_id = p.id 
AND basket_id = (SELECT basket_id FROM user WHERE id = 2);

-- fetching total cost of products of organization
SELECT sum(p.cost)
FROM product AS p
where p.organization_id = 1;

-- fetching all products for organization, where user is manager
SELECT p.*
FROM product as p
WHERE organization_id = (SELECT organization_id FROM organization_manager WHERE user_id = 5);

-- fetching all products in order of user
SELECT p.*
FROM product as p 
INNER JOIN basket_product as bp
ON bp.product_id = p.id
AND basket_id = (SELECT basket_id FROM user_order WHERE user_order.user_id = 2);

-- fetching all total cost of every user's orders 
SELECT o.id, o.status, (SELECT sum(p.cost) 
			  FROM product as p
			  INNER JOIN basket_product as bp 
              ON bp.basket_id = o.basket_id
			  AND bp.product_id = p.id)
FROM user_order as o;

-- fetching all transactions for user
SELECT t.*
FROM transaction as t
WHERE t.wallet_id = (SELECT w.id FROM user_wallet as w WHERE w.user_id = 2);

-- fetching all products from current basket of user
SELECT p.*
FROM product as p
INNER JOIN basket_product as bp
ON bp.product_id = p.id
AND bp.basket_id = (SELECT u.basket_id FROM user as u WHERE u.id = 2);

-- fetching sum of all products in the current basket, grouped by organizations
SELECT p.organization_id, sum(p.cost)
FROM product AS p
INNER JOIN basket_product as bp
ON bp.product_id = p.id
AND bp.basket_id = (SELECT u.basket_id FROM user as u WHERE u.id = 2)
GROUP BY p.organization_id;

-- fetching the most expensive order of user
SELECT o.id, (SELECT sum(p.cost) 
			  FROM product as p
			  INNER JOIN basket_product as bp 
              ON bp.basket_id = o.basket_id
			  AND bp.product_id = p.id) as amount
FROM user_order as o
ORDER BY amount DESC LIMIT 1; 

SELECT o.*
FROM organization as o
INNER JOIN organization_manager as om
ON om.organization_id = o.id
AND om.user_id = 5;

SELECT *
FROM user;

SELECT *
FROM user
WHERE basket_id IS NOT NULL;