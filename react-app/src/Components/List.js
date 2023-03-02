import Cards from './Cards'

function List(props) {
  return (
    <div>
      <ul>
        <li>
          <span className="label">{props.data.listId}</span>
        </li>
        <li>
          <span className="label">{props.data.title}</span>
        </li>
      </ul>

      <Cards cards={props.data.cards}/>
    </div>
  );
}
  
export default List;