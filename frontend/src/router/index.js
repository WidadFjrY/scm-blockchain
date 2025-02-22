import Login from "@/pages/LoginPage.vue";
import Register from "../pages/RegisterPage.vue";
import Dashboard from "@/pages/DashboardPage.vue";
import Order from "@/pages/OrderPage.vue";
import Distributor from "@/pages/DistributorPage.vue";
import Product from "@/pages/ProductPage.vue";
import Report from "@/pages/ReportPage.vue";
import Store from "@/pages/StorePage.vue";
import Users from "@/pages/UsersPage.vue";
import User from "@/pages/UserPage.vue";
import FormProduct from "@/components/FormProduct.vue";

import { createRouter, createWebHistory } from "vue-router";

const routes = [
  { path: "/login", name: "Login", component: Login },
  { path: "/register", name: "Register", component: Register },
  { path: "/", name: "Dashboard", component: Dashboard },
  { path: "/orders", name: "Manajemen Pesanan", component: Order },
  {
    path: "/distributors",
    name: "Manajemen Distributor",
    component: Distributor,
  },

  { path: "/store", name: "Manajemen Toko", component: Store },

  { path: "/products", name: "Produk", component: Product },
  { path: "/add/product", name: "Tambah Produk", component: FormProduct },

  { path: "/reports", name: "Laporan", component: Report },
  { path: "/users", name: "Pengguna", component: Users },
  { path: "/profile", name: "Profil", component: User },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
