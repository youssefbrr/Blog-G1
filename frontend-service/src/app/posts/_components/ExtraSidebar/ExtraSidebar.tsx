import React from 'react'

const users = [
  {
    name: 'Nir Zicherman',
    bio: 'CEO of Oboe (http://oboe.fyi). Former VP of Audiobooks at...',
    avatar: 'https://avatar.iran.liara.run/public/boy?username=Nir Zicherman',
  },
  {
    name: 'The Belladonna Comedy',
    bio: 'Comedy and satire by women and marginalized genders, fo...',
    avatar: 'https://avatar.iran.liara.run/public/boy?username=hazem',
  },
  {
    name: 'Ian Williams',
    bio: 'I write things about video games, labor, pop culture....',
    avatar: 'https://avatar.iran.liara.run/public/boy?username=Ian Williams',
  },
]

const topics = [
  'Programming',
  'Writing',
  'Self Improvement',
  'Data Science',
  'Relationships',
  'Technology',
  'Politics',
]

const ExtraSidebar = () => {
  return (
    <div className="flex flex-col gap-y-3">
      <div className="bg-white rounded-lg p-4">
        <h2 className="text-lg font-semibold mb-4">Recommended topics</h2>
        <div className="flex flex-wrap gap-3">
          {topics.map((topic, index) => (
            <span
              key={index}
              className="px-4 py-2 bg-gray-100 text-gray-700 text-sm rounded-full hover:bg-gray-200 cursor-pointer"
            >
              {topic}
            </span>
          ))}
        </div>
      </div>
      <div className="bg-white rounded-lg p-4">
        <h2 className="text-lg font-semibold mb-4">Who to follow</h2>
        <div className="space-y-4">
          {users.map((user, index) => (
            <div key={index} className="flex items-center justify-between">
              <div className="flex items-center space-x-3">
                <img
                  src={user.avatar}
                  alt={`${user.name}'s avatar`}
                  className="w-10 h-10 min-w-10 rounded-full"
                />
                <div>
                  <h3 className="text-sm font-semibold">{user.name}</h3>
                  <p className="text-xs text-gray-500">{user.bio}</p>
                </div>
              </div>
              <button className="px-4 py-1 text-sm border border-gray-300 rounded-full hover:bg-gray-100">
                Follow
              </button>
            </div>
          ))}
        </div>
        <a href="#" className="block mt-4 text-sm text-indigo-500 hover:underline">
          See more suggestions
        </a>
      </div>
    </div>
  )
}

export default ExtraSidebar
