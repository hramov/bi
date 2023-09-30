import {CHART_TYPE, CHART_TYPES} from "../config/config";
import {data} from "../formatter";

export function chartMapper(options: any): any {
    switch(CHART_TYPE) {
        case CHART_TYPES.ECHARTS:
            return eChartsMapper(options);
        default:
            return '';
    }
}

/**
 *   title: '',
 *   description: '',
 *   dashboard: null,
 *   options: {
 *     x: {
 *       title: '',
 *       fn: null,
 *     },
 *     y: [] as any[]
 *   },
 */
function eChartsMapper(options: any): any {
    const dataQueries = {};

    const result = {
        id: options.id,
        grid: {
            containLabel: true,
            x: "7%",
            y: "7%",
            x2: "5%",
            y2: "7%"
        },
        labels: [] as any,
        yAxis: [] as any,
        labelFormatter: options.x.fn,
        xAxis: {
            field: options.y[0].xField,
            type: "category"
        },
        tooltip: {
            trigger: "axis",
            cross: true,
            axisPointer: {
                type: "shadow"
            },
            formatter: {}
        },
        legend: {}
    }

    for (const y of options.y) {
        result.yAxis.push({
            type: "value",
            position: "left",
            yAxisID: y.id,
            name: y.title,
            nameLocation: "middle",
            nameGap: 70,
            axisLabel: {
                formatter: y.fn,
            },
            min: "dataMin"
        });

        result.labels.push({
            type: y.type,
            title: y.yAxis,
            display: y.title,
            formatter: y.fn
        });

        dataQueries[y.id] = {
            source: y.source,
            query: y.query,
        }
    }
    return {
        rawOptions: JSON.stringify(result),
        dataQueries: dataQueries,
    }
}