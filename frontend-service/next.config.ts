import type { NextConfig } from 'next'

const nextConfig: NextConfig = {
  /* config options here */
  async rewrites() {
    return [
      {
        source: '/search',
        destination: 'https://blog1-search-posts.tamtasks.com',
      },
    ]
  },
}

export default nextConfig
