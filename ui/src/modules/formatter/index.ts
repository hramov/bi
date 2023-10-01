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
    if (!options) return {};
    replaceFunctionsImpl(options);
    return options;
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