package rawsql

var SqlUserWithProfile = `
SELECT
    "users"."id",
    "users"."password",
    "users"."mobile",
    "users"."is_active",
    "users"."is_admin",
    "users"."created_at",
    "users"."updated_at",
    "users"."deleted_at",
    "Profile"."id"         AS "Profile__id",
    "Profile"."user_id"    AS "Profile__user_id",
    "Profile"."first_name" AS "Profile__first_name",
    "Profile"."last_name"  AS "Profile__last_name",
    "Profile"."avatar"     AS "Profile__avatar"
FROM
    "users"
    LEFT JOIN "profiles" "Profile"
              ON "users"."id" = "Profile"."user_id"
WHERE
    users.id = ?
ORDER BY
    "users"."id"
LIMIT 1
`
