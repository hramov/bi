import {defineStore} from "pinia";
import {ref} from "vue";
import ApiManager from "../api";

export const useDashboardStore = defineStore('dashboard', () => {

    const availableTypes = ref([]);
    const dashboard = ref({});
    const dashboards = ref([] as any);

    const getAvailableTypes = async () => {
        availableTypes.value = await ApiManager.get('/dashboards/types');
    }

    const getDashboard = async (id: number) => {
        // dashboard.value = await ApiManager.get('/dashboards/' + id);
        dashboard.value = dashboards.value[0];
    }

    const getItemById = async (id: number) => {
       return ApiManager.get('/dashboards/item/' + id);
    }

    const getDashboards = async () => {
        dashboards.value = [{
            id: 1,
            title: 'Показатели филиала',
            description: 'Показатели уровня ОАО РЖД',
            last_updated: new Date(),
        }]
    }

    const getAvailability = async (period: number) => {
        return [
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
    }

    const saveChart = async (options: any) => {
        const result = await ApiManager.post('/dashboards/item', options);
        if (result.data) {
            return result.data;
        }
        return null;
    }

    return {
        availableTypes,
        dashboards,
        dashboard,
        getAvailableTypes,
        getDashboard,
        getDashboards,
        getAvailability,
        saveChart,
        getItemById,
    }
});