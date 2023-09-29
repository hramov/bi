import {optionsStr, replaceFunctions} from "./index.ts";

test('Replace Function', () => {
    const result = replaceFunctions(optionsStr);

    console.log(result)
});