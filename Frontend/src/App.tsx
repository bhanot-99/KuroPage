import React, { useState, useContext } from 'react';
import { Navbar } from './components/Navbar';
import { Home } from './pages/Home';
import { Products } from './pages/Products';
import { AuthProvider, AuthContext } from './context/AuthContext';
import { CartProvider } from './context/CartContext';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

function App() {
  const [currentPage, setCurrentPage] = useState<'home' | 'products'>('home');
  const { isAuthenticated, logout } = useContext(AuthContext);

  const handleNavigate = (page: 'home' | 'products') => {
    setCurrentPage(page);
  };

  return (
    <AuthProvider>
      <CartProvider>
        <Router>
          <div className="min-h-screen bg-gray-900">
            <Navbar
              onNavigate={handleNavigate}
              currentPage={currentPage}
              isAuthenticated={isAuthenticated}
              onLogout={logout}
            />
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/home" element={<Home />} />
              <Route path="/products" element={<Products />} />
            </Routes>
          </div>
        </Router>
      </CartProvider>
    </AuthProvider>
  );
}

export default App;