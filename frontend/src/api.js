import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add token to every request
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Auth APIs
export const signup = (data) => api.post('/signup', data);
export const login = (data) => api.post('/login', data);

// Product APIs
export const getProducts = (params) => api.get('/products', { params });

// Cart APIs
export const addToCart = (data) => api.post('/cart/add', data);
export const getCart = () => api.get('/cart');

// Order APIs
export const checkout = (data) => api.post('/checkout', data);
export const getOrders = () => api.get('/orders');

export default api;