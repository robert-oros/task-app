import List from "./List";
import '../css/Board.css'
import { useState } from "react";
import AddList from "./AddList"
import Popup from "./Popup";

function Board(props) {
  let lists;
  if (props.data.lists !== null) {
    lists = props.data.lists.map(l => {
      return <List data={l}/>
    })
  }
  const [ showPopup, setShowPopup] = useState(false)
  const handleClick = () => {
    setShowPopup(!showPopup)
  }
  
  return (
    <div class="container board-container">
      {/* <ul className="board-text">
        <li>
          <span className="label">{props.data.boardId}</span>
        </li> 
        <li>
          <span className="label">{props.data.name}</span>
        </li>
      </ul> */}

      <div class="container">
      <button onClick={handleClick}>Adauga Lista</button>
      {showPopup ? <Popup content={<><AddList boardId={props.data.boardId} close={true}/></>}  handleClose={handleClick} />: <div></div>}
        <div class="row align-items-center">
          {lists}
        </div>
        
      </div>
    </div>
  );
}
  
export default Board;