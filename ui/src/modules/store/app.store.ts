import {defineStore} from "pinia";
import {computed, ref} from "vue";
import {useRouter} from "vue-router";

export const useAppStore = defineStore('app', () => {
   const router = useRouter();

   const pagesMapper = {
      'dashboards': 'Дашборды',
      'dashboard': 'Дашборд',
   }

   const calcBreadcrumbs = computed(() => {
      return Array.from(new Set(router.currentRoute.value.fullPath.split('/').filter(el => el != 1).map(el => el === "" ? "Главная" : pagesMapper[el])));
   });

   return {
      calcBreadcrumbs,
   }
});