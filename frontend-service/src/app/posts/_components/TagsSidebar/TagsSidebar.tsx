import React from 'react'

const TagsSidebar = () => {
  const menuItems = [
    { icon: '🏠', label: 'Home' },
    { icon: '⚙️', label: 'DEV++' },
    { icon: '📚', label: 'Reading List' },
    { icon: '🎙️', label: 'Podcasts' },
    { icon: '🎥', label: 'Videos' },
    { icon: '🏷️', label: 'Tags' },
    { icon: '💡', label: 'DEV Help' },
    { icon: '🛍️', label: 'Forem Shop' },
    { icon: '❤️', label: 'Advertise on DEV' },
    { icon: '🏆', label: 'DEV Challenges' },
    { icon: '✨', label: 'DEV Showcase' },
    { icon: '🛡️', label: 'About' },
    { icon: '📞', label: 'Contact' },
    { icon: '📘', label: 'Guides' },
    { icon: '🤔', label: 'Software comparisons' },
  ]

  return (
    <div className="h-screen">
      <ul className="flex flex-col gap-3">
        {menuItems.map((item, index) => (
          <li
            key={index}
            className="flex items-center gap-3 hover:bg-[#e5e5f4] hover:text-[#2f3ab2] cursor-pointer rounded-lg p-1 pl-2"
          >
            <span className="text-xl">{item.icon}</span>
            <span className="font-lg">{item.label}</span>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default TagsSidebar
