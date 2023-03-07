import React, { useState } from "react";

function addList(name, boardId){
    fetch("http://localhost:8081/add_list", {
      method: 'POST',
      body: JSON.stringify({
        BoardId: boardId,
        Title: name
      }),
    })
    .then(res => console.log(res))
    .catch(error => {
      console.error(error)
    })
}

function AddList(props){
    
    const [close, setClose] = useState(props.close)
    const [name, setValue] = useState("")

    const boardId = props.boardId

    const handleChange = (event) => {
        setValue(event.target.value)
    }

    const handleSubmit = () => {
        setClose(!close)
    }

    return (
        <div>
            {console.log(close)}
            <input type="text" value={name} placeholder="List Name" onChange={handleChange}/>
            <input type="submit" value="Submit" onClick={handleSubmit}/>
            {close === false ? addList(name, boardId) : <div></div>}
      
        </div>
    )
}

export default AddList