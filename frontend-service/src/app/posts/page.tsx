import React from 'react'
import PostsItems from './_components/PostsItems/PostsItems'
import TagsSidebar from './_components/TagsSidebar/TagsSidebar'
import ExtraSidebar from './_components/ExtraSidebar/ExtraSidebar'

const Page = () => {
  return (
    <div className="flex justify-between gap-3">
      <div className="w-[240px]  rounded-md">
        <TagsSidebar />
      </div>
      <div className="flex-1">
        <PostsItems />
      </div>
      <div className="w-[300px] rounded-md">
        <ExtraSidebar />
      </div>
    </div>
  )
}

export default Page
