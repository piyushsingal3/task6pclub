
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import axios from "axios";
import './Admindashboard.css'; 

export default function AdminDashboard({handlelogout, handleclickeduser,token }) {
    const navigate = useNavigate();
    const [userdata, setUserdata] = useState([]);
   
    
    useEffect(() => {
        axios.post('http://localhost:9000/admin/userslist', {}, {
            headers: {
                'Content-Type': 'application/json',
                'token': token
            }
        })
        .then(response => {
            setUserdata(response.data);
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }, [token]);
const handlecreate=()=>{
navigate("/admin/usercreate");
}
const handleuserclick=(email)=>{
    handleclickeduser(email);
    navigate("/admin/attendance")
}
const handleLogout = () => {
    handlelogout();
    navigate('/admin');
  };

    return (
        <div className="admin-dashboard"  style={{ backgroundImage: 'none' }}>
            <div className="sidebar">
                <h2>Dashboard</h2>
                <ul>
                    <li>Home</li>
                    <li onClick={handlecreate}>Create Users</li>
                    <li onClick={handleLogout}>Log Out</li>
                </ul>
            </div>

            <div className="content">
            <h1>Students List</h1>
            <h3>Click on any user to see his/her attendance</h3>
                {userdata.length === 0 ? (
                    <p>No student data found or admin not logged in properly.Please login again</p>
                ) : (
                    
                    userdata.map((entry, index) => (
                        
                        entry.email && (
                            <button className="user" onClick={() =>handleuserclick(entry.email)}><div key={index} className="attendance-entry">
                                <h3>{entry.name}</h3>
                                <p>{entry.email}</p>
                                {entry.image && (
                                    <img src={`data:image/png;base64,${entry.image}`} alt="User" className="user-image" />
                                )}
                                <br />
                            </div>
                            </button>
                        )
                    ))
                )}
            </div>
        </div>
    );
}

