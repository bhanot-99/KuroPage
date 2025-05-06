import React from 'react';
import { motion } from 'framer-motion';
import { Truck, Shield, Gift, ArrowRight, Star } from 'lucide-react';
import { MOCK_PRODUCTS } from '../config';

const fadeIn = {
  initial: { opacity: 0, y: 20 },
  animate: { opacity: 1, y: 0 },
  transition: { duration: 0.6 }
};

const staffPicks = MOCK_PRODUCTS.slice(0, 4);

export function Home() {
  return (
    <div className="bg-gray-900">
      {/* Hero Section */}
      <div className="relative min-h-[80vh] flex items-center">
        <div className="absolute inset-0 overflow-hidden">
          <img
            src="https://images.unsplash.com/photo-1581833971358-2c8b550f87b3?auto=format&fit=crop&q=80&w=2000"
            alt="Manga Collection"
            className="w-full h-full object-cover opacity-20"
          />
          <div className="absolute inset-0 bg-gradient-to-b from-gray-900/80 via-gray-900/60 to-gray-900" />
        </div>
        
        <motion.div 
          className="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-32"
          initial="initial"
          animate="animate"
          variants={fadeIn}
        >
          <div className="text-center max-w-3xl mx-auto">
            <h1 className="text-5xl md:text-7xl font-bold tracking-tight text-white mb-8">
              Discover Rare & <span className="bg-gradient-to-r from-indigo-400 to-violet-400 bg-clip-text text-transparent">
                Bestselling Manga
              </span>
            </h1>
            <p className="text-xl text-gray-300 mb-12">
              Your gateway to exclusive editions, limited releases, and the finest manga collection in the world.
            </p>
            <a href="/products" className="btn-primary text-lg px-8 py-4">
              Browse Collection <ArrowRight className="ml-2 h-6 w-6" />
            </a>
          </div>
        </motion.div>
      </div>

      {/* Value Propositions */}
      <div className="py-24 bg-gray-900/50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <motion.div 
              className="card p-6 text-center"
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ delay: 0.2 }}
            >
              <Truck className="h-12 w-12 mx-auto mb-4 text-indigo-400" />
              <h3 className="text-xl font-bold text-white mb-2">Worldwide Shipping</h3>
              <p className="text-gray-400">Fast delivery to your doorstep, wherever you are</p>
            </motion.div>
            
            <motion.div 
              className="card p-6 text-center"
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ delay: 0.4 }}
            >
              <Shield className="h-12 w-12 mx-auto mb-4 text-indigo-400" />
              <h3 className="text-xl font-bold text-white mb-2">Authentic Editions</h3>
              <p className="text-gray-400">100% genuine manga from official publishers</p>
            </motion.div>
            
            <motion.div 
              className="card p-6 text-center"
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ delay: 0.6 }}
            >
              <Gift className="h-12 w-12 mx-auto mb-4 text-indigo-400" />
              <h3 className="text-xl font-bold text-white mb-2">Member Rewards</h3>
              <p className="text-gray-400">Exclusive discounts and early access</p>
            </motion.div>
          </div>
        </div>
      </div>

      {/* Staff Picks */}
      <div className="py-24">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <motion.div 
            className="text-center mb-16"
            initial={{ opacity: 0 }}
            whileInView={{ opacity: 1 }}
            viewport={{ once: true }}
          >
            <h2 className="text-4xl font-bold text-white mb-4">Staff Picks</h2>
            <p className="text-xl text-gray-400">Curated selections from our manga experts</p>
          </motion.div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            {staffPicks.map((manga, index) => (
              <motion.div
                key={manga.id}
                className="card overflow-hidden"
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true }}
                transition={{ delay: index * 0.2 }}
              >
                <img
                  src={manga.imageUrl}
                  alt={manga.name}
                  className="w-full h-64 object-cover"
                />
                <div className="p-6">
                  <div className="flex items-center gap-1 text-amber-400 mb-2">
                    <Star className="h-4 w-4 fill-current" />
                    <span className="text-sm font-medium">{manga.rating}</span>
                  </div>
                  <h3 className="text-lg font-bold text-white mb-2">{manga.name}</h3>
                  <p className="text-sm text-gray-400 mb-4">"{manga.description}"</p>
                  <a href={`/products#${manga.id}`} className="btn-secondary w-full justify-center">
                    View Details
                  </a>
                </div>
              </motion.div>
            ))}
          </div>
        </div>
      </div>

      {/* Newsletter */}
      <div className="py-24 bg-gray-800/50">
        <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
          >
            <h2 className="text-3xl font-bold text-white mb-4">Stay Updated</h2>
            <p className="text-gray-400 mb-8">Get notified about new releases and exclusive offers</p>
            <form className="flex gap-4 max-w-md mx-auto">
              <input
                type="email"
                placeholder="Enter your email"
                className="flex-1 px-4 py-3 rounded-lg bg-gray-800 border border-gray-700 text-white 
                         placeholder-gray-400 focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
              />
              <button type="submit" className="btn-primary whitespace-nowrap">
                Subscribe
              </button>
            </form>
          </motion.div>
        </div>
      </div>
    </div>
  );
}