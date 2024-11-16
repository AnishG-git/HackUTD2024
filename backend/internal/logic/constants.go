package logic

const getExistingUserQuery = `SELECT id FROM users WHERE email = $1`
