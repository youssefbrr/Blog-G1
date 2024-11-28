'use client'

import { useEffect, useState } from 'react'

const Posts = () => {
  const [posts, setPosts] = useState([])

  useEffect(() => {
    fetch('https://jsonplaceholder.typicode.com/posts')
      .then((response) => response.json())
      .then((data) => setPosts(data))
  }, [])

  console.log({ posts })

  return <div className="posts">Posts</div>
}

export default Posts
