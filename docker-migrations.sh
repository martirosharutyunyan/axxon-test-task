DB_DSN="host=$DB_HOST user=$DB_USERNAME password=$DB_PASSWORD dbname=$DB_NAME sslmode=$DB_SSL"

goose postgres "$DB_DSN" up