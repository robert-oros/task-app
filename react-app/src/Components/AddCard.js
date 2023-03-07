import React, {useState} from "react"

function addCard(name, boardId, listId){
    fetch("http://localhost:8081/add_card", {
      method: 'POST',
      body: JSON.stringify({
        BoardID: boardId,
        ListId: listId,
        Text: name,
      }),
    })
    .then(res => console.log(res))
    .catch(error => {
      console.error(error)
    })
}

function AddCard(props){

    const [close, setClose] = useState(props.close)
    const [name, setValue] = useState("")
    const boardId = props.boardId
    const listId = props.listId

    const handleChange = (event) => {
        setValue(event.target.value)
    }

    const handleSubmit = () => {
        setClose(!close)
    }

    return(
        <div>
            <input type="text" value={name} placeholder="Card Name" onChange={handleChange}/>
            <input type="submit" value="Submit" onClick={handleSubmit}/>
            {close === false ? addCard(name, boardId, listId): <div></div>}
      
        </div>
    )
}

export default AddCard