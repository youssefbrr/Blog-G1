'use client'
import { useEffect } from 'react'
import { useRouter } from 'next/navigation'

const Search = () => {
  const router = useRouter()

  useEffect(() => {
    router.push('/posts')
  }, [router])

  return null
}

export default Search
