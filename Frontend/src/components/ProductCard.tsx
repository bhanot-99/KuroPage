import React from 'react';
import { Product } from '../types';
import { useCart } from '../context/CartContext';
import { ShoppingCart, Heart, Star } from 'lucide-react';

interface ProductCardProps {
  product: Product;
}

export function ProductCard({ product }: ProductCardProps) {
  const { addToCart } = useCart();
  const [isHovered, setIsHovered] = React.useState(false);

  const discount = product.originalPrice 
    ? Math.round(((product.originalPrice - product.price) / product.originalPrice) * 100)
    : 0;

  return (
    <div
      className="card group"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      <div className="relative overflow-hidden rounded-t-xl">
        <img
          src={product.imageUrl}
          alt={product.name}
          className="w-full h-[400px] object-cover transform transition-transform duration-300 group-hover:scale-105"
        />
        <div className="absolute top-4 right-4 flex flex-col gap-2">
          <button className="p-2 rounded-full bg-gray-900/80 backdrop-blur-sm shadow-sm hover:bg-gray-800 
                           transition-all duration-200">
            <Heart className="h-5 w-5 text-gray-300" />
          </button>
        </div>
        {product.isNewRelease && (
          <div className="absolute top-4 left-4">
            <span className="badge-new">New Release</span>
          </div>
        )}
        {product.rating >= 4.8 && (
          <div className="absolute top-14 left-4">
            <span className="badge-bestseller">Bestseller</span>
          </div>
        )}
        <div className="absolute inset-0 bg-gradient-to-t from-gray-900 to-transparent opacity-0 
                       group-hover:opacity-100 transition-opacity duration-300" />
      </div>
      
      <div className="p-6">
        <div className="flex items-start justify-between mb-2">
          <div>
            <h3 className="text-lg font-semibold text-white group-hover:text-indigo-400 
                         transition-colors duration-200">
              {product.name}
            </h3>
            <p className="text-sm text-gray-400">by {product.author}</p>
          </div>
          <div className="text-right">
            <div className="flex items-center gap-1 text-amber-400">
              <Star className="h-4 w-4 fill-current" />
              <span className="text-sm font-medium">{product.rating}</span>
            </div>
            <div className="text-sm text-gray-400">
              {product.reviewCount} reviews
            </div>
          </div>
        </div>
        
        <div className="flex flex-wrap gap-2 mb-4">
          {product.tags.map(tag => (
            <span key={tag} className="tag">
              {tag}
            </span>
          ))}
        </div>
        
        <div className="flex items-baseline gap-2 mb-4">
          <span className="text-2xl font-bold text-white">${product.price}</span>
          {product.originalPrice && (
            <>
              <span className="text-lg text-gray-500 line-through">
                ${product.originalPrice}
              </span>
              <span className="text-sm text-emerald-500 font-medium">
                Save {discount}%
              </span>
            </>
          )}
        </div>
        
        <div className="flex items-center justify-between text-sm text-gray-400 mb-4">
          <span>{product.stock} in stock</span>
          <span>{product.category}</span>
        </div>
        
        <button
          onClick={() => addToCart(product, 1)}
          className="mt-2 w-full btn-primary group/button"
        >
          <ShoppingCart className="h-5 w-5 mr-2 transition-transform duration-200 
                                group-hover/button:scale-110" />
          <span>Add to Cart</span>
        </button>
      </div>
    </div>
  );
}