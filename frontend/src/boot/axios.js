import axios from 'axios'
import { Notify } from 'quasar'

const api = axios.create({ baseURL: '/api/v1' })

api.interceptors.response.use((response) => response, (error) => {
    const rstatus = error.response.status

    const res = {
        color: 'negative',
        message: 'Network error'
    }

    if (rstatus >= 500 && rstatus < 600) {
        res.message = 'Server error'
    } else if (rstatus >= 400 && rstatus < 500) {
        res.message = 'Invalid data'
    }

    Notify.create(res)

    throw error;
});

export default ({ app }) => {
    app.config.globalProperties.$axios = axios
    app.config.globalProperties.$api = api
}

export { axios, api }