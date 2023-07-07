import { createRouter, createWebHistory } from 'vue-router';
import HomePage from './components/HomePage.vue'
import UploadPage from './components/UploadPage.vue';
import DownloadPage from './components/DownloadPage.vue';

const routes = [
    {
        path:'/',
        name:'Home',
        component: HomePage
    },
    {
        path:'/upload',
        name:'Upload',
        component: UploadPage
    },
    {
        path:'/download',
        name:'Download',
        component: DownloadPage
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router;