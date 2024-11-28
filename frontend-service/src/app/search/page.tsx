'use client'
import { useEffect, useState, useCallback } from 'react'
import { Post } from './post.types'
import PostCard from '../posts/_components/PostCard/PostCard'

const Search = () => {
  const [articles, setArticles] = useState<Post[]>([])
  const [query, setQuery] = useState('')
  const [page, setPage] = useState(1)
  const [hasMore, setHasMore] = useState(true)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const fetchArticles = useCallback(async () => {
    let data
    const url = `https://blog1-search-posts.tamtasks.com/search?title=${query}&content=${query}&author=${query}&page=${page}`
    setLoading(true)
    setError(null)

    try {
      const response = await fetch(url, {
        method: 'GET',
        headers: {
          Accept: '*',
          'Content-Type': 'application/json',
        },
      })

      if (!response.ok) {
        throw new Error(`Response status: ${response.status}`)
      }

      data = await response.json()
      console.log(data)

      setArticles(data)
      setHasMore(true)
      setPage((prevPage) => prevPage + 1)
    } catch (error: any) {
      console.error(error.message)
      setError(error instanceof Error ? error.message : 'An unknown error occurred')
      setHasMore(false)
    } finally {
      setLoading(false)
    }
  }, [page, query])

  useEffect(() => {
    fetchArticles()
  }, [])
  console.log()
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
        {articles.map((article:any, index) => (
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
        <p className="text-center mt-6 text-gray-500">No articles found</p>
      )}
    </div>
  )
}

export default Search
