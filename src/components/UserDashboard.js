
    
import { useNavigate } from "react-router-dom";
import "./UserDashboard.css";
import { useEffect, useState } from "react";
import axios from "axios";

export default function UserDashboard({ handlelogout,token,name, email }) {
  const navigate = useNavigate();
  const [attendancedata, setAttendancedata] = useState([]);



console.log(email);

useEffect(() => {
    axios.post('http://localhost:9000/user/dashboard', {
        email: email
    }, {
        headers: {
            'Content-Type': 'application/json',
            'token': token
        }
    })
    .then(response => {
        setAttendancedata(response.data);
    })
    .catch(error => {
        console.error('Error:', error);
    });
}, [email, token]);
  const handleattend=()=>{
   
      navigate('/user/markattendance/webcam');
    
    };
    const handleuserprofile=()=>{
   
        navigate('/user/userprofile');
      
      };
      const handleLogout = () => {
        handlelogout();
        navigate('/user/login');
      };
    
  return (
    <div className="page">
      <div className="sidebar">
        <h2>Dashboard</h2>
        <ul>
          <li>Home</li>
          <li onClick={handleattend}>Mark Attendance for today</li>
          <li onClick={handleuserprofile}>User Profile</li>
          
          <li onClick={handleLogout}>Log Out</li>
        </ul>
      </div>
      <div className="attendance">
        <h1 className="title">Attendance Checking System IITK</h1>
        <h2 className="title">Hello {name} you can find your attendance here</h2>

        {attendancedata.length === 0 ? (
          <p>No attendance data to show</p>
        ) : (
          attendancedata.map((entry, index) => (
            entry.email && (
              <div key={index} className="attendance-entry">
                <h3>{entry.rollno}</h3>
                <h3>{entry.date} : {entry.status}</h3>
                <br></br>
              </div>
            )
          ))
        )}
      </div>
    </div>
  );
}
