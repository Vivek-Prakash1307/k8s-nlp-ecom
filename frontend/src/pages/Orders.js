import React, {useEffect, useState} from 'react';
import API from '../api';

export default function Orders(){
  const [orders, setOrders] = useState([]);
  useEffect(()=>{ API.get('/api/orders').then(r=>setOrders(r.data.orders)).catch(()=>setOrders([])) }, []);
  return (
    <div>
      <h2>Your Orders</h2>
      {orders.length===0 ? <p>No orders</p> : (
        <ul>
          {orders.map(o => <li key={o.id || o._id}>Order {o._id || o.id} — Total: ₹{o.total} — Status: {o.status}</li>)}
        </ul>
      )}
    </div>
  );
}
