import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import Login from './views/LoginPage.vue'

// creating page router
const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', name: 'Login', component: Login },
        //        {path: '/home', name: 'Home', component: Home}
    ]
})

// creating app instance and mounting it to the dom
createApp(App).use(router).mount('#app')
