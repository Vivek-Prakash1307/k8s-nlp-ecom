import React, {useEffect, useState} from 'react';
import API from '../api';
import ProductCard from '../components/ProductCard';

export default function Products(){
  const [products, setProducts] = useState([]);
  const [sort, setSort] = useState('relevance');
  const [q, setQ] = useState('');

  useEffect(()=> {
    fetchProducts();
  }, [sort]);

  const fetchProducts = async () => {
    const res = await API.get('/api/products', { params: { sort, q }});
    setProducts(res.data.products);
  };

  const addToCart = async (productId) => {
    try{
      await API.post('/api/cart/add', { productId, quantity: 1 });
      alert('Added to cart');
    }catch(err){
      alert('Please login to add to cart');
    }
  };

  const onSearch = (e) => {
    e.preventDefault();
    fetchProducts();
  };

  return (
    <div>
      <div className="toolbar">
        <form onSubmit={onSearch}>
          <input value={q} onChange={e=>setQ(e.target.value)} placeholder="Search products"/>
          <button>Search</button>
        </form>
        <select value={sort} onChange={e=>setSort(e.target.value)}>
          <option value="relevance">Relevance</option>
          <option value="price_asc">Price low→high</option>
          <option value="price_desc">Price high→low</option>
          <option value="rating">Rating</option>
          <option value="popularity">Popularity</option>
        </select>
      </div>
      <div className="grid">
        {products.map(p => <ProductCard key={p.id || p._id} p={p} onAdd={addToCart} />)}
      </div>
    </div>
  );
}
