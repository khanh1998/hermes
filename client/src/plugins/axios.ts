// User Definition
import axios from 'axios'

let apiUri;
if (import.meta.env.DEV) {
  apiUri = import.meta.env.VITE_API_HOST as string;
} else {
  apiUri = `${import.meta.env.VITE_HOST}${import.meta.env.VITE_API_PATH}`;
}
const axiosConfig = {
  baseURL: apiUri,
  timeout: 30000,
};

console.log(apiUri);

const axiosInstance = axios.create(axiosConfig);
axiosInstance.interceptors.request.use(
  async (config) => {
    const myConfig = { ...config };
    // Do something before request is sent
    // If request is different than any of the URLS in urlsExcludedForBearerHeader
    // then send Authorization header with token from localstorage
    const urlsExcludedForBearerHeader = [
      '/authentication/main',
      '/register',
    ];
    const notInExcludedList = urlsExcludedForBearerHeader.every((url) => myConfig.url && !myConfig.url.includes(url))
    if (notInExcludedList && myConfig.headers) {
      myConfig.headers.Authorization = `Bearer ${localStorage.getItem('main_token')}`;
    }
    return myConfig;
  },
  (error) => {
    // Do something with request error
    Promise.reject(error);
  },
);

export default {
  install(app: any, options: any) {
    app.config.globalProperties.$http = axiosInstance
    window.axios = axiosInstance;
  }
}