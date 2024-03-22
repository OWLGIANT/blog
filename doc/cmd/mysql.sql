-- 删除外键
ALTER TABLE alert DROP FOREIGN KEY `fk_alert_robot`;
SHOW INDEX FROM alert;
SHOW INDEX FROM account;
DROP INDEX sub_account ON account;

-- 大后台
-- mysql -u qqlh_rich -p
-- pwd  = "W5UmPr8AD7wDMUZ4"


select * from customer;

UPDATE account SET status = 2 WHERE plat_form='coinex_usdt_swap';

source /root/update.sql;