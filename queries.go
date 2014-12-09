package gomes

const (
	createPushTokenTable = `
    CREATE TABLE push_tokens (
        uid      VARCHAR PRIMARY KEY,
        arn      VARCHAR NOT NULL,
        arn_type VARCHAR NOT NULL,
        token    VARCHAR NOT NULL
        )
    `

	insertPushToken = `
    INSERT INTO push_tokens (uid, arn, arn_type, token)
    VALUES ($1, $2, $3, $4)
    `

	selectPushToken = `
    SELECT (uid, arn, arn_type, token) FROM push_tokens
    `

	selectPushTokenByUid = `
    SELECT (uid, arn, arn_type, token)
    FROM push_tokens
    WHERE uid = $1
    `
)
