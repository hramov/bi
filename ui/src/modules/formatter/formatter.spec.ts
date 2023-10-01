import {replaceFunctions} from ".";

const optionsStr = `{
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

test('Replace Function', () => {
    const result = replaceFunctions(optionsStr);

    console.log(result)
});