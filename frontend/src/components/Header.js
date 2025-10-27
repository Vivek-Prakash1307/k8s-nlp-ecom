import React from 'react';
import { Link, useNavigate } from 'react-router-dom';

export default function Header() {
  const navigate = useNavigate();
  const token = localStorage.getItem('token');
  const user = JSON.parse(localStorage.getItem('user') || '{}');

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    navigate('/login');
  };

  return (
    <header className="header">
      <div className="header-content">
        <Link to="/" className="logo">
          <span>🛒</span> NLP E-Commerce
        </Link>
        <nav>
          <ul className="nav-links">
            <li><Link to="/">Products</Link></li>
            {token ? (
              <>
                <li><Link to="/cart">🛍️ Cart</Link></li>
                <li><Link to="/orders">📦 Orders</Link></li>
                <li>
                  <span style={{ color: '#6366f1', fontWeight: '600' }}>
                    👤 {user.name || 'User'}
                  </span>
                </li>
                <li>
                  <button className="btn-logout" onClick={handleLogout}>
                    Logout
                  </button>
                </li>
              </>
            ) : (
              <>
                <li><Link to="/login">Login</Link></li>
                <li><Link to="/signup">Sign Up</Link></li>
              </>
            )}
          </ul>
        </nav>
      </div>
    </header>
  );
}