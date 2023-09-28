import axios from 'axios'

const headers = {
  'Access-Control-Allow-Origin': '*',
  'Content-Type': 'application/json',
}

class GetDataError extends Error {
  constructor(message = '') {
    super(message)
    this.name = 'GetDataError'
  }
}

/**
 * Interceptors are used to analyze coming data before client does
 * Here is errors checking and errors handling
 */
axios.interceptors.response.use(
  function (response) {
    return response
  },
  function (error) {
    if (error.response)
      return {
        status: error.response.status,
        data: error.response.data,
      }
    return {
      status: error.status,
      data: error,
    }
  },
)

export default class ApiManager {
  static apiUri = import.meta.env.VITE_APP_API_URL
  static userUri = import.meta.env.VITE_APP_USER_URL

  static async get(endpoint: string) {
    const result = await axios.get(ApiManager.apiUri + endpoint, {
      headers,
    })
    return result.data
  }

  static async post<T>(endpoint: string, data: T) {
    return await axios.post(ApiManager.apiUri + endpoint, data)
  }

  static async put<T>(endpoint: string, data: T) {
    return await axios.put(ApiManager.apiUri + endpoint, data)
  }

  static async delete(id: number) {
    const result = await axios.delete(ApiManager.apiUri + id)
    return result.data
  }

  static async getUser() {
    const user = await axios.get(ApiManager.userUri)
    if (user.data) return user.data
    return new GetDataError()
  }
}
