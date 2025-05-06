import React, { useState } from 'react';
import { motion } from 'framer-motion';
import { ProductCard } from '../components/ProductCard';
import { MOCK_PRODUCTS, GENRES } from '../config';
import { Filter, Search } from 'lucide-react';

export function Products() {
  const [selectedGenre, setSelectedGenre] = useState<string>('All');
  const [priceRange, setPriceRange] = useState<string>('all');
  const [sortBy, setSortBy] = useState<string>('featured');
  const [isFilterOpen, setIsFilterOpen] = useState(false);

  const filteredProducts = MOCK_PRODUCTS.filter(product => {
    if (selectedGenre !== 'All' && !product.tags.includes(selectedGenre)) {
      return false;
    }
    if (priceRange === 'under25' && product.price >= 25) {
      return false;
    }
    if (priceRange === '25to50' && (product.price < 25 || product.price > 50)) {
      return false;
    }
    if (priceRange === 'over50' && product.price <= 50) {
      return false;
    }
    return true;
  }).sort((a, b) => {
    if (sortBy === 'priceAsc') return a.price - b.price;
    if (sortBy === 'priceDesc') return b.price - a.price;
    if (sortBy === 'rating') return b.rating - a.rating;
    return 0;
  });

  return (
    <div className="min-h-screen bg-gray-900 pt-24">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Mobile Filter Button */}
        <div className="lg:hidden mb-4">
          <button
            onClick={() => setIsFilterOpen(!isFilterOpen)}
            className="btn-secondary w-full justify-center"
          >
            <Filter className="h-5 w-5 mr-2" />
            Filters
          </button>
        </div>

        <div className="flex flex-col lg:flex-row gap-8">
          {/* Sidebar Filters */}
          <motion.aside
            className={`lg:w-64 flex-shrink-0 ${isFilterOpen ? 'block' : 'hidden'} lg:block`}
            initial={{ x: -20, opacity: 0 }}
            animate={{ x: 0, opacity: 1 }}
          >
            <div className="sticky top-24 card p-6">
              <h2 className="text-xl font-bold text-white mb-6">Filters</h2>
              
              <div className="space-y-6">
                <div>
                  <label className="block text-sm font-medium text-gray-300 mb-2">
                    Genre
                  </label>
                  <select
                    value={selectedGenre}
                    onChange={(e) => setSelectedGenre(e.target.value)}
                    className="w-full rounded-lg border border-gray-700 py-2 px-3 bg-gray-800 text-gray-300 
                             focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                  >
                    <option value="All">All Genres</option>
                    {GENRES.map(genre => (
                      <option key={genre} value={genre}>{genre}</option>
                    ))}
                  </select>
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-300 mb-2">
                    Price Range
                  </label>
                  <select
                    value={priceRange}
                    onChange={(e) => setPriceRange(e.target.value)}
                    className="w-full rounded-lg border border-gray-700 py-2 px-3 bg-gray-800 text-gray-300 
                             focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                  >
                    <option value="all">All Prices</option>
                    <option value="under25">Under $25</option>
                    <option value="25to50">$25 to $50</option>
                    <option value="over50">Over $50</option>
                  </select>
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-300 mb-2">
                    Sort By
                  </label>
                  <select
                    value={sortBy}
                    onChange={(e) => setSortBy(e.target.value)}
                    className="w-full rounded-lg border border-gray-700 py-2 px-3 bg-gray-800 text-gray-300 
                             focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                  >
                    <option value="featured">Featured</option>
                    <option value="priceAsc">Price: Low to High</option>
                    <option value="priceDesc">Price: High to Low</option>
                    <option value="rating">Highest Rated</option>
                  </select>
                </div>
              </div>
            </div>
          </motion.aside>

          {/* Product Grid */}
          <div className="flex-1">
            <div className="mb-8">
              <h1 className="text-3xl font-bold text-white mb-4">Browse Manga</h1>
              <div className="relative">
                <input
                  type="text"
                  placeholder="Search manga titles, authors, or genres..."
                  className="w-full px-4 py-3 rounded-lg bg-gray-800 border border-gray-700 text-white 
                           placeholder-gray-400 focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                />
                <div className="absolute inset-y-0 right-0 flex items-center pr-3">
                  <Search className="h-5 w-5 text-gray-400" />
                </div>
              </div>
            </div>

            <motion.div 
              className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8"
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              transition={{ duration: 0.5 }}
            >
              {filteredProducts.map((product, index) => (
                <motion.div
                  key={product.id}
                  initial={{ opacity: 0, y: 20 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ delay: index * 0.1 }}
                >
                  <ProductCard product={product} />
                </motion.div>
              ))}
            </motion.div>
          </div>
        </div>
      </div>
    </div>
  );
}