// import Cookies from "js-cookie";
// import axios from "axios";

// import { useRouter } from "vue-router";

// export function CheckAuth() {
//   const router = useRouter();
//   const BACKEND_BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
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
//       const response = await axios.get(`${BACKEND_BASE_URL}/checkToken`);
//       role = response.data.data;
//     } catch (error) {
//       router.push("/login");
//     }
//   }

//   return role;
// }
