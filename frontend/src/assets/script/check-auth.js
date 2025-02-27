// import Cookies from "js-cookie";
// import axios from "axios";

// import { useRouter } from "vue-router";

// export function CheckAuth() {
//   const router = useRouter();
//   const BASE_URL_BACKEND = import.meta.env.VITE_BACKEND_BASE_URL;
//   const token = Cookies.get("AUTH_TOKEN");

//   let role = "";

//   if (!token) {
//     router.push("/login");
//   } else {
//     checkToken();
//   }

//   async function checkToken() {
//     try {
//       axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
//       const response = await axios.get(`${BASE_URL_BACKEND}/checkToken`);
//       role = response.data.data;
//     } catch (error) {
//       router.push("/login");
//     }
//   }

//   return role;
// }
