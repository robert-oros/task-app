import React, { Component } from 'react';
import Board from './Board';
import Navbar from './Navbar'

class App extends Component {
  constructor(){
    super()
    this.state = {
      boards: [],
      isLoaded: false,
      showComponent: false,
      showInput: false
    }

    this.setShowComponent = this.setShowComponent.bind(this)
    this.setStateForInput = this.setStateForInput.bind(this)
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
  
  setShowComponent(boardId) {
    this.setState({showComponent: boardId})

    if (this.state.showComponent == boardId) {
      this.setState({showComponent: false})
    }

    console.log(this.state.showComponent)
  }

  addBoard(name, list){
    fetch("http://localhost:8081/add_board", {
      method: 'POST',
      body: JSON.stringify({
        Name: '',
      }),
    })
    .then(res => res.json())
    .catch(error => {
      console.log(error)
    })
  }

  setStateForInput(){
    this.setState({
      showInput: true
    })
  }

  render() {
    let boardsName;
    if (typeof this.state.boards !== 'undefined') {
      boardsName = this.state.boards?.map(b => {
        return <div>
            <p style={{cursor: 'pointer'}} onClick={() => this.setShowComponent(b.boardId)}>{b.name}</p>
            {this.state.showComponent && (b.boardId == this.state.showComponent) && <Board data={b} />}
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
