import Vue from 'vue'
import Router from 'vue-router'
import LoginForm from '@/components/LoginForm'

Vue.use(Router)

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        {
            path: '/login',
            name: 'Login',
            component: LoginForm
        },
        // ...其他路由...
    ]
})
