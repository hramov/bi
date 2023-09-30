import {
    defaultLabelFormatter,
    defaultTooltipFormatter,
    percentFormatter,
    percentLabelFormatter
} from "./formatFuncs";

export const funcMap = {
    'fn_percentFormatter': percentFormatter,
    'fn_percentLabelFormatter': percentLabelFormatter,
    'fn_defaultLabelFormatter': defaultLabelFormatter,
    'fn_defaultTooltipFormatter': defaultTooltipFormatter,
} as any;

export function replaceFunctions (options: string) {
    const optObject = JSON.parse(options);
    replaceFunctionsImpl(optObject);
    return optObject;
}

function replaceFunctionHandler(root: any, key: string, parent: any, parentKey: string) {
    if (typeof root[key] === 'string' && root[key].startsWith('fn_') && key !== 'fn') {
        const handler = root[key] as string;

        if (typeof funcMap[handler] === 'string') {
            root[key] = new Function('return ' + funcMap[handler].replace('\n', ''))();
        } else {
            root[key] = funcMap[handler];
        }
    } else if (typeof root[key] === 'string' && root[key].startsWith('fn_') && key === 'fn') {
        const handler = root[key];
        if (typeof funcMap[handler] === 'string') {
            root[key] = new Function('return ' + funcMap[handler].replace('\n', ''))().apply(null, root['params']);
        } else {
            parent[parentKey] = funcMap[handler].call(null, root['params']);
        }
    } else if (typeof root[key] === 'object') {
        replaceFunctionsImpl(root[key], root, key)
    }
}

export function replaceFunctionsImpl(root: any, parent = null, parentKey = '') {
    if (typeof root === 'function') {
        throw new Error('Functions are not allowed');
    }

    if (typeof root === 'object') {
        if (Array.isArray(root)) {
            for (let r of root) {
                if (typeof r === 'object') {
                    for (let key in r) {
                        replaceFunctionHandler(r, key, parent, parentKey!)
                    }
                }
            }
        } else {
            for (let key in root) {
                replaceFunctionHandler(root, key, parent, parentKey!)
            }
        }
    }
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

export const optionsStr = `{
    "id": "availability",
    "grid": {
        "containLabel": true,
        "x": "7%",
        "y": "7%",
        "x2": "5%",
        "y2": "7%"
    },
    "labels": [
        {
            "type": "line",
            "title": "Значение",
            "display": "Доступность",
            "formatter": "fn_percentFormatter"
        },
        {
            "yAxisID": 1,
            "type": "bar",
            "title": "Всего",
            "display": "Всего",
            "color": "#c0bdbd",
            "opacity": 0.5
        }
    ],
    "labelFormatter": "fn_defaultLabelFormatter",
    "yAxis": [
        {
            "type": "value",
            "position": "left",
            "name": "Доступность",
            "nameLocation": "middle",
            "nameGap": 70,
            "axisLabel": {
                "formatter": "fn_percentLabelFormatter"
            },
            "min": "dataMin"
        },
        {
            "type": "value",
            "position": "right",
            "splitLine": {
                "show": false
            },
            "axisLabel": {
                "show": false
            }
        }
    ],
    "xAxis": {
        "field": "Дата",
        "type": "category"
    },
    "tooltip": {
        "trigger": "axis",
        "cross": true,
        "axisPointer": {
            "type": "shadow"
        },
        "formatter": {
            "fn": "fn_defaultTooltipFormatter",
            "params": [
                "default",
                "percent-sign"
            ]
        }
    },
    "legend": {}
}`;