

import styles from "./Markattendance.css";
import React, { useCallback, useRef, useState } from "react";
import Webcam from "react-webcam";
import axios from "axios";
import { useNavigate } from "react-router-dom";


const API_KEY = "ZNoMmPcy7VxEqFNNB9jSE8TCPqg0Ntwr";
const API_SECRET = "ah7Gwk9k92O3qGS0cnzb_6CNvTkUOLdN";
const CompareFaceAPI = "https://api-us.faceplusplus.com/facepp/v3/compare";

export default function MarkAttendance({name,image,email,rollno}) {
  const [img, setImg] = useState(null);
  const [comparisonResult, setComparisonResult] = useState(null);
  const webcamRef = useRef(null);
const navigate =useNavigate();
  const capture = useCallback(() => {
    const imageSrc = webcamRef.current.getScreenshot();
    setImg(imageSrc);
  }, [webcamRef]);

  const compareFaces = async (sourceImage, image) => {
    const formData = new FormData();
    formData.append("api_key", API_KEY);
    formData.append("api_secret", API_SECRET);
    formData.append("image_base64_1", sourceImage.replace(/^data:image\/\w+;base64,/, ""));
    formData.append("image_base64_2", image.replace(/^data:image\/\w+;base64,/, ""));

    try {
      const response = await axios.post(CompareFaceAPI, formData, {
        headers: { "Content-Type": "multipart/form-data" },
      });
      return response.data;
    } catch (error) {
      console.error("Error comparing faces: ", error);
      return null;
    }
  };
  const appStyle = {
    backgroundColor: 'white',
    

    margin: 0,
    padding: 10
    
  };
  const markAttendance = async () => {
    try {
        const response = await axios.post('http://localhost:9000/markattendance', {
            email: email,
            rollno: rollno,
            date: new Date().toISOString().split('T')[0],
            status: "Present",
        });
        if (response.status === 200) {
            alert('Attendance marked successfully!.Click here to go to dashboard again');
            navigate("/user/dashboard")
        }
    } catch (error) {
       
        if (error.response) {
            if (error.response.status === 500) {
                alert('Attendance already marked.Click here to go to dashboard again');
                navigate("/user/dashboard")
            } else {
                alert('An error occurred while marking attendance.');
            }
        } else {
            alert('An error occurred while marking attendance.');
        }
    }
};

  const handleCompare = async () => {
    if (img) {
      const result = await compareFaces(img, image);
      setComparisonResult(result.confidence);
    }
  };

  return (
    <div  style={appStyle} className="Container">
    <h1>Capture your Image to mark your attendance</h1>
      {img === null ? (
        <>
          <Webcam className="cam"
            audio={false}
            mirrored={true}
            height={600}
            width={600}
            ref={webcamRef}
            screenshotFormat="image/jpeg"
          />
          <button className="button" onClick={capture}>Capture photo</button>
        </>
      ) : (
        <>
          <img src={img} alt="screenshot" />
          <button className="button" onClick={() => {setImg(null);setComparisonResult(null)}}>Retake</button>
          <button className="button" onClick={handleCompare}>Compare Faces</button>
        </>
      )}
      
     { comparisonResult>85 && (
        <div className="compare">
            <h2>Comparison Result</h2>
            <h3>Hello {name}!</h3>
        <p >Your image matched click this button to mark your attendance</p>
        <button className="button" onClick={markAttendance}>Mark Attendance for today</button>
        </div>
      )}{ comparisonResult<=85 && comparisonResult!=null && (
        <div>
        <h3>Your image didn't match please retake</h3>
        
        </div>
      )}
    </div>
  );
}
