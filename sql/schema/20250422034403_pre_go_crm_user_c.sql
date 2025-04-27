-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_crm_user_c (
  usr_id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
  usr_email VARCHAR(30) NOT NULL DEFAULT '' COMMENT 'Email',
  usr_phone VARCHAR(15) NOT NULL DEFAULT '' COMMENT 'Phone Number',
  usr_username VARCHAR(30) NOT NULL DEFAULT '' COMMENT 'Username',
  usr_password VARCHAR(32) NOT NULL DEFAULT '' COMMENT 'Password',
  usr_created_at INT(11) NOT NULL DEFAULT '0' COMMENT 'Creation Time',
  usr_updated_at INT(11) NOT NULL DEFAULT '0' COMMENT 'Update time',
  user_create_ip_at VARCHAR(12) NOT NULL DEFAULT '' COMMENT 'Creation IP',
  usr_last_login_at INT(11) NOT NULL DEFAULT '0' COMMENT 'Last Login Time',
  usr_login_times INT(11) NOT NULL DEFAULT '0' COMMENT 'Login Times',
  usr_status TINYINT(
    1
  ) NOT NULL DEFAULT '0' COMMENT 'Status 1:enable, 0:disable, -1:deleted',
  PRIMARY KEY (usr_id),
  KEY IDX_EMAIL (usr_email),
  KEY IDX_PHONE (usr_phone),
  KEY IDX_USERNAME (usr_username)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Account';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_crm_user_c;
-- +goose StatementEnd
