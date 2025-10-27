import axiosInstance from './axios';
import { LoginRequest, RegisterRequest, AuthResponse, User } from '../types/auth';

export const authApi = {
  // Login
  login: async (data: LoginRequest): Promise<AuthResponse> => {
    const response = await axiosInstance.post<AuthResponse>('/login', data);
    return response.data;
  },

  // Register
  register: async (data: RegisterRequest): Promise<AuthResponse> => {
    const response = await axiosInstance.post<AuthResponse>('/register', data);
    return response.data;
  },

  // Get current user profile
  getProfile: async (): Promise<{ user: User }> => {
    const response = await axiosInstance.get<{ user: User }>('/profile');
    return response.data;
  },
};