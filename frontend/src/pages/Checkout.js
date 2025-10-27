import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { checkout } from '../api';
import { toast } from 'react-toastify';

export default function Checkout() {
  const [address, setAddress] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    if (!address.trim()) {
      toast.warning('Please enter your delivery address');
      return;
    }

    setLoading(true);
    try {
      await checkout({ address });
      toast.success('Order placed successfully! 🎉');
      navigate('/orders');
    } catch (error) {
      toast.error(error.response?.data?.error || 'Checkout failed');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="checkout-container">
      <h1>💳 Checkout</h1>
      <form onSubmit={handleSubmit} className="checkout-form">
        <div className="form-group">
          <label>Delivery Address</label>
          <textarea
            placeholder="Enter your complete delivery address..."
            value={address}
            onChange={(e) => setAddress(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="btn-primary" disabled={loading}>
          {loading ? 'Placing Order...' : 'Place Order 🚀'}
        </button>
      </form>
    </div>
  );
}