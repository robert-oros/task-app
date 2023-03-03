import Card from './Card';
import '../css/List.css'

function List(props) {
  const cards = props.data.cards.map(c => {
    return <Card data={c}/>
  })

  return (

      <div class="col list-container">
        <ul className='list-text'>
          {/* <li>
            <span className="label">{props.data.listId}</span>
          </li> */}
          <li>
            <span className="label">{props.data.title}</span>
          </li>
        </ul>

        <div className="card-list">
          {cards}

          <ul className="card-text-container">
            <button type="button">Adauga Todo</button>
          </ul>

        </div>
      </div>
  );
}
  
export default List;