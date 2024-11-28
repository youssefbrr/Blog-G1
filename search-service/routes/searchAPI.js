const express = require('express');
const db = require('../config/db.config');
const router = express.Router();

router.get('/search', (req, res) => {
    const { title, content, author, createdAt } = req.query;

    let query = 'SELECT * FROM Post WHERE 1=1';
    const params = [];

    if (title) {
        query += ' AND title LIKE ?';
        params.push(`%${title}%`);
    }

    if (content) {
        query += ' AND content LIKE ?';
        params.push(`%${content}%`);
    }

    if (author) {
        query += ' AND author LIKE ?';
        params.push(`%${author}%`);
    }

    if (createdAt) {
        query += ' AND DATE(created_at) = ?';
        params.push(createdAt);
    }
    const limit = req.query.limit ? parseInt(req.query.limit) : 10;
    const offset = req.query.offset ? parseInt(req.query.offset) : 0;
    query += ' LIMIT ? OFFSET ?';
    params.push(limit, offset);

    db.query(query, params, (err, results) => {
        if (err) {
            return res.status(500).json({ error: err.message });
        }
        res.json(results);
    });
});
module.exports = router;
