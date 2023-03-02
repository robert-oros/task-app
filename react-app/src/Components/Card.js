import '../css/Card.css'

function Card(props) {
  return (
    <ul>
      <li>
        <span className="label">{props.data.cardId}</span>
      </li>
      <li>
        <span className="label">{props.data.text}</span>
      </li>
    </ul>
  );
}

export default App;