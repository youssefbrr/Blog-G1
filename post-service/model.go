package main

import "time"

type Post struct {
    ID        string     `json:"id"`
    Title     string     `json:"title"`
    Content   string     `json:"content"`
    Author    string     `json:"author"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at,omitempty"`
    DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
