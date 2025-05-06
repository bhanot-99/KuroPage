import { Product } from '../types';

export const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

export const MOCK_PRODUCTS: Product[] = [
  {
    id: '1',
    name: 'One Piece Vol. 100 - Collector\'s Edition',
    description: 'The 100th volume of the bestselling manga series featuring exclusive artwork and bonus content',
    price: 29.99,
    originalPrice: 39.99,
    category: 'Shonen',
    imageUrl: 'https://images.unsplash.com/photo-1612404730960-5c71577fca11?auto=format&fit=crop&q=80&w=500',
    stock: 50,
    author: 'Eiichiro Oda',
    tags: ['Action', 'Adventure', 'Bestseller'],
    isNewRelease: true,
    rating: 4.9,
    reviewCount: 128,
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  },
  {
    id: '2',
    name: 'Demon Slayer Complete Box Set',
    description: 'All 23 volumes of the hit series in a premium collector\'s box with exclusive art prints',
    price: 199.99,
    originalPrice: 249.99,
    category: 'Shonen',
    imageUrl: 'https://images.unsplash.com/photo-1613376023733-0a73315d9b06?auto=format&fit=crop&q=80&w=500',
    stock: 25,
    author: 'Koyoharu Gotouge',
    tags: ['Action', 'Supernatural', 'Limited Edition'],
    isNewRelease: false,
    rating: 4.8,
    reviewCount: 256,
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  },
  // Add 18 more manga products here with similar structure
];

export const GENRES = [
  'Shonen',
  'Seinen',
  'Shoujo',
  'Josei',
  'Action',
  'Romance',
  'Horror',
  'Slice of Life',
  'Sports',
  'Mystery',
];