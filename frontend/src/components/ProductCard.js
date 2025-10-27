import React from 'react';
import { addToCart } from '../api';
import { toast } from 'react-toastify';

export default function ProductCard({ product }) {
  const handleAddToCart = async () => {
    const token = localStorage.getItem('token');
    if (!token) {
      toast.warning('Please login to add items to cart');
      return;
    }

    try {
      await addToCart({
        productId: product.id,
        quantity: 1
      });
      toast.success(`${product.title} added to cart!`);
    } catch (error) {
      toast.error('Failed to add to cart');
      console.error(error);
    }
  };

  return (
    <div className="product-card">
      <div className="product-image">
        {product.title.charAt(0)}
      </div>
      <div className="product-content">
        <span className="product-category">{product.category}</span>
        <h3 className="product-title">{product.title}</h3>
        <p className="product-description">{product.description}</p>
        <div className="product-meta">
          <span className="product-price">${product.price.toFixed(2)}</span>
          <span className="product-rating">
            ⭐ {product.rating.toFixed(1)}
          </span>
        </div>
        <button className="btn-add-cart" onClick={handleAddToCart}>
          Add to Cart 🛒
        </button>
      </div>
    </div>
  );
}