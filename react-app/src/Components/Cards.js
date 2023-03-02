import Card from './Card'

function Cards(props) {
    const cards = props.cards.map(c => {
        return <Card data={c}/>
    })

    return (
        <div>
            {cards}
        </div>
    );
}
  
export default Cards;