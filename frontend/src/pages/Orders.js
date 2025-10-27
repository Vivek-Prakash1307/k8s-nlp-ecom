import React, { useState, useEffect } from 'react';
import { getOrders } from '../api';
import { toast } from 'react-toastify';

export default function Orders() {
  const [orders, setOrders] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchOrders();
  }, []);

  const fetchOrders = async () => {
    try {
      const res = await getOrders();
      setOrders(res.data.orders || []);
    } catch (error) {
      toast.error('Failed to fetch orders');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const formatDate = (timestamp) => {
    return new Date(timestamp * 1000).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  if (loading) {
    return (
      <div className="loading">
        <div className="spinner"></div>
      </div>
    );
  }

  if (orders.length === 0) {
    return (
      <div className="orders-container">
        <div className="empty-state">
          <div className="empty-state-icon">📦</div>
          <h3>No orders yet</h3>
          <p>Start shopping to see your orders here!</p>
        </div>
      </div>
    );
  }

  return (
    <div className="orders-container">
      <h1>📦 My Orders</h1>
      {orders.map((order) => (
        <div key={order.id} className="order-card">
          <div className="order-header">
            <span className="order-id">Order #{order.id.slice(-8)}</span>
            <span className={`order-status ${order.status}`}>
              {order.status}
            </span>
          </div>
          <div className="order-info">
            <strong>Date:</strong> {formatDate(order.createdAt)}
          </div>
          <div className="order-info">
            <strong>Items:</strong> {order.items.length} item(s)
          </div>
          <div className="order-info">
            <strong>Delivery Address:</strong> {order.address}
          </div>
          <div className="order-total">
            Total: ${order.total.toFixed(2)}
          </div>
        </div>
      ))}
    </div>
  );
}