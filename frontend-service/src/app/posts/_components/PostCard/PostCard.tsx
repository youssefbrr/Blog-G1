import React from 'react'
import { IPost } from '../PostsItems/PostsItems'
import PostActions from '@/components/PostActions/PostActions'
import Link from 'next/link'
interface PostCardProps {
  post: IPost
  actions?: {
    delete: () => void
    edit: (newPost: IPost) => void
    open: () => void
  }
}

const PostCard = ({ post, actions }: PostCardProps) => {
  const { title, content, author, createdAt } = post

  return (
    <div className="w-full bg-white border decoration-neutral-300 rounded-xl relative group">
      <div className="absolute right-2 top-2 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
        {actions &&
        <PostActions
        onDelete={actions.delete}
        update={() => actions.edit(post)}
        open={actions.open}
        post={post}
        />
      }
      </div>
      <div className="md:flex">
        <div className="p-8">
          <div className="flex items-center gap-3 uppercase tracking-wide text-sm text-indigo-500 font-semibold mb-2">
            <img
              className="w-8 h-8 rounded-full"
              src={`https://avatar.iran.liara.run/public/boy?username=${author}`}
              alt="Rounded avatar"
            />

            {author}
          </div>

          <Link
            href={`/posts/${post.id}`}
            className="block mt-1 text-lg leading-tight font-medium text-black hover:underline"
          >
            {title}
          </Link>
          <p className="mt-2 text-gray-500">{content}</p>
          <div className="mt-4 text-gray-500 text-sm">{createdAt}</div>
        </div>
      </div>
    </div>
  )
}

export default PostCard
