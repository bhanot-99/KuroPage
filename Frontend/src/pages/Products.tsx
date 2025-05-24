import React, { useState, useEffect } from 'react';
import { motion } from 'framer-motion';
import { Filter, Search } from 'lucide-react';

interface Manga {
  id: string;
  attributes: {
    titles: { en_jp?: string };
    posterImage: { small: string };
    averageRating?: string;
    // Add more fields as needed
  };
}

export function Products() {
  const [manga, setManga] = useState<Manga[]>([]);
  const [loading, setLoading] = useState(true);
  const [selectedGenre, setSelectedGenre] = useState<string>('All');
  const [sortBy, setSortBy] = useState<string>('featured');
  const [isFilterOpen, setIsFilterOpen] = useState(false);
  const [search, setSearch] = useState('');

  useEffect(() => {
    // Fetch first 20 popular manga
    fetch('https://kitsu.io/api/edge/manga?page[limit]=20&page[offset]=0&sort=popularityRank')
      .then(res => res.json())
      .then(data1 => {
        // Fetch next 20 popular manga
        fetch('https://kitsu.io/api/edge/manga?page[limit]=20&page[offset]=20&sort=popularityRank')
          .then(res => res.json())
          .then(data2 => {
            setManga([...data1.data, ...data2.data]);
            setLoading(false);
          });
      });
  }, []);
  // Filter and sort logic
  const filteredManga = manga
    .filter(item => {
      if (selectedGenre !== 'All' && !item.attributes.titles.en_jp?.toLowerCase().includes(selectedGenre.toLowerCase())) {
        return false;
      }
      if (search && !item.attributes.titles.en_jp?.toLowerCase().includes(search.toLowerCase())) {
        return false;
      }
      return true;
    })
    .sort((a, b) => {
      if (sortBy === 'rating') {
        return (parseFloat(b.attributes.averageRating || '0') - parseFloat(a.attributes.averageRating || '0'));
      }
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
                    Genre (by title keyword)
                  </label>
                  <select
                    value={selectedGenre}
                    onChange={(e) => setSelectedGenre(e.target.value)}
                    className="w-full rounded-lg border border-gray-700 py-2 px-3 bg-gray-800 text-gray-300 
                             focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                  >
                    <option value="All">All Genres</option>
                    <option value="Naruto">Naruto</option>
                    <option value="Attack">Attack</option>
                    <option value="Dragon">Dragon</option>
                    {/* Add more keywords or genres as needed */}
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
                  placeholder="Search manga titles..."
                  value={search}
                  onChange={e => setSearch(e.target.value)}
                  className="w-full px-4 py-3 rounded-lg bg-gray-800 border border-gray-700 text-white 
                           placeholder-gray-400 focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                />
                <div className="absolute inset-y-0 right-0 flex items-center pr-3">
                  <Search className="h-5 w-5 text-gray-400" />
                </div>
              </div>
            </div>

            {loading ? (
              <div className="text-white">Loading manga...</div>
            ) : (
              <motion.div
                className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8"
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{ duration: 0.5 }}
              >
                {filteredManga.map((item, index) => (
                  <motion.div
                    key={item.id}
                    initial={{ opacity: 0, y: 20 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ delay: index * 0.05 }}
                  >
                    <div className="bg-gray-800 rounded shadow p-2 flex flex-col items-center">
                      <img
                        src={item.attributes.posterImage.small}
                        alt={item.attributes.titles.en_jp || 'Manga'}
                        className="w-full h-48 object-cover rounded"
                      />
                      <h3 className="text-white mt-2 text-center text-lg font-semibold">
                        {item.attributes.titles.en_jp || 'No Title'}
                      </h3>
                      {item.attributes.averageRating && (
                        <p className="text-gray-400 text-sm mt-1">
                          Rating: {item.attributes.averageRating}
                        </p>
                      )}
                    </div>
                  </motion.div>
                ))}
              </motion.div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}