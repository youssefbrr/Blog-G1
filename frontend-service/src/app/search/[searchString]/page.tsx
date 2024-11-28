'use client'
import React from 'react'
import { useParams } from 'next/navigation'

import { useState, useEffect, useCallback } from 'react'
import { Post } from '../post.types'
import PostCard from '@/app/posts/_components/PostCard/PostCard'
import { getSearchPostsApi } from '@/lib/api'

interface SearchProps {}

const Search = ({}: SearchProps) => {
  const { searchString } = useParams()
  const [articles, setArticles] = useState<Post[]>([])
  const [page, setPage] = useState(1)
  const [hasMore, setHasMore] = useState(true)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const fetchArticles = useCallback(async () => {
    try {
      setLoading(true)
      setError(null)
      getSearchPostsApi({
        title: searchString,
        content: searchString,
        author: searchString,
        page,
      }).then((data) => {
        setArticles(data)
        setHasMore(data.length > 0)
        setPage((prevPage) => prevPage + 1)
      })
    } catch (error: any) {
      setError(error instanceof Error ? error.message : 'An unknown error occurred')
      setHasMore(false)
    } finally {
      setLoading(false)
    }
  }, [page, searchString])

  // Initial fetch
  useEffect(() => {
    fetchArticles()
  }, [])

  return (
    <div className="max-w-3xl mx-auto py-8 px-4">
      {error && (
        <div
          className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4"
          role="alert"
        >
          <span className="block sm:inline">{error}</span>
        </div>
      )}

      <div className="space-y-6">
        {articles.map((article: any, index) => (
          <PostCard post={article} key={index} />
        ))}
      </div>

      {loading && (
        <div className="flex justify-center items-center mt-6">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900" />
          <p className="ml-3 text-gray-700">Loading more articles...</p>
        </div>
      )}

      {!hasMore && articles.length > 0 && (
        <p className="text-center mt-6 text-gray-500">No more articles to load</p>
      )}

      {!hasMore && articles.length === 0 && (
        <p className="text-center mt-6 text-gray-500">No Posts found</p>
      )}
    </div>
  )
}

export default Search
