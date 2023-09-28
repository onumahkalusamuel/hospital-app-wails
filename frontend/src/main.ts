import { createApp } from 'vue';
import './style.css';
import App from './App.vue';
import { router } from './router';

// create app and attach router
createApp(App).use(router).mount('#app');