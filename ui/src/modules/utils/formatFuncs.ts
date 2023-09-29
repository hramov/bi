import moment from "moment/moment";

export function percentFormatter(val: number) {
    if (Number(val)) {
        return (val * 100).toFixed(2)
    }
    return String(val);
}

export function percentLabelFormatter(val: number) {
    return val + '%'
}

export function twoDigitsFormatter(val: any) {
    if (Number(val)) {
        return val.toFixed(2)
    }
    return val
}

export function formatDuration(days: any, addSign?: any) {
    if (days === null) {
        return '--'
    }
    const isNegative = days < 0
    let mm = days * 1440
    mm *= isNegative ? -1 : 1
    let HH = Math.floor(mm / 60)
    mm = Math.round(mm - HH * 60)
    if (mm === 60) {
        mm = 0
        HH++
    }
    return (isNegative ? '-' : addSign ? '+' : '') + HH + ':' + (mm < 10 ? '0' + mm : mm)
}

export function timeFormatter(params: any) {
    let result = `${params && params[0] ? params[0].axisValueLabel : ''}<br/>`
    for (const data of params.sort((a: any, b: any) => (a.data > b.data ? -1 : 1))) {
        result += `${data.marker} ${data.seriesName} <strong>${formatDuration(data.data)}</strong><br/>`
    }
    return result
}

export function defaultTooltipFormatter(seriesTypes: string[]) {
    return function (params: any) {
        let result = `${params && params[0] ? params[0].axisValueLabel : ''}<br/>`
        for (const [index, data] of params.entries()) {
            if (seriesTypes[index] === 'time') {
                result += `${data.marker} ${data.seriesName} <strong>${formatDuration(data.data)}</strong><br/>`
            } else if (seriesTypes[index] === 'percent') {
                result += `${data.marker} ${data.seriesName} <strong>${(Number(data.data) * 100).toFixed(2)}%</strong><br/>`
            } else if (seriesTypes[index] === 'percent-sign') {
                result += `${data.marker} ${data.seriesName} <strong>${data.data ? data.data : 0}%</strong><br/>`
            } else {
                result += `${data.marker} ${data.seriesName} <strong>${Number(data.data).toLocaleString()}</strong><br/>`
            }
        }
        return result
    }
}

export function sortValuesFormatter(params: any) {
    let result = `${params && params[0] ? params[0].axisValueLabel : ''}<br/>`
    for (const data of params.sort((a: any, b: any) => (a.data > b.data ? -1 : 1))) {
        result += `${data.marker} ${data.seriesName} <strong>${data.data}</strong><br/>`
    }
    return result
}

export function firstCapitalLetter(str: string) {
    return str.charAt(0).toUpperCase()
}

export function calcMinValuePercent(chartData: any, field: string, options: any, axisId: number) {
    let min = Number.MAX_VALUE, max = Number.MIN_VALUE;
    for (const d of chartData.value) {
        min = Math.min(d[field], min)
        max = Math.max(d[field], max)
    }
    options.value.yAxis[axisId].min = ((min - ((max - min) * 0.1)) * 100).toFixed(2);
}

export function calcMinValue(chartData: any, field: string, options: any, axisId: number) {
    let min = Number.MAX_VALUE, max = Number.MIN_VALUE;
    for (const d of chartData.value) {
        min = Math.min(d[field], min)
        max = Math.max(d[field], max)
    }
    options.value.yAxis[axisId].min = (min - ((max - min) * 0.1)).toFixed(2);
}

export function calcMinValueTime(chartData: any, field: string, options: any, axisId: number) {
    let min = Number.MAX_VALUE, max = Number.MIN_VALUE;
    for (const d of chartData) {
        min = Math.min(d[field], min)
        max = Math.max(d[field], max)
    }
    options.value.yAxis[axisId].min = min - ((max - min) * 0.1);
}

export function defaultLabelFormatter (val: string) {
    return moment(val).format('DD.MM.YYYY')
}