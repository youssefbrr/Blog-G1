import { IPost } from '@/app/posts/_components/PostsItems/PostsItems'
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/react'
import React from 'react'

interface PostActionsProps {
  onDelete: () => void
  open: () => void
  update: () => void
  post: IPost
}

const PostActions = ({ onDelete, open, update, post }: PostActionsProps) => {
  return (
    <Menu as="div" className="relative inline-block text-left">
      <div>
        <MenuButton className="inline-flex w-full justify-center gap-x-1.5 rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50">
          ...
        </MenuButton>
      </div>

      <MenuItems
        transition
        className="absolute right-0 z-10 mt-2 w-56 origin-top-right divide-y divide-gray-100 rounded-md bg-white shadow-lg ring-1 ring-black/5 transition focus:outline-none data-[closed]:scale-95 data-[closed]:transform data-[closed]:opacity-0 data-[enter]:duration-100 data-[leave]:duration-75 data-[enter]:ease-out data-[leave]:ease-in"
      >
        <div className="py-1">
          <MenuItem>
            <button
              onClick={open}
              className="block w-full text-left px-4 py-2 text-sm text-gray-700 data-[focus]:bg-gray-100 data-[focus]:text-gray-900 data-[focus]:outline-none"
            >
              Open
            </button>
          </MenuItem>
          <MenuItem>
            <button
              onClick={update}
              className="block w-full text-left px-4 py-2 text-sm text-gray-700 data-[focus]:bg-gray-100 data-[focus]:text-gray-900 data-[focus]:outline-none"
            >
              Edit
            </button>
          </MenuItem>
        </div>
        <div className="py-1">
          <MenuItem>
            <button
              onClick={open}
              className="block w-full text-left px-4 py-2 text-sm text-gray-700 data-[focus]:bg-gray-100 data-[focus]:text-gray-900 data-[focus]:outline-none"
            >
              Add to favorites
            </button>
          </MenuItem>
        </div>
        <div className="py-1">
          <MenuItem>
            <button
              onClick={onDelete}
              className="block w-full text-left px-4 py-2 text-sm text-gray-700 data-[focus]:bg-gray-100 data-[focus]:text-gray-900 data-[focus]:outline-none"
            >
              Delete
            </button>
          </MenuItem>
        </div>
      </MenuItems>
    </Menu>
  )
}

export default PostActions
