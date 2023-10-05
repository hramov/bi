import { createApp } from 'vue'
import App from './App.vue'
import { createPinia } from 'pinia'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { VSkeletonLoader } from 'vuetify/labs/VSkeletonLoader'

import router from "./modules/router";
import '@mdi/font/css/materialdesignicons.css'

const vuetify = createVuetify({
    components: {
        ...components,
        VSkeletonLoader
    },
    directives,
    theme: {
        defaultTheme: 'dark',
    }
});

const app = createApp(App);
app.use(vuetify)
app.use(router)
app.use(createPinia())
app.mount('#app')