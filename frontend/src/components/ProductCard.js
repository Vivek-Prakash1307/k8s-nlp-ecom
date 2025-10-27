import React from 'react';

export default function ProductCard({p, onAdd}) {
  return (
    <div className="card">
      <h3>{p.title}</h3>
      <p>{p.description}</p>
      <p>₹ {p.price.toFixed(2)}</p>
      <p>Rating: {p.rating} • Pop: {p.popularity}</p>
      <button onClick={() => onAdd(p._id || p.id)}>Add to cart</button>
    </div>
  );
}
