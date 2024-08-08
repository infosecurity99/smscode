alter table users
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;
