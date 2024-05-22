import React from 'react';
import './Userprofile.css';
import { useNavigate } from 'react-router-dom';

function Userprofile({handlelogout, name, rollno, email,image }) {
  const navigate = useNavigate();

  const handleLogout = () => {
    handlelogout();
    navigate('/user/login');
  };

  return (
    <div>
      <h2>Your Account</h2>
      {email==""?(
        <p>No user logged in</p>
      ):(

      
      <div className='account-box'>
        <div className='box'>
        <img src={`data:image/png;base64,${image}`} alt="User" />
          <h3>Name: {name}</h3>
          <h3>Roll No: {rollno}</h3>
          <h3>Email: {email}</h3>
          <div className='button1'>
            <button id='account1' onClick={handleLogout}>Log out</button>
          </div>
        </div>
      </div>)}
    </div>
  );
}

export default Userprofile;


