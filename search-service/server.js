const express = require('express');
const searchRoutes = require('./searchAPI');
const app = express();
const PORT = (+process.env.PORT) || 8080;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use('/', searchRoutes);

app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});