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

    const getDashboard = async (id: string) => {
        dashboard.value = await ApiManager.get('/dashboards/' + id);
    }

    const getItemById = async (id: number) => {
       dashboard.value = await ApiManager.get('/dashboards/item/' + id);
    }

    const getDashboards = async () => {
        dashboards.value = (await ApiManager.get('/dashboards')).data;
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
        saveChart,
        getItemById,
    }
});