import {createRouter, createWebHashHistory} from "vue-router";
import AccountsView from "../views/AccountsView.vue";
import HomeView from "@/views/HomeView.vue";
import HelpView from "@/views/HelpView.vue";
import UserAgreementView from "@/views/UserAgreementView.vue";

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: "/",
            name: "accounts",
            component: AccountsView,
        },
        {
            path: "/about",
            name: "about",
            // route level code-splitting
            // this generates a separate chunk (About.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import("../views/AboutView.vue"),
        },
        {
            path: "/home",
            name: "home",
            component: HomeView,
        },
        {
            path: "/help",
            name: "help",
            component: HelpView,
        },
        {
            path: "/userAgreement",
            name: "userAgreement",
            component: UserAgreementView,
        },
    ],
});

export default router;
