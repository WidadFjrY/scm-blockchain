import Login from "@/pages/LoginPage.vue";
import Register from "@/pages/RegisterPage.vue";
import Home from "@/pages/HomePage.vue";
import Dashboard from "@/pages/DashboardPage.vue";
import Product from "@/pages/ProductPage.vue";
import Order from "@/pages/OrderPage.vue";
import Report from "@/pages/ReportPage.vue";
import User from "@/pages/UserPage.vue";

import { createRouter, createWebHistory } from "vue-router";

import Cookies from "js-cookie";
import axios from "axios";

const routes = [
  { path: "/login", name: "Login", component: Login },
  { path: "/register", name: "Register", component: Register },
  { path: "/", name: "Home", component: Home },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: Dashboard,
    meta: { requiresAdmin: true },
  },
  {
    path: "/products",
    name: "Produk",
    component: Product,
    meta: { requiresAdmin: true },
  },
  {
    path: "/orders",
    name: "Manajemen Pesanan",
    component: Order,
    meta: { requiresAdmin: true },
  },
  {
    path: "/users",
    name: "Pengguna",
    component: User,
    meta: { requiresAdmin: true },
  },
  {
    path: "/reports",
    name: "Laporan",
    component: Report,
    meta: { requiresAdmin: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, from, next) => {
  const token = Cookies.get("AUTH_TOKEN");
  const auth = await isAuthenticated(token);

  if (to.name !== "Login" && !auth.status) {
    return next({ name: "Login" });
  }

  if (auth.status && (to.name === "Login" || to.name === "Register")) {
    return next(from.name ? { name: from.name } : { name: "Home" });
  }

  if (to.meta.requiresAdmin && auth.role !== "Admin") {
    return next({ name: "Home" });
  }

  next();
});

async function isAuthenticated(token) {
  const BASE_URL_BACKEND = import.meta.env.VITE_BACKEND_BASE_URL;
  try {
    axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
    const response = await axios.get(`${BASE_URL_BACKEND}/checkToken`);

    return { status: response.status === 200, role: response.data.data };
  } catch (error) {
    console.log(error);
    return { status: false, role: "" };
  }
}

export default router;
