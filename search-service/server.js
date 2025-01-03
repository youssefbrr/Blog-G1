const express = require('express');
const cors = require('cors');
const searchRoutes = require('./routes/searchAPI');
const config = require('./utils/config');

const app = express();
const PORT = (+config.port) || 8080;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(
    cors({origin: '*'})
);
app.use('/', searchRoutes);

app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});