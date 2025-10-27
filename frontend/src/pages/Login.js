import React, {useState} from 'react';
import API from '../api';
import {useNavigate} from 'react-router-dom';

export default function Login(){
  const [email,setEmail]=useState(""); const [pass,setPass]=useState("");
  const nav = useNavigate();
  const submit = async e => {
    e.preventDefault();
    try{
      const res = await API.post('/api/login', {email, password: pass});
      localStorage.setItem('token', res.data.token);
      nav('/');
    }catch(err){ alert(err.response?.data?.error || "Error"); }
  };
  return (
    <form onSubmit={submit} className="form">
      <h2>Login</h2>
      <input value={email} onChange={e=>setEmail(e.target.value)} placeholder="Email"/>
      <input type="password" value={pass} onChange={e=>setPass(e.target.value)} placeholder="Password"/>
      <button type="submit">Login</button>
    </form>
  );
}
