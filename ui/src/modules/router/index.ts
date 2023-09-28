import {createRouter, createWebHashHistory} from "vue-router";
import Main from "../../views/main/Main.vue";
import Dashboards from "../../views/dashboards/Dashboards.vue";
import Settings from "../../views/settings/Settings.vue";
import Help from "../../views/help/Help.vue";
import SingleDashboard from "../../views/single_dashboard/SingleDashboard.vue";

const routes = [
    { path: '/', component: Main },
    { path: '/dashboards', component: Dashboards },
    { path: '/dashboards/:dashboard_id', component: SingleDashboard },
    { path: '/data_sources', component: Dashboards },
    { path: '/settings', component: Settings },
    { path: '/help', component: Help },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router;