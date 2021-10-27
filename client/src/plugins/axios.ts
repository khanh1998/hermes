// User Definition
import axios from 'axios'

let apiUri;
let authUri;
if (import.meta.env.DEV) {
  apiUri = import.meta.env.VITE_API_HOST as string;
  authUri = import.meta.env.VITE_AUTH_SERVICE as string;
} else {
  // config for production
  apiUri = `${import.meta.env.VITE_HOST}${import.meta.env.VITE_API_PATH}`;
}

const axiosConfigApi = {
  baseURL: apiUri,
  timeout: 30000,
};
const axiosInstance = axios.create(axiosConfigApi);

const axiosConfigAuth = {
  baseURL: authUri,
  timeout: 30000,
};
const axiosInstanceAuth = axios.create(axiosConfigAuth);
const interceptors = async (config: any) => {
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
}
axiosInstance.interceptors.request.use(interceptors,
  (error) => {
    // Do something with request error
    Promise.reject(error);
  },
);
axiosInstanceAuth.interceptors.request.use(interceptors,
  (error) => {
    // Do something with request error
    Promise.reject(error);
  },
);

export default {
  install(app: any, options: any) {
    app.config.globalProperties.$http = axiosInstance
    window.axios = axiosInstance;
    window.axiosAuth = axiosInstanceAuth;
  }
}