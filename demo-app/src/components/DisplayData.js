import React, { useState} from "react";
import { useQuery, gql, useMutation } from "@apollo/client";

const QUERY_ALL_USERS = gql`
    query getAllStudents {
    students {
        id
        name 
        age
        major
        friends {
            name 
            }
        }
    }
`

const CREATE_STUDENT_MUTATION = gql`
    mutation CreateStudent($createStudentInput: CreateStudentInput!){
        createStudent(input: $createStudentInput) {
            id
            name
            age
            major
        }
    }
`

const DELETE_STUDENT_MUTATION = gql`
    mutation($deleteStudentID: ID!){
      deleteStudent(id: $deleteStudentID) {
        id
        }
    }
`

function DisplayData() {
    const {data, loading, error} = useQuery(QUERY_ALL_USERS);
    const [name, setName] = useState("");
    const [age, setAge] = useState("");
    const [major, setMajor] = useState("");

    const [createStudent] = useMutation(CREATE_STUDENT_MUTATION, {
        refetchQueries: [{ query: QUERY_ALL_USERS }],
    } )
    const [deleteOneStudent] = useMutation(DELETE_STUDENT_MUTATION, {
        refetchQueries: [{ query: QUERY_ALL_USERS }],
    } )

    if(loading) {
        return <h1>Data is loading.....</h1>
    }

    if (data){
        console.log(data.students)
    }

    if(error){
        console.log(error)
    }

    return(

        <div>
            <div className="addNewStudent">
                <input
                    type="text"
                    placeholder="Name..."
                    value={name}
                    onChange={(event) => {
                        setName(event.target.value);
                    }}
                />
                <input
                    type="number"
                    placeholder="Age..."
                    value={age}
                    onChange={(event) => {
                        setAge(event.target.value);
                    }}
                />
                <input
                    type="text"
                    placeholder="Major..."
                    value={major}
                    onChange={(event) => {
                        setMajor(event.target.value);
                    }}
                />

            <button onClick={()=>{
                if( name && age && major ){
                    console.log({name: name, age: Number(age), major: major})
                    createStudent({
                        variables:{
                            createStudentInput: {name: name, age: Number(age), major: major},
                        },
                    })
                    setName("");
                    setAge("");
                    setMajor("");
                    // data.push({name: name, age: Number(age), major: major})
                }
        }}>Add Student</button>

            </div>
            
            {data && data.students.map((student)=>{
                return(
                <div key={student.id} className="studentList">
                    <p>Name: {student.name} </p>
                    <p>Age: {student.age} </p>
                    <p>Major:  {student.major}</p>
                    {student.friends ? <p>Friends: {student.friends.map((f,i)=>(<span key={i}>{f.name}, </span>))}</p> : null}
                    {/* <button>Update Name</button> */}
                    <button onClick={()=>{
                        console.log(typeof student['id'])
                        deleteOneStudent({
                            variables: {
                                deleteStudentID: student['id']
                            } 
                        })
                    } 
                    }>Delete</button>
                </div >
                )
            })}
        </div>

    )
 
}

export default DisplayData;