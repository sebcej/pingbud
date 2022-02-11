// Parse stats for apex charts usage
export function parsedStats (state) {
    const stats = state.graph
    const chartSeries = [
        {name: 'Average', data: []},
        {name: 'Max', data: []},
        {name: 'Min', data: []},
        {name: 'Jitter', data: []}
    ]
    const chartCategories = []
    
    for (const stat of stats) {
        chartCategories.push(new Date(stat.time * 1000).toLocaleTimeString())

        chartSeries[0].data.push(stat.isOnline ? stat.avg : null)
        chartSeries[1].data.push(stat.isOnline ? stat.max : null)
        chartSeries[2].data.push(stat.isOnline ? stat.min : null)
        chartSeries[3].data.push(stat.isOnline ? stat.jitter : null)
    }

    return {
        series: chartSeries,
        categories: chartCategories
    }
}
