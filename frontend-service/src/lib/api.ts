import { IPost } from '@/app/posts/_components/PostsItems/PostsItems'
import { axiosInstance, searchServiceAxiosInstance } from './axios'

export const getPostsApi = async (params: any = {}) => {
  const response = await axiosInstance.get('/api/v1/posts', { params })
  return response.data
}

export const getPostApi = async (id: string) => {
  const response = await axiosInstance.get(`/api/v1/posts/${id}`)
  return response.data
}

export const createPostApi = async (post: Omit<IPost, 'id'>) => {
  const response = await axiosInstance.post('/api/v1/posts', post)
  return response.data
}

export const updatePostApi = async (id: number, post: Omit<IPost, 'id'>) => {
  const response = await axiosInstance.put(`/api/v1/posts/${id}`, post)
  return response.data
}

export const deletePostApi = async (id: number) => {
  await axiosInstance.delete(`/api/v1/posts/${id}`)
}

export const getSearchPostsApi = async (params: any = {}) => {
  const response = await searchServiceAxiosInstance.get('/search', { params })
  return response.data
}
