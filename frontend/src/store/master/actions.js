import { api } from 'boot/axios'

export async function getSettings (context) {
    const results = await api.get('/settings')

    context.commit('setSettings', results.data.settings)
}

export async function getStats (context) {
    const results = await api({
        method: 'GET',
        url: '/stats',
        params: {
            'filter': context.state.dateFilter
        }
    })

    context.commit('setGraph', results.data.aggregated)
    context.commit('setStats', results.data.stats)
    context.commit('setLatest', results.data.latest)
    context.commit('setErrors', results.data.errors)
}

export async function saveSettings (context, settingsObj) {
    await api.post('/settings', settingsObj)
}

export async function toggleEnabled (context) {
    const settings = {...context.state.settings}
    settings.enabled = !settings.enabled

    await api.post('/settings', settings)
}