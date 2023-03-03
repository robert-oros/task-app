import List from "./List";

function Board(props) {
  const lists = props.data.lists.map(l => {
    return <List data={l}/>
  })
  
  return (
    <div>
      <ul>
        <li>
          <span className="label">{props.data.boardId}</span>
        </li>
        <li>
          <span className="label">{props.data.name}</span>
        </li>
      </ul>
      
      {lists}
    </div>
  );
}
  
export default Board;