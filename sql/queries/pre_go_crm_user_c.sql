-- name: GetUserByEmail :one
SELECT
    usr_email,
    usr_id
FROM pre_go_crm_user_c
WHERE usr_email = $1
LIMIT 1;

-- name: UpdateUserStatusByUserId :exec
UPDATE pre_go_crm_user_c
SET
    usr_status = $1,
    usr_updated_at = $2
WHERE usr_id = $3;
