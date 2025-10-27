import React, {useState} from 'react';
import API from '../api';
import {useNavigate} from 'react-router-dom';

export default function Checkout(){
  const [address,setAddress] = useState("");
  const nav = useNavigate();

  const submit = async (e) => {
    e.preventDefault();
    try {
      const res = await API.post('/api/checkout', { address });
      alert('Order placed');
      nav('/orders');
    } catch(err) {
      alert(err.response?.data?.error || "Error");
    }
  };

  return (
    <form onSubmit={submit} className="form">
      <h2>Checkout</h2>
      <textarea value={address} onChange={e=>setAddress(e.target.value)} placeholder="Shipping address"/>
      <button type="submit">Place Order</button>
    </form>
  );
}
