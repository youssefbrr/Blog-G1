import axios from 'axios'

const baseURL = 'https://blog1-posts.tamtasks.com/api/v1'

const headers = {
  Accept: 'application/json',
  'Content-Type': 'application/json',
}

const axiosInstance = axios.create({
  baseURL,
  responseType: 'json',
  headers,
})

export default axiosInstance
