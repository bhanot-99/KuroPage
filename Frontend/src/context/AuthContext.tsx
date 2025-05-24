import React, { createContext, useContext, useState, ReactNode } from 'react';
import { User } from '../types';

interface AuthContextType {
  user: User | null;
  isAuthenticated: boolean;
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
  loginWithGoogle?: (googleUser: any) => void; // Add this
}

export const AuthContext = createContext<AuthContextType>({
  user: null,
  isAuthenticated: false,
  login: async () => {},
  logout: () => {},
});

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null);

  const login = async (email: string, password: string) => {
    // Simulate API call
    const mockUser: User = {
      id: '1',
      email,
      name: 'John Doe',
    };
    setUser(mockUser);
    console.log("login success");
  };

  const logout = () => {
    setUser(null);
  };

  const loginWithGoogle = (googleUser: any) => {
    const mockUser: User = {
      id: googleUser.sub || 'google-id',
      email: googleUser.email || 'google@example.com',
      name: googleUser.name || 'Google User',
    };
    setUser(mockUser);
    console.log("login success (Google)");
  };

  const isAuthenticated = !!user;

  return (
    <AuthContext.Provider value={{ user, isAuthenticated, login, logout, loginWithGoogle }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}