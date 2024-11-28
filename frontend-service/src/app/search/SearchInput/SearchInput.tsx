
import { SetStateAction, Dispatch } from 'react'

interface ISeachInput {
  setQuery: Dispatch<SetStateAction<string>>
  query: string
}
export default function SearchInput({ setQuery, query }: ISeachInput) {
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setQuery(e.target.value)
  }

  return (
    <form className="relative w-full max-w-lg mx-auto">
      <input
        type="text"
        value={query}
        onChange={handleInputChange}
        placeholder="Search..."
        className="w-full px-4 py-2 text-gray-800 bg-gray-100 border border-gray-300 rounded-full focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
    </form>
  )
}
