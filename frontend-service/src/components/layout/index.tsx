'use client'
import type { ReactNode } from 'react'
import { useRouter } from 'next/navigation'
import { useCallback } from 'react'
import debounce from 'lodash/debounce'
export function Layout({ children }: { children: ReactNode }) {
  const router = useRouter()

  const onInputChange = useCallback(
    debounce((e: any) => {
      const value = e.target.value
      if (value) {
        router.push(`/search/${value}`)
      } else {
        router.push('/posts')
      }
    }, 300),
    [],
  )

  return (
    <div className="flex flex-col mx-auto min-h-screen bg-neutral-100">
      <header className="h-14 border-b-[1px] decoration-neutral-300 bg-background sticky top-0 left-0 w-full z-10">
        <div className="flex justify-between items-center max-w-[1300px] m-auto h-full p-3">
          <div className="h-min">
            <input
              type="text"
              onChange={onInputChange}
              id="default-input"
              placeholder="What is on your mind?"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 min-w-[400px]"
            />
          </div>
          <div>
            <button className="bg-white hover:bg-gray-100  font-semibold py-2 px-4 border text-indigo-500 border-indigo-500 rounded-lg shadow">
              Create Post
            </button>
          </div>
        </div>
      </header>
      <div className="max-w-[1300px] w-[1300px] mx-auto p-3 flex-1">{children}</div>
    </div>
  )
}
