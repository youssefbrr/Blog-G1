import axios from 'axios'

const baseURL = 'https://blog1-posts.tamtasks.com'
const searchServiceURL = 'https://blog1-search-posts.tamtasks.com'

const headers = {
  Accept: 'application/json',
  'Content-Type': 'application/json',
}

export const axiosInstance = axios.create({
  baseURL,
  responseType: 'json',
  headers,
})

export const searchServiceAxiosInstance = axios.create({
  baseURL: searchServiceURL,
  responseType: 'json',
  headers,
})
