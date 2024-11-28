require('dotenv').config({ path: './.env' });
const host = process.env.MYSQL_HOST;
const database = process.env.MYSQL_DATABASE;
const user = process.env.MYSQL_USER;
const password = process.env.MYSQL_PASSWORD;
const port = process.env.PORT

module.exports = {
    host,
    database,
    user,
    password,
    port
};