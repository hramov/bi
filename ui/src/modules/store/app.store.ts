import {defineStore} from "pinia";
import {computed} from "vue";
import {useRouter} from "vue-router";

export const useAppStore = defineStore('app', () => {
   const router = useRouter();

   const pagesMapper = {
      'dashboards': {
         title: 'Дашборды',
         path: '/dashboards',
      },
      'data_sources': {
         title: 'Источники данных',
         path: '/data_sources'
      }
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
      return Array.from(new Set(router.currentRoute.value.fullPath.split('/').filter(el => el != 'Главная').map(el => el === "" ? { title: 'Главная', path: '/'} : pagesMapper[el]))).filter(el => el);
   });

   return {
      calcBreadcrumbs,
      sortOptions
   }
});