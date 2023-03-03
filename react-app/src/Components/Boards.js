import Board from "./Board";

function Boards(props) {
    const boards = props.data.boards.map(b => {
        return <Board data={b}/>
    })

    return (
        <div>
            {boards}
        </div>
    );
}
    
export default Boards;