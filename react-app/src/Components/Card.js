import '../css/Card.css'

function Card(props) {
  return (
    <ul className="card-text-container">
      {/* <li>
        <span className="label">{props.data.cardId}</span>
      </li> */}
      <li>
        <span className="card-text">{props.data.text}</span>
      </li>
    </ul>
  );
}

export default Card;