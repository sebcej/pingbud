import { api } from 'boot/axios'

export async function getSettings (context) {
    const results = await api.get('/settings')

    context.commit('setSettings', results.data.settings)
}

export async function getStats (context) {
    const results = await api.get('/stats')

    context.commit('setGraph', results.data.aggregated)
    context.commit('setStats', results.data.stats)
    context.commit('setLatest', results.data.latest)
}

export async function saveSettings (context, settingsObj) {
    await api.post('/settings', settingsObj)
}