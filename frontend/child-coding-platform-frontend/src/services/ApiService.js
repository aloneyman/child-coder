import axios from 'axios';

const apiClient = axios.create({
    baseURL: `http://localhost:8080`, // 根据实际后端服务地址调整
    withCredentials: false, // 这是默认值
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
    }
});

export default {
    register(credentials) {
        return apiClient.post('/register', credentials);
    },
    login(credentials) {
        return apiClient.post('/login', credentials);
    },
    // ... 添加其他 API 方法 ...
};
