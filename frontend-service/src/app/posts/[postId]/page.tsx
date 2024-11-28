'use client'

import React, { useEffect, useState } from 'react'
import { getPostApi } from '@/lib/api'

interface PostDetailsProps {
  params: {
    postId: string
  }
}
interface IPost {
  id: number
  title: string
  content: string
  author: string
  created_at: string
}
const PostDetails = ({ params }: PostDetailsProps) => {
  const [post, setPost] = useState<IPost>()
  const [postId, setPostId] = useState<string | undefined>(undefined)

  useEffect(() => {
    ;(async () => {
      const resolvedParams = await params
      setPostId(resolvedParams.postId)

      if (resolvedParams.postId) {
        const data = await getPostApi(resolvedParams.postId)
        setPost(data)
      }
    })()
  }, [params])

  return (
    <div className="rounded-lg flex gap-4">
      <div className="max-w-7xl mx-auto p-6 bg-white shadow-md rounded-lg">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {/* Main Content */}
          <div className="col-span-2">
            <header className="flex flex-col mb-8">
              <h1 className="text-4xl font-bold text-gray-900 mb-2">{post?.title}</h1>
              <p className="text-gray-600 text-lg">
                <span>Published on: {post && new Date(post.created_at).toDateString()}</span>
              </p>
              <div className="flex justify-between items-center text-sm text-gray-500 mt-8">
                <div className="flex items-center">
                  <img
                    src="../avatar.png"
                    alt="Author"
                    className="w-10 h-10 rounded-full mr-4 object-cover"
                  />
                  <span>
                    By <strong className="text-gray-800">{post?.author}</strong>
                  </span>
                </div>
              </div>
            </header>

            <section className="mb-8">
              <div className="text-gray-700 leading-relaxed">
                <p>{post?.content}</p>
              </div>
            </section>
          </div>

          {/* Sidebar - Right Container */}
        </div>
      </div>{' '}
      <div className="md:col-span-1">
        <div className="bg-white p-6 rounded-lg shadow-md h-full">
          <h2 className="text-2xl font-bold text-gray-800 mb-4">Author Bio</h2>
          <p className="text-gray-700 mb-4">
            Author is a developer, writer, and open-source enthusiast. They write about technology,
            programming, and software development.
          </p>

          <div className="space-y-4">
            <div className="flex items-center space-x-2">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="w-5 h-5 text-blue-600"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M12 11c1.104 0 2-.896 2-2V4h-4v5c0 1.104.896 2 2 2zM5 8V6c0-1.104.896-2 2-2h10c1.104 0 2 .896 2 2v2c0 1.104-.896 2-2 2H7c-1.104 0-2-.896-2-2z"
                />
              </svg>
              <a
                href="https://twitter.com/author"
                target="_blank"
                rel="noopener noreferrer"
                className="text-blue-600 hover:underline"
              >
                @author
              </a>
            </div>

            <div className="flex items-center space-x-2">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="w-5 h-5 text-gray-800"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M15 19l-7-7 7-7"
                />
              </svg>
              <a href="/related-posts" className="text-gray-800 hover:underline">
                View Related Posts
              </a>
            </div>
            <button className="bg-blue-500 text-white rounded-[5px] w-full p-2">Follow</button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default PostDetails
