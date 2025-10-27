import React, {useEffect, useState} from 'react';
import API from '../api';
import {useNavigate} from 'react-router-dom';

export default function Cart(){
  const [cart, setCart] = useState({items:[]});
  const nav = useNavigate();
  const fetch = async () => {
    try {
      const res = await API.get('/api/cart');
      setCart(res.data.cart);
    } catch(e) { setCart({items:[]}); }
  };
  useEffect(()=>{ fetch(); }, []);
  const goCheckout = ()=> nav('/checkout');

  return (
    <div>
      <h2>Your Cart</h2>
      {cart.items.length===0 ? <p>Empty</p> : <>
        <ul>
          {cart.items.map(it => <li key={it.productId}>{it.productId} x {it.quantity}</li>)}
        </ul>
        <button onClick={goCheckout}>Checkout</button>
      </>}
    </div>
  );
}
