package main

import (
    "fmt"
)

func validatePostCreateBody(post Post) error {
    if post.Title == "" {
        return fmt.Errorf("title is required")
    }
    if post.Content == "" {
        return fmt.Errorf("content is required")
    }
    if post.Author == "" {
        return fmt.Errorf("author is required")
    }
    if len(post.Title) > 255 {
        return fmt.Errorf("title cannot exceed 255 characters")
    }
    if len(post.Author) > 50 {
        return fmt.Errorf("author cannot exceed 50 characters")
    }
    return nil
}


func validatePostUpdateBody(post Post) error {
		if post.Title == "" {
				return fmt.Errorf("title is required")
		}
		if post.Content == "" {
				return fmt.Errorf("content is required")
		}
		if post.Author != "" {
			  return fmt.Errorf("author could not be updated")
		}
		if len(post.Title) > 255 {
				return fmt.Errorf("title cannot exceed 255 characters")
		}
		return nil
}

func validatePostPartialUpdateBody(post Post) error {
		if post.Author != "" {
			return fmt.Errorf("author could not be updated")
		}
		if len(post.Title) > 255 {
				return fmt.Errorf("title cannot exceed 255 characters")
		}
		return nil
}