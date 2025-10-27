import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { getCart, getProducts } from '../api';
import { toast } from 'react-toastify';

export default function Cart() {
  const [cart, setCart] = useState({ items: [] });
  const [products, setProducts] = useState({});
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    fetchCartAndProducts();
  }, []);

  const fetchCartAndProducts = async () => {
    try {
      const cartRes = await getCart();
      const cartData = cartRes.data.cart;
      setCart(cartData);

      // Fetch product details for cart items
      if (cartData.items && cartData.items.length > 0) {
        const productsRes = await getProducts();
        const productsMap = {};
        productsRes.data.products.forEach((p) => {
          productsMap[p.id] = p;
        });
        setProducts(productsMap);
      }
    } catch (error) {
      toast.error('Failed to fetch cart');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const calculateTotal = () => {
    return cart.items.reduce((total, item) => {
      const product = products[item.productId];
      return total + (product ? product.price * item.quantity : 0);
    }, 0);
  };

  const handleCheckout = () => {
    if (cart.items.length === 0) {
      toast.warning('Your cart is empty');
      return;
    }
    navigate('/checkout');
  };

  if (loading) {
    return (
      <div className="loading">
        <div className="spinner"></div>
      </div>
    );
  }

  if (!cart.items || cart.items.length === 0) {
    return (
      <div className="cart-container">
        <div className="empty-state">
          <div className="empty-state-icon">🛒</div>
          <h3>Your cart is empty</h3>
          <p>Add some products to get started!</p>
          <button className="btn-primary" onClick={() => navigate('/')}>
            Start Shopping
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="cart-container">
      <h1>🛒 Shopping Cart</h1>
      <div className="cart-items">
        {cart.items.map((item, index) => {
          const product = products[item.productId];
          if (!product) return null;

          return (
            <div key={index} className="cart-item">
              <div className="cart-item-image">
                {product.title.charAt(0)}
              </div>
              <div className="cart-item-details">
                <h3 className="cart-item-title">{product.title}</h3>
                <p className="cart-item-price">${product.price.toFixed(2)}</p>
                <p className="cart-item-quantity">Quantity: {item.quantity}</p>
                <p className="cart-item-quantity">
                  Subtotal: ${(product.price * item.quantity).toFixed(2)}
                </p>
              </div>
            </div>
          );
        })}
      </div>
      <div className="cart-summary">
        <div className="cart-total">
          <span>Total Amount:</span>
          <span className="cart-total-amount">${calculateTotal().toFixed(2)}</span>
        </div>
        <button className="btn-checkout" onClick={handleCheckout}>
          Proceed to Checkout 💳
        </button>
      </div>
    </div>
  );
}