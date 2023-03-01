import axios from 'axios';

const API = axios.create({ baseURL: 'http://localhost:8080' });

export const fetchPosts = () => API.get('/posts');
export const createPost = (values) => API.post('/posts', values);
export const deletePost = (id) => API.delete(`/posts/${id}`, id);