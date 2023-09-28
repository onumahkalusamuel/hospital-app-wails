import { __api, __multi, handleError, handleDownload, handleDownloadError } from './axios';

const get = async (url: string, filter?: Record<string, string>) => {
  let queryString = new URLSearchParams(filter).toString();
  return await __api.get(`${url}${queryString ? '?' + queryString : ''}`)
    .then(response => { return response?.data; })
    .catch(e => { return handleError(e); });
};

const getBlob = async (url: string, filter?: Record<string, string>) => {
  let queryString = new URLSearchParams(filter).toString();
  return await __api.get(`${url}${queryString ? '?' + queryString : ''}`, { responseType: 'blob' })
    .then(blob => { return handleDownload(blob); })
    .catch(e => { return handleDownloadError(e.response.data); });
};

const post = async (url: string, formData: object) => {
  return await __api.post(url, formData)
    .then(response => { return response?.data; })
    .catch(e => { return handleError(e); });
};

const postBlob = async (url: string, formData: object) => {
  return await __api.post(url, formData, { responseType: 'blob' })
    .then(blob => { return handleDownload(blob); })
    .catch(e => { return handleDownloadError(e.response.data); });
};

const postMulti = async (url: string, form: object) => {
  return await __multi.post(url, form)
    .then(response => { return response?.data; })
    .catch(e => handleError(e));
};

const putMulti = async (url: string, form: object) => {
  return await __multi.put(url, form)
    .then(response => { return response?.data; })
    .catch(e => handleError(e));
};

const deleteRecord = async (url: string, silent?: boolean) => {
  if(!silent && !confirm('Are you sure you want to delete this record?')) return;
  return await __api.delete(url)
    .then(response => { return response?.data; })
    .catch(e => { return handleError(e); });
};

const put = async (url: string, formData: object) => {
  return await __api.put(url, formData)
    .then(response => { return response?.data; })
    .catch(e => { return handleError(e); });
};

const apiRequest = { get, getBlob, post, postBlob, postMulti, putMulti, deleteRecord, put };

export default apiRequest;
