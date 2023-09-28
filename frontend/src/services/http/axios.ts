import axios, { AxiosError, AxiosResponse } from 'axios';
import { router } from '@/router';
import { auth } from '@/stores/auth';
import { toasts } from '@/stores/toasts';

// axios instance
let __api = axios.create({ baseURL: `http://${window.location.hostname}:8788/api/`, proxy: false });
if (!window.location.hostname.match(/localhost|192.*|127.*|0.*/)) {
  __api = axios.create({ baseURL: `https://${window.location.hostname}/api/`, proxy: false });
}

// interceptor to add bearer token
__api.interceptors.request.use((request) => {
  if (auth.getJwt()) request.headers.Authorization = `Bearer ${auth.getJwt()}`;
  return request;
}, error => error);

// interceptor to redirect to login or other pages
__api.interceptors.response.use(async (response) => response, (error) => errorHandler(error));

interface ActivationObject {
  code?: number;
  message?: string;
  installation_code?: string;
}

function errorHandler (error: AxiosError) {
  if (error.response && error.response.status === 401) {
    const activationCheck = error.response.data as ActivationObject;

    if (activationCheck.code === 1) {
      return router.push({ name: 'activate', params: { code: activationCheck.installation_code } });
    }
    if (activationCheck.code === 2) {
      return router.push({ name: 'create-admin' });
    }
    if (activationCheck.code === 3) {
      return router.push({ name: 'hospital-setup' });
    }
    return router.push({ name: 'login' });
  }
  throw error;
}

const __multi = __api;
__multi.defaults.headers.common['Content-Type'] = 'multipart/form-data';

// for errors
const handleError = function(e: AxiosError) {
  let message;
  if (e.response && e.response.data) message = (e.response.data as { message: string; }).message;
  else message = e.message;

  toasts.addToast({message, type: 'error', title: 'Error'})
  return;
};

// for downloads
const handleDownload = function(response: AxiosResponse) {
  if (response && response.data) {
    try {
      const url = window.URL.createObjectURL(new Blob([response.data]));
      const a = document.createElement('a');
      a.href = url;
      a.download = response.headers.filename;
      a.click();
      a.remove();
      setTimeout(() => window.URL.revokeObjectURL(url), 100);
    } catch (e) {}
  }
};

const handleDownloadError = function(blob: Blob) {
  const fileReader = new FileReader();
  return new Promise((resolve) => {
    fileReader.onload = () => {
      try {
        const message = JSON.parse(fileReader.result as string).message;
        resolve(toasts.addToast({ message, type: 'error', title: 'Error' }));
      } catch (e) {}
    };
    fileReader.readAsText(blob);
  });
};

export { __api, __multi, axios, handleError, handleDownload, handleDownloadError };
