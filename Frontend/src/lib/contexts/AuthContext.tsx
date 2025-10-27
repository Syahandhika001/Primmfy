'use client';

import React, { createContext, useContext, useState, useEffect, useCallback } from 'react';
import { useRouter } from 'next/navigation';
import Cookies from 'js-cookie';
import { authApi } from '../api/auth';
import { User, AuthContextType, RegisterRequest } from '../types/auth';

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const router = useRouter();

  // Load user from cookies on mount
  useEffect(() => {
    const loadUser = () => {
      const savedToken = Cookies.get('token');
      const savedUser = Cookies.get('user');

      if (savedToken && savedUser) {
        try {
          const userData = JSON.parse(savedUser);
          setToken(savedToken);
          setUser(userData);
        } catch (error) {
          console.error('Failed to parse user data:', error);
          Cookies.remove('token');
          Cookies.remove('user');
        }
      }
      setIsLoading(false);
    };

    loadUser();
  }, []);

  // Login function
  const login = useCallback(async (email: string, password: string) => {
    try {
      const response = await authApi.login({ email, password });
      
      // Save token and user to cookies
      Cookies.set('token', response.token, { expires: 7 }); // 7 days
      Cookies.set('user', JSON.stringify(response.user), { expires: 7 });
      
      setToken(response.token);
      setUser(response.user);

      // Redirect based on role
      if (response.user.role === 'teacher') {
        router.push('/teacher/dashboard');
      } else {
        router.push('/student/dashboard');
      }
    } catch (error) {
      throw error;
    }
  }, [router]);

  // Register function
  const register = useCallback(async (data: RegisterRequest) => {
    try {
      const response = await authApi.register(data);
      
      // Save token and user to cookies
      Cookies.set('token', response.token, { expires: 7 });
      Cookies.set('user', JSON.stringify(response.user), { expires: 7 });
      
      setToken(response.token);
      setUser(response.user);

      // Redirect based on role
      if (response.user.role === 'teacher') {
        router.push('/teacher/dashboard');
      } else {
        router.push('/student/dashboard');
      }
    } catch (error) {
      throw error;
    }
  }, [router]);

  // Logout function
  const logout = useCallback(() => {
    Cookies.remove('token');
    Cookies.remove('user');
    setToken(null);
    setUser(null);
    router.push('/login');
  }, [router]);

  const value: AuthContextType = {
    user,
    token,
    isLoading,
    isAuthenticated: !!token && !!user,
    login,
    register,
    logout,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

// Custom hook to use auth context
export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}