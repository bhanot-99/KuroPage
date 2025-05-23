import React, { useState, useEffect } from 'react';
import { ShoppingCart, User, Menu, X } from 'lucide-react';
import { useAuth } from '../context/AuthContext';
import { useCart } from '../context/CartContext';
import { motion } from 'framer-motion';
// Add this import:
import { GoogleLogin } from '@react-oauth/google';
import { Link } from 'react-router-dom';

type NavbarProps = {
  onNavigate: (page: 'home' | 'products') => void;
  currentPage: 'home' | 'products';
  isAuthenticated?: boolean;
  onLogout?: () => void;
};

export const Navbar: React.FC<NavbarProps> = ({
  onNavigate,
  currentPage,
  isAuthenticated,
  onLogout,
}) => {
  const { user, logout, loginWithGoogle } = useAuth();
  const { items } = useCart();
  const [isScrolled, setIsScrolled] = useState(false);
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      setIsScrolled(window.scrollY > 10);
    };
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  // Handler for Google login success
  const handleGoogleLoginSuccess = (credentialResponse: any) => {
    // Parse credentialResponse if needed
    if (loginWithGoogle) {
      loginWithGoogle(credentialResponse);
    }
  };

  // Handler for Google login error
  const handleGoogleLoginError = () => {
    console.log('Google login failed');
  };

  return (
    <motion.nav 
      className={`fixed w-full z-50 transition-all duration-200 ${
        isScrolled ? 'glass-nav' : 'bg-transparent'
      }`}
      initial={{ y: -100 }}
      animate={{ y: 0 }}
      transition={{ duration: 0.5 }}
    >
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center">
            <Link to="/home" className="navbar-brand">
              <span className="text-2xl font-display bg-gradient-to-r from-indigo-400 to-violet-400 
                             bg-clip-text text-transparent">
                MangaVerse
              </span>
            </Link>
            <div className="hidden sm:ml-8 sm:flex sm:space-x-4">
              <Link 
                to="/home"
                className={`nav-link ${currentPage === 'home' ? 'text-white' : ''}`}
              >
                Home
              </Link>
              <Link
                to="/products"
                className={`nav-link ${currentPage === 'products' ? 'text-white' : ''}`}
              >
                Browse Manga
              </Link>
            </div>
          </div>
          
          <div className="hidden sm:flex sm:items-center sm:space-x-4">
            <button className="relative p-2 text-gray-400 hover:text-white transition-colors duration-200">
              <ShoppingCart className="h-6 w-6" />
              {items.length > 0 && (
                <span className="absolute -top-1 -right-1 flex items-center justify-center w-5 h-5 text-xs 
                               font-bold text-white bg-indigo-600 rounded-full animate-fade-in">
                  {items.length}
                </span>
              )}
            </button>
            
            {user ? (
              <div className="flex items-center space-x-4">
                <div className="flex items-center space-x-2 px-4 py-2 rounded-lg bg-gray-800">
                  <User className="h-5 w-5 text-gray-400" />
                  <span className="text-sm font-medium text-white">{user.name}</span>
                </div>
                <button
                  onClick={logout}
                  className="btn-secondary"
                >
                  Logout
                </button>
              </div>
            ) : (
              // Replace the button with GoogleLogin
              <GoogleLogin
                onSuccess={handleGoogleLoginSuccess}
                onError={handleGoogleLoginError}
                width="140"
                theme="filled_black"
                size="medium"
                text="signin_with"
                shape="pill"
              />
            )}
          </div>

          <div className="flex items-center sm:hidden">
            <button
              onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
              className="p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-800 
                       focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500"
            >
              {isMobileMenuOpen ? (
                <X className="h-6 w-6" />
              ) : (
                <Menu className="h-6 w-6" />
              )}
            </button>
          </div>
        </div>
      </div>

      {/* Mobile menu */}
      {isMobileMenuOpen && (
        <motion.div 
          className="sm:hidden bg-gray-800 border-t border-gray-700"
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.2 }}
        >
          <div className="px-2 pt-2 pb-3 space-y-1">
            <Link
              to="/home"
              onClick={() => setIsMobileMenuOpen(false)}
              className="block w-full text-left px-3 py-2 rounded-md text-base font-medium text-white 
                         hover:bg-gray-700"
            >
              Home
            </Link>
            <Link
              to="/products"
              onClick={() => setIsMobileMenuOpen(false)}
              className="block w-full text-left px-3 py-2 rounded-md text-base font-medium text-white 
                         hover:bg-gray-700"
            >
              Browse Manga
            </Link>
          </div>
        </motion.div>
      )}
    </motion.nav>
  );
}