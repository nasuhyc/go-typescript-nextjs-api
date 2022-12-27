import axios from "axios";
import Router from "next/router";

//* Create an axios instance with a env variable for the baseURL and a timeout of 1 second
const instance = axios.create({
  baseURL: process.env.BACKEND_API,
  timeout: 1000,
  headers: { "X-Custom-Header": "foobar" },
});

// Add a request interceptor
instance.interceptors.request.use(
  function (config) {
    // Do something before request is sent
    return config;
  },
  function (error) {
    // Do something with request error
    return Promise.reject(error);
  }
);

// Add a response interceptor
instance.interceptors.response.use(
  function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data
    return response;
  },
  function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error

    if (error.response.status === 400 || error.response.status === 404) {
      alert(error.response.data.message);
    }

    // Redirect to 500 page if the server is down
    if (error.response.status === 500) {
      Router.push("/500");
    }

    return Promise.reject(error);
  }
);

export default instance;
