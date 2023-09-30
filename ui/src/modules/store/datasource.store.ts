import {defineStore} from "pinia";
import {ref} from "vue";
import ApiManager from "../api";

export const useDatasourceStore = defineStore('datasource', () => {

    const drivers = ref([]);
    const driver = ref({});

    const sources = ref([]);
    const source = ref({});

    const getDrivers = async () => {
        drivers.value = await ApiManager.get('/datasource/driver');
    }

    const getDriverById = async (id: number) => {
        driver.value = await ApiManager.get('/datasource/driver/' + id);
    }

    const getSources = async () => {
        sources.value = await ApiManager.get('/datasource');
    }

    const getSourceById = async (id: number) => {
        source.value = await ApiManager.get('/datasource/' + id);
    }

    const saveSource = async (data: any) => {
        let result;
        if (data.id) {
            result = await ApiManager.put('/datasource/' + data.id, data);
        } else {
            result = await ApiManager.post('/datasource', data);
        }
        return result;
    }

    const deleteSource = async (id: number) => {
        const result = await ApiManager.delete('/datasource/', id);
        return result;
    }

    return {
        getDrivers,
        getDriverById,
        getSources,
        getSourceById,
        saveSource,
        deleteSource,
        drivers,
        driver,
        sources,
        source,
    }
});