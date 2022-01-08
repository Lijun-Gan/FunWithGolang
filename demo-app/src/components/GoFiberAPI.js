
import React, {useState} from "react";
import axios from 'axios';

const GoFiberAPI = () => {

    const [docs, setDocs] = useState([])

    const handleSearch = () =>{
     
      axios.get('http://localhost:5000/api/external-api')
      .then(data => {
        console.log(data.data.docs)
        setDocs(data.data.docs)
      })
    }
    
    return (
        <div className="App">
          <h1>Welcome to Lijun's book store !</h1>
          
          <button onClick={handleSearch} >Search API</button>
          <ul>
            {docs.map(data => (
              <li key={data._id}>id: {data._id} , name: {data.name} </li>
              ))}
          </ul>

          
    
        </div>
      );
 }



export default GoFiberAPI
