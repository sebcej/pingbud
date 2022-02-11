import axios from 'axios'

const api = axios.create({ baseURL: '/api/v1' })

export default ({ app }) => {
    app.config.globalProperties.$axios = axios
    app.config.globalProperties.$api = api
}

export { axios, api }