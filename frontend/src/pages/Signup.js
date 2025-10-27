import React, {useState} from 'react';
import API from '../api';
import {useNavigate} from 'react-router-dom';

export default function Signup(){
  const [name,setName]=useState(""); const [email,setEmail]=useState(""); const [pass,setPass]=useState("");
  const nav = useNavigate();
  const submit = async e => {
    e.preventDefault();
    try{
      const res = await API.post('/api/signup', {name, email, password: pass});
      localStorage.setItem('token', res.data.token);
      nav('/');
    }catch(err){ alert(err.response?.data?.error || "Error"); }
  };
  return (
    <form onSubmit={submit} className="form">
      <h2>Sign Up</h2>
      <input value={name} onChange={e=>setName(e.target.value)} placeholder="Name"/>
      <input value={email} onChange={e=>setEmail(e.target.value)} placeholder="Email"/>
      <input type="password" value={pass} onChange={e=>setPass(e.target.value)} placeholder="Password"/>
      <button type="submit">Signup</button>
    </form>
  );
}
