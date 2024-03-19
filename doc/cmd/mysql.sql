-- 删除外键
ALTER TABLE alert DROP FOREIGN KEY `fk_alert_robot`;
SHOW INDEX FROM alert;


-- 大后台
-- mysql -u qqlh_rich -p
-- pwd  = "W5UmPr8AD7wDMUZ4"


select * from customer;