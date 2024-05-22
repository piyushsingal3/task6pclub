import { useNavigate } from "react-router-dom";
import "./UserDashboard.css";
import { useEffect, useState } from "react";
import axios from "axios";

export default function Adminattendance({ token, clickeduser }) {
  const navigate = useNavigate();
  const [userattendancedata, setUserattendancedata] = useState([]);



console.log(clickeduser);
console.log(token);

useEffect(() => {
    axios.post('http://localhost:9000/user/dashboard', {
      "email":clickeduser
    }, {
        headers: {
            'Content-Type': 'application/json',
            'token': token
        }
    })
    .then(response => {
        setUserattendancedata(response.data);
    })
    .catch(error => {
        console.error('Error:', error);
    });
}, [clickeduser, token]);
  const handleattend=()=>{
   
      navigate('/user/markattendance/webcam');
    
    };
    console.log(clickeduser);
  return (
    <div className="page">
     
      <div className="attendance">
        <h1 className="title">Attendance Checking System IITK</h1>
        <h2 className="title">Attendance of {clickeduser}</h2>

        {userattendancedata.length === 0 ? (
          <p>No attendance data to show</p>
        ) : (
          userattendancedata.map((entry, index) => (
            entry.email && (
              <div key={index} className="attendance-entry">
               
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
