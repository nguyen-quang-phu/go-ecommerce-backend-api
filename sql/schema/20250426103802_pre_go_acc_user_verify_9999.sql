-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_verify_9999`(
`verify_id` int not null auto_increment,
`verify_otp` varchar(6) not null,
  `verify_key` varchar(255) not null,
  `verify_key_hash` varchar(255) not null,
  `verify_type` int default '1',
  `is_verified` int default '0',
  `verify_created_at` timestamp null default current_timestamp,
  `verify_updated_at` timestamp null default current_timestamp on update current_timestamp,
  primary key (`verify_id`),
  unique key `unique_verify_key`(`verify_key`),
  key `idex_verify_otp` (`verify_otp`)
) Engine= InnoDB default charset=utf8mb4 collate=utf8mb4_unicode_ci comment='account_user_verify';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
