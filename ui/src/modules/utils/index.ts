import moment from "moment/moment";

export function twoDigitsFormatter(val: any) {
    if (Number(val)) {
        return val.toFixed(2)
    }
    return val
}

export function percentFormatter(val: any) {
    if (Number(val)) {
        return (val * 100).toFixed(2)
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

export function calcMinValueTime(data: any, field: string, options: any, axisId: number) {
    let min = Number.MAX_VALUE, max = Number.MIN_VALUE;
    for (const d of data) {
        min = Math.min(d[field], min)
        max = Math.max(d[field], max)
    }
    options.value.yAxis[axisId].min = min - ((max - min) * 0.1);
}

export const data = [
    {
        "Дата": "2023-06-26T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 177445,
        "Значение": 0.9925441686156273
    },
    {
        "Дата": "2023-07-03T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 161094,
        "Значение": 0.9932461792493823
    },
    {
        "Дата": "2023-07-10T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 160937,
        "Значение": 0.9936745434548923
    },
    {
        "Дата": "2023-07-17T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 148097,
        "Значение": 0.9937675982633004
    },
    {
        "Дата": "2023-07-24T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 143132,
        "Значение": 0.9941801973003941
    },
    {
        "Дата": "2023-07-31T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 153210,
        "Значение": 0.9945695450688598
    },
    {
        "Дата": "2023-08-07T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 175685,
        "Значение": 0.9951276432250904
    },
    {
        "Дата": "2023-08-14T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 178647,
        "Значение": 0.9944471499661344
    },
    {
        "Дата": "2023-08-21T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 176909,
        "Значение": 0.9945847865286673
    },
    {
        "Дата": "2023-08-28T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 165068,
        "Значение": 0.9943053771778904
    },
    {
        "Дата": "2023-09-04T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 208920,
        "Значение": 0.9955389622822133
    },
    {
        "Дата": "2023-09-11T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 183107,
        "Значение": 0.9942383415161626
    },
    {
        "Дата": "2023-09-18T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 211798,
        "Значение": 0.9956231881320882
    },
    {
        "Дата": "2023-09-25T00:00:00.000Z",
        "Показатель_ID": 60,
        "Всего": 111374,
        "Значение": 0.9950886203243127
    }
]

export const options = {
    id: 'availability',
    grid: { containLabel: true, x: '7%', y: '7%', x2: '5%', y2: '7%' },
    labels: [
        {
            type: 'line',
            title: 'Значение',
            display: 'Доступность',
            formatter: (val: number) => {
                if (Number(val)) {
                    return (val * 100).toFixed(2)
                }
                return val
            },
        },
        {
            yAxisID: 1,
            type: 'bar',
            title: 'Всего',
            display: 'Всего',
            color: '#c0bdbd',
            opacity: 0.5,
        },
    ],
    labelFormatter: (val: string) => {
        return moment(val).format('DD.MM.YYYY')
    },
    yAxis: [
        {
            type: 'value',
            position: 'left',
            name: 'Доступность',
            nameLocation: 'middle',
            nameGap: 70,
            axisLabel: {
                formatter: (val: number) => {
                    return val + '%'
                },
            },
            min: 'dataMin',
        },
        {
            type: 'value',
            position: 'right',
            splitLine: {
                show: false,
            },
            axisLabel: {
                show: false,
            },
        },
    ],
    xAxis: {
        field: 'Дата',
        type: 'category',
    },
    tooltip: {
        trigger: 'axis',
        cross: true,
        axisPointer: {
            type: 'shadow',
        },
        formatter: defaultTooltipFormatter(['default', 'percent-sign']),
    },
    legend: {},
}

export const options2 = {
    id: 'availability2',
    grid: { containLabel: true, x: '7%', y: '7%', x2: '5%', y2: '7%' },
    labels: [
        {
            type: 'line',
            title: 'Значение',
            display: 'Доступность',
            formatter: (val: number) => {
                if (Number(val)) {
                    return (val * 100).toFixed(2)
                }
                return val
            },
        },
        {
            yAxisID: 1,
            type: 'bar',
            title: 'Всего',
            display: 'Всего',
            color: '#c0bdbd',
            opacity: 0.5,
        },
    ],
    labelFormatter: (val: string) => {
        return moment(val).format('DD.MM.YYYY')
    },
    yAxis: [
        {
            type: 'value',
            position: 'left',
            name: 'Доступность',
            nameLocation: 'middle',
            nameGap: 70,
            axisLabel: {
                formatter: (val: number) => {
                    return val + '%'
                },
            },
            min: 'dataMin',
        },
        {
            type: 'value',
            position: 'right',
            splitLine: {
                show: false,
            },
            axisLabel: {
                show: false,
            },
        },
    ],
    xAxis: {
        field: 'Дата',
        type: 'category',
    },
    tooltip: {
        trigger: 'axis',
        cross: true,
        axisPointer: {
            type: 'shadow',
        },
        formatter: defaultTooltipFormatter(['default', 'percent-sign']),
    },
    legend: {},
}