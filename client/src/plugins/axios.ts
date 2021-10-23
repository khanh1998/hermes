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

export default {
  install(app: any, options: any) {
    app.config.globalProperties.$http = axiosInstance
  }
}