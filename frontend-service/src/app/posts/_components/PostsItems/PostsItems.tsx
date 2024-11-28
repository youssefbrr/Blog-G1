'use client'
import React, { useEffect, useState } from 'react'
import PostCard from '../PostCard/PostCard'
import { useRouter } from 'next/navigation'
import { deletePostApi, getPostsApi } from '@/lib/api'

export interface IPost {
  id: number
  title: string
  content: string
  author: string
  createdAt: string
  created_at: string
}

const PostsItems = () => {
  const [posts, setPosts] = useState<IPost[]>([])
  const [loading, setLoading] = useState(true)
  const router = useRouter()
  useEffect(() => {
    setLoading(true)
    getPostsApi({ limit: 100 }).then((data) => {
      setPosts(data)
      setLoading(false)
    })
  }, [])

  const deletePost = async (id: number) => {
    const filteredPosts = posts.filter((post) => post.id !== id)
    setPosts(filteredPosts)
    try {
      await deletePostApi(id)
    } catch (error) {
      setPosts(posts)
      alert('Error deleting post')
    }
  }

  const editPost = (id: number) => openPost(id)

  const openPost = (id: number) => router.push(`/posts/${id}`)

  return (
    <div className="flex flex-col gap-3">
      {loading
        ? Array.from({ length: 10 }).map((_, index) => (
            <div
              key={index}
              className="flex items-center justify-center h-40 bg-gray-300 rounded-lg animate-pulse light:bg-gray-700"
            ></div>
          ))
        : posts.map((post) => (
            <PostCard
              key={post.id}
              post={post}
              actions={{
                delete: () => deletePost(post.id),
                edit: () => editPost(post.id),
                open: () => openPost(post.id),
              }}
            />
          ))}
    </div>
  )
}

export default PostsItems
