import React from 'react';
import { Link, useNavigate } from 'react-router-dom';

export default function Header(){
  const navigate = useNavigate();
  const token = localStorage.getItem('token');
  const logout = () => { localStorage.removeItem('token'); navigate('/'); }
  return (
    <header className="header">
      <Link to="/"><h2>DummyShop</h2></Link>
      <nav>
        <Link to="/">Products</Link>
        <Link to="/cart">Cart</Link>
        { token ? <>
          <Link to="/orders">Orders</Link>
          <button onClick={logout} className="linklike">Logout</button>
        </> : <>
          <Link to="/login">Login</Link>
          <Link to="/signup">SignUp</Link>
        </> }
      </nav>
    </header>
  );
}
