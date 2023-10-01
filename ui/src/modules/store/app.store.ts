import {defineStore} from "pinia";
import {computed, ref} from "vue";
import {useRouter} from "vue-router";

export const useAppStore = defineStore('app', () => {
   const router = useRouter();

   const pagesMapper = {
      'dashboards': 'Дашборды',
      'dashboard': 'Дашборд',
      'data_sources': 'Источники данных'
   }

   const sortOptions = [
      {
         label: 'Сначала новые',
         value: 'desc',
      },
      {
         label: 'Сначала старые',
         value: 'asc',
      },
      {
         label: 'От А до Я',
         value: 'atoz',
      },
      {
         label: 'От Я до А',
         value: 'ztoa',
      },
   ]

   const calcBreadcrumbs = computed(() => {
      return Array.from(new Set(router.currentRoute.value.fullPath.split('/').filter(el => el != 1).map(el => el === "" ? "Главная" : pagesMapper[el]))).filter(el => el);
   });

   return {
      calcBreadcrumbs,
      sortOptions
   }
});