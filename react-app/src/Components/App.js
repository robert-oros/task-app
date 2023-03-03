import React, { Component } from 'react';
import Board from './Board';
import Navbar from './Navbar'

class App extends Component {
  constructor(){
    super()
    this.state = {
      boards: [],
      isLoaded: false,
      showComponent: false
    }

    this.displayBoard = this.displayBoard.bind(this)
  }

  componentDidMount(){
    fetch("http://localhost:8081/get_boards", {
      method: 'GET'
    })
    .then(res => res.json())
    .then(data => {
      this.setState({ boards: data})
      this.setState({ isLoaded: true })
    })
  }
  
  displayBoard(boardId) {
    this.setState({showComponent: boardId})

    if (this.state.showComponent == boardId) {
      this.setState({showComponent: false})
    }

    console.log(this.state.showComponent)
  }

  render() {
    let boards, boardsName;
    if (typeof this.state.boards !== 'undefined') {
      // boards = this.state.boards?.map(b => {
      //   return <Board data={b} />
      // })

      boardsName = this.state.boards?.map(b => {
        return <div>
            <p onClick={() => this.displayBoard(b.boardId)}>{b.name}</p>
            {this.state.showComponent && <Board data={b} />}
          </div>
          
      })
    }

    if (this.state.isLoaded) {
      return (
        <div>
          <Navbar />
          <div className='container'>
            {boardsName}
          </div>
        </div>
      );
    } else {
      return (
        <div>
          <Navbar />
          <h4 style={{textAlign: "center", paddingTop: '50%', paddingBottom: '50%'}}>Loading Board...</h4>
        </div>
      );
    }
  }
}

export default App;
