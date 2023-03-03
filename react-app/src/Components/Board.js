import List from "./List";
import '../css/Board.css'

function Board(props) {
  const lists = props.data.lists.map(l => {
    return <List data={l}/>
  })
  
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
        <div class="row align-items-center">
          {lists}
        </div>
      </div>
    </div>
  );
}
  
export default Board;