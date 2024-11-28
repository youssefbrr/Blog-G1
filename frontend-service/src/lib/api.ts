import { IPost } from '@/app/posts/_components/PostsItems/PostsItems'
import axiosInstance from './axios'

export const getPostsApi = async (params: any = {}) => {
  const response = await axiosInstance.get('/posts', { params })
  return response.data
}

export const getPostApi = async (id: string) => {
  const response = await axiosInstance.get(`/posts/${id}`)
  return response.data
}

export const createPostApi = async (post: Omit<IPost, 'id'>) => {
  const response = await axiosInstance.post('/posts', post)
  return response.data
}

export const updatePostApi = async (id: number, post: Omit<IPost, 'id'>) => {
  const response = await axiosInstance.put(`/posts/${id}`, post)
  return response.data
}

export const deletePostApi = async (id: number) => {
  await axiosInstance.delete(`/posts/${id}`)
}
