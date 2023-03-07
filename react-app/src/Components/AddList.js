import React, { useState } from "react";

function addList(name, boardId){
    fetch("http://localhost:8081/add_list", {
      method: 'POST',
      body: JSON.stringify({
        ListId:1,
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
    const [boardId, setBoardId] = useState(props.boardId)

    const handleChange = (event) => {
        setValue(event.target.value)
    }

    const Change = (event) => {
        setBoardId(event.target.value)
    }

    const handleSubmit = () => {
        setClose(!close)
    }

    return (
        <div>
            {console.log(close)}
            <input type="text" value={boardId} placeholder="Board Id" onChange={Change}/>
            <input type="text" value={name} placeholder="List Name" onChange={handleChange}/>
            <input type="submit" value="Submit" onClick={handleSubmit}/>
            {close === false ? addList(name, boardId) : <div>else</div>}
      
        </div>
    )
}

export default AddList