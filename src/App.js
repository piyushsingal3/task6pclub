import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Loginadmin from './components/Loginadmin';
import Loginuser from './components/Loginuser';
import axios from "axios";
import { useState } from 'react';
import AdminDashboard from './components/Admindashboard';
import SignUp from './components/Signup';
import Markattendance from './components/Markattendance';
import UserDashboard from './components/UserDashboard';
import Adminattendance from './components/Adminattendance';
import Userprofile from './components/Userprofile';
function App(){
    const [token, setToken] = useState('');
    const [clickeduser, setClickeduser] = useState("");
    const [email,setEmail]=useState("");
    const [error, setError] = useState("");
    const [image, setImage] = useState("");
    const [rollno,setRollno]=useState("");
    const [name,setName]=useState("");
    const handleclickeduser = (email) => {
      console.log(email);
      setClickeduser(email);
      console.log(email);
    
    };
    const handlelogout = (event) => {
      setToken('');
            setEmail("");
            setRollno("");
            setName("");
            setImage("");
    
    };
    const handleSubmit= (event) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);
        const email = data.get('email');
        const password = data.get('password');
        fetch('http://localhost:9000/admin', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: email,
            password: password,
          }),
        })
        .then(response => response.json())
        .then(data => {
          if (data.token) {
            setError("");
            setEmail(data.email);
            setImage(data.image);
            setRollno(data.rollno);
            setName(data.name);
            setToken(data.token);
            console.log(data.token);
           
            axios.post('http://localhost:9000/admin/userslist', {
              headers: {
                'token': `${data.token}`
              }
            })           
            .catch(error => {
              console.error('Error:', error);
            });
          } else {
            setToken('');
            setEmail("");
            setRollno("");
            setName("");
            setError("Email or Password is incorrect please try again");
          }
        })
        .catch(error => {
          console.error('Error:', error);
        });}
        const handleuserSubmit= (event) => {
            event.preventDefault();
            const data = new FormData(event.currentTarget);
            const email = data.get('email');
            const password = data.get('password');
          
            fetch('http://localhost:9000/user/login', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify({
                email: email,
                password: password,
              }),
            })
            .then(response => response.json())
            .then(data => {
              if (data.token) {
                setError("");
                setRollno(data.rollno);
                setEmail(data.email);
                setImage(data.image);
                setName(data.name);
                setToken(data.token);
                console.log(data.token);
                console.log(data.image); 
               
                axios.post('http://localhost:9000/user/dashboard', {
                  "email":email,
                 "password":password
                },
                 {headers: {
                    'token': `${data.token}`
                  }
    
                })
                
                .catch(error => {
                  console.error('Error:', error);
                });
              } else {setToken('');
           
                setEmail("");
                setRollno("");
                setImage("");
              
                setError("Email or Password is incorrect please try again");
              }
            })
            .catch(error => {
              console.error('Error:', error);
            });}

return (
    <Router>
         <div className="App"> 
     <Routes>
       <Route path="/" element={<Loginadmin token={token} handleSubmit={handleSubmit} error={error} />} />
       <Route path="/admin" element={<Loginadmin token={token} handleSubmit={handleSubmit} error={error} />} />
       <Route path="/user/login" element={<Loginuser token={token} image={image} handleuserSubmit={handleuserSubmit} error={error} />} />
       <Route path="/admin/dashboard" element={<AdminDashboard handlelogout={handlelogout} handleclickeduser={handleclickeduser} token={token}/>}/>
       <Route path="/admin/usercreate" element={<SignUp />} />
       <Route path="/user/markattendance/webcam" element={<Markattendance name={name} image={image} email={email} rollno={rollno} />} />
       <Route path="/user/dashboard" element={<UserDashboard  handlelogout={handlelogout} token={token} name={name} email={email} />} />
       <Route path="/admin/attendance" element={<Adminattendance token={token} clickeduser={clickeduser} />} />
       <Route path="/user/userprofile" element={<Userprofile handlelogout={handlelogout} name={name} rollno ={rollno} email={email} image={image} />} />
     </Routes>
   </div>
    </Router>
);
}
export default App