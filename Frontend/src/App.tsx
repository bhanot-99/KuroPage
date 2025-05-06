import React, { useState } from 'react';
import { Navbar } from './components/Navbar';
import { Home } from './pages/Home';
import { Products } from './pages/Products';
import { AuthProvider } from './context/AuthContext';
import { CartProvider } from './context/CartContext';

function App() {
  const [currentPage, setCurrentPage] = useState<'home' | 'products'>('home');

  return (
    <AuthProvider>
      <CartProvider>
        <div className="min-h-screen bg-gray-900">
          <Navbar onNavigate={setCurrentPage} currentPage={currentPage} />
          {currentPage === 'home' ? <Home /> : <Products />}
        </div>
      </CartProvider>
    </AuthProvider>
  );
}

export default App;